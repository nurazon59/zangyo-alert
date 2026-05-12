package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

type MonthlyAttendance struct {
	Period          string
	Month           int
	WeekdayOvertime float64
	LateNightHours  float64
	HolidayOvertime float64
}

var expectedHeaders = []struct {
	Cell, Want string
}{
	{"A1", "勤怠月度"},
	{"Q1", "深夜時間合計"},
	{"Y1", "法定外平日フレックス残業時間"},
	{"Z1", "休日フレックス残業時間"},
}

func LoadAttendance(path string) (MonthlyAttendance, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return MonthlyAttendance{}, fmt.Errorf("open xlsx %s: %w", path, err)
	}
	defer f.Close()

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return MonthlyAttendance{}, fmt.Errorf("no sheet in %s", path)
	}
	sheet := sheets[0]

	for _, h := range expectedHeaders {
		got, err := f.GetCellValue(sheet, h.Cell)
		if err != nil {
			return MonthlyAttendance{}, fmt.Errorf("read %s: %w", h.Cell, err)
		}
		if got != h.Want {
			return MonthlyAttendance{}, fmt.Errorf("unexpected header at %s: expected %q but got %q", h.Cell, h.Want, got)
		}
	}

	num := func(addr string) (float64, error) {
		v, err := f.GetCellValue(sheet, addr)
		if err != nil {
			return 0, err
		}
		v = strings.TrimSpace(v)
		if v == "" {
			return 0, nil
		}
		return strconv.ParseFloat(v, 64)
	}

	period, err := f.GetCellValue(sheet, "A2")
	if err != nil {
		return MonthlyAttendance{}, err
	}
	month, err := parsePeriodMonth(period)
	if err != nil {
		return MonthlyAttendance{}, err
	}

	a := MonthlyAttendance{Period: period, Month: month}
	for _, p := range []struct {
		addr string
		dst  *float64
	}{
		{"Q2", &a.LateNightHours},
		{"Y2", &a.WeekdayOvertime},
		{"Z2", &a.HolidayOvertime},
	} {
		v, err := num(p.addr)
		if err != nil {
			return MonthlyAttendance{}, fmt.Errorf("parse %s: %w", p.addr, err)
		}
		*p.dst = v
	}
	return a, nil
}

func parsePeriodMonth(period string) (int, error) {
	digits := []rune{}
	prev := rune(0)
	for _, r := range period {
		if r >= '0' && r <= '9' {
			digits = append(digits, r)
		} else if len(digits) > 0 {
			if prev >= '0' && prev <= '9' {
				if m, err := strconv.Atoi(string(digits)); err == nil && m >= 1 && m <= 12 {
					if r == '月' {
						return m, nil
					}
				}
			}
			digits = digits[:0]
		}
		prev = r
	}
	for _, tok := range strings.FieldsFunc(period, func(r rune) bool {
		return r < '0' || r > '9'
	}) {
		if m, err := strconv.Atoi(tok); err == nil && m >= 1 && m <= 12 {
			return m, nil
		}
	}
	return 0, fmt.Errorf("cannot parse month from period %q", period)
}
