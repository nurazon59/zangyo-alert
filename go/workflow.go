package main

import (
	"fmt"
	"io"
	"os"
)

type Args struct {
	ConfigPath string
	GradesPath string
	ImportPath string
	Stdin      io.Reader
	Stdout     io.Writer
	Stderr     io.Writer
}

func Run(args Args) int {
	if args.Stdin == nil {
		args.Stdin = os.Stdin
	}
	if args.Stdout == nil {
		args.Stdout = os.Stdout
	}
	if args.Stderr == nil {
		args.Stderr = os.Stderr
	}

	cfg, err := LoadConfig(args.ConfigPath)
	if err != nil {
		fmt.Fprintln(args.Stderr, err)
		return 1
	}
	if cfg.HourlyRateDerived {
		fmt.Fprintf(args.Stdout, "[時給]  %s（自動計算: 基本給 %s ÷ %.1fh/月）\n",
			yen(int(cfg.HourlyRate)), yen(cfg.BaseSalary), cfg.MonthlyWorkingHours)
	}

	table, err := LoadGrades(args.GradesPath)
	if err != nil {
		fmt.Fprintln(args.Stderr, err)
		return 1
	}

	importedAmounts := map[int]int{}
	if args.ImportPath != "" {
		a, err := LoadAttendance(args.ImportPath)
		if err != nil {
			fmt.Fprintln(args.Stderr, err)
			return 1
		}
		bd := cfg.OvertimePayFromAttendance(a)
		amount := cfg.EstimateGrossFromAttendance(a)
		importedAmounts[a.Month] = amount
		fmt.Fprintf(args.Stdout,
			"[import] %s 月度: 平日残業=%.2fh (請求対象=%.2fh), 深夜=%.2fh, 休日=%.2fh, 残業代=¥%d, 見込み総支給=¥%d\n",
			a.Period, bd.WeekdayHours, bd.Chargeable, bd.LateNightHours, bd.HolidayHours, bd.Pay, amount)
	}

	prompter := NewPrompterFrom(args.Stdin, args.Stdout)

	var entries []MonthlyEntry
	for _, m := range []int{4, 5, 6} {
		imported, hasImported := importedAmounts[m]
		entry, err := promptMonth(prompter, cfg, m, imported, hasImported)
		if err != nil {
			fmt.Fprintln(args.Stderr, err)
			return 1
		}
		entries = append(entries, entry)
	}

	monthlies := make([]int, len(entries))
	for i, e := range entries {
		monthlies[i] = cfg.MonthlyGross(e)
	}
	avg := Average(monthlies...)
	j, err := Judge(table, avg)
	if err != nil {
		fmt.Fprintln(args.Stderr, err)
		return 1
	}
	alert := cfg.ComputeAlert(j, entries, avg)

	fmt.Fprint(args.Stdout, "\n")
	if err := Render(args.Stdout, cfg, entries, monthlies, avg, j, alert); err != nil {
		fmt.Fprintln(args.Stderr, err)
		return 1
	}
	return 0
}

func promptMonth(p *Prompter, cfg Config, month int, importedAmount int, hasImported bool) (MonthlyEntry, error) {
	fmt.Fprintf(p.w, "\n[%d月]\n", month)
	useActual, err := p.AskYesNo("実績で入力しますか？ (y=実績/n=見込み):")
	if err != nil {
		return MonthlyEntry{}, err
	}

	var input MonthInput
	if useActual {
		var def *int
		if hasImported {
			def = &importedAmount
		}
		amount, err := p.AskInt("総支給額:", def)
		if err != nil {
			return MonthlyEntry{}, err
		}
		input = MonthInput{IsActual: true, Amount: amount}
	} else {
		zero := 0.0
		hours, err := p.AskFloat("追加残業時間(h):", &zero)
		if err != nil {
			return MonthlyEntry{}, err
		}
		input = MonthInput{IsActual: false, ExtraHours: hours}
	}

	usages := make([]BenefitUsage, 0, len(cfg.Benefits))
	for _, b := range cfg.Benefits {
		zero := 0
		n, err := p.AskInt(fmt.Sprintf("%s (¥%d/回) を何回利用？:", b.Name, b.UnitAmount), &zero)
		if err != nil {
			return MonthlyEntry{}, err
		}
		usages = append(usages, BenefitUsage{Benefit: b, Count: n})
	}
	return MonthlyEntry{Month: month, Input: input, BenefitUsages: usages}, nil
}
