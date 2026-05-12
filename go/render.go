package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Render(w io.Writer, cfg Config, entries []MonthlyEntry, monthlies []int, avg int, j GradeJudgement, a Alert) error {
	parts := []string{"[算定] "}
	for i, e := range entries {
		label := "見込"
		if e.Input.IsActual {
			label = "実績"
		}
		parts = append(parts, fmt.Sprintf(" %d月: %s %s", e.Month, label, yen(monthlies[i])))
	}
	if _, err := fmt.Fprintln(w, strings.Join(parts, "")); err != nil {
		return err
	}

	if _, err := fmt.Fprintf(w, "[平均]  %s\n", yen(avg)); err != nil {
		return err
	}

	kenpo := "-"
	if j.KenpoGrade != nil {
		kenpo = fmt.Sprintf("%d級", *j.KenpoGrade)
	}
	if _, err := fmt.Fprintf(w, "[等級]  健保 %s / 厚年 %d級   標準報酬月額 %s\n", kenpo, j.Pension, yen(j.StandardAmount)); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(w); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(w, "[アラート]"); err != nil {
		return err
	}
	if j.NextThreshold == nil {
		_, err := fmt.Fprintln(w, "  - 既に最上位等級です")
		return err
	}
	if _, err := fmt.Fprintf(w, "  - 次の等級まであと %s（3ヶ月合計）\n", yen(a.AmountUntil)); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(w, "  - 残り月で許容できる追加残業: 約 %.1f 時間（時給 %s × 1.25 換算）\n", a.AllowedHours, yen(int(cfg.HourlyRate))); err != nil {
		return err
	}
	for _, b := range cfg.Benefits {
		if _, err := fmt.Fprintf(w, "  - %s: あと %d 回まで利用可\n", b.Name, a.BenefitLimits[b.Name]); err != nil {
			return err
		}
	}
	return nil
}

func yen(v int) string {
	sign := ""
	if v < 0 {
		sign = "-"
		v = -v
	}
	s := strconv.Itoa(v)
	var out []byte
	for i, c := range []byte(s) {
		if i > 0 && (len(s)-i)%3 == 0 {
			out = append(out, ',')
		}
		out = append(out, c)
	}
	return "¥" + sign + string(out)
}
