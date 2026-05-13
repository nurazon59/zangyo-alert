package main

import "testing"

func TestBenefitTotal(t *testing.T) {
	usages := []BenefitUsage{
		{Benefit: Benefit{Name: "住宅", UnitAmount: 10000}, Count: 2},
		{Benefit: Benefit{Name: "食事", UnitAmount: 500}, Count: 10},
	}
	if got := BenefitTotal(usages); got != 25000 {
		t.Fatalf("benefit total = %d, want 25000", got)
	}
}

func TestMonthlyGross_Actual(t *testing.T) {
	cfg := Config{}
	entry := MonthlyEntry{
		Input: MonthInput{IsActual: true, Amount: 350000},
		BenefitUsages: []BenefitUsage{
			{Benefit: Benefit{Name: "住宅", UnitAmount: 10000}, Count: 1},
		},
	}
	if got := cfg.MonthlyGross(entry); got != 360000 {
		t.Fatalf("monthly gross = %d, want 360000", got)
	}
}

func TestMonthlyGross_Estimate(t *testing.T) {
	cfg := Config{RawConfig: RawConfig{
		BaseSalary:      300000,
		FixedOverSalary: 60000,
		FixedBenefits:   40000,
		HourlyRate:      2000,
	}}
	entry := MonthlyEntry{
		Input: MonthInput{ExtraHours: 5},
	}
	want := 300000 + 60000 + 40000 + int(5*1.25*2000)
	if got := cfg.MonthlyGross(entry); got != want {
		t.Fatalf("monthly gross = %d, want %d", got, want)
	}
}

func TestAverage(t *testing.T) {
	if got := Average(300000, 310000, 320000); got != 310000 {
		t.Fatalf("avg = %d, want 310000", got)
	}
	if got := Average(100, 100, 101); got != 100 {
		t.Fatalf("trunc avg = %d, want 100", got)
	}
	if got := Average(); got != 0 {
		t.Fatalf("empty avg = %d, want 0 (zero-divide guard)", got)
	}
}

func baseRateConfig() Config {
	return Config{RawConfig: RawConfig{
		BaseSalary:      300000,
		FixedOvertime:   50.0,
		FixedOverSalary: 80000,
		FixedBenefits:   10000,
		HourlyRate:      2000,
		OvertimeRates:   OvertimeRates{LateNight: 0.25, Holiday: 1.35},
	}}
}

func TestOvertimePayFromAttendance_WithinFixedOvertime(t *testing.T) {
	cfg := baseRateConfig()
	a := MonthlyAttendance{WeekdayOvertime: 40}
	bd := cfg.OvertimePayFromAttendance(a)
	if bd.WeekdayHours != 40 {
		t.Fatalf("weekdayHours = %v, want 40", bd.WeekdayHours)
	}
	if bd.Chargeable != 0 {
		t.Fatalf("chargeable = %v, want 0 (within fixed)", bd.Chargeable)
	}
	if bd.Pay != 0 {
		t.Fatalf("pay = %d, want 0", bd.Pay)
	}
}

func TestOvertimePayFromAttendance_OverFixedOvertime(t *testing.T) {
	cfg := baseRateConfig()
	a := MonthlyAttendance{WeekdayOvertime: 60}
	bd := cfg.OvertimePayFromAttendance(a)
	if bd.Chargeable != 10 {
		t.Fatalf("chargeable = %v, want 10", bd.Chargeable)
	}
	want := int(10.0 * 1.25 * 2000)
	if bd.Pay != want {
		t.Fatalf("pay = %d, want %d", bd.Pay, want)
	}
}

func TestOvertimePayFromAttendance_LateNightOnly(t *testing.T) {
	cfg := baseRateConfig()
	a := MonthlyAttendance{LateNightHours: 8}
	bd := cfg.OvertimePayFromAttendance(a)
	want := int(8.0 * 0.25 * 2000)
	if bd.Pay != want {
		t.Fatalf("pay = %d, want %d", bd.Pay, want)
	}
	if bd.LateNightHours != 8 {
		t.Fatalf("lateNightHours = %v", bd.LateNightHours)
	}
}

func TestOvertimePayFromAttendance_HolidayOnly(t *testing.T) {
	cfg := baseRateConfig()
	a := MonthlyAttendance{HolidayOvertime: 10}
	bd := cfg.OvertimePayFromAttendance(a)
	if bd.HolidayHours != 10 {
		t.Fatalf("holidayHours = %v, want 10", bd.HolidayHours)
	}
	want := int(10.0 * 1.35 * 2000)
	if bd.Pay != want {
		t.Fatalf("pay = %d, want %d", bd.Pay, want)
	}
}

func TestOvertimePayFromAttendance_Mixed(t *testing.T) {
	cfg := baseRateConfig()
	a := MonthlyAttendance{
		WeekdayOvertime: 60,
		LateNightHours:  4,
		HolidayOvertime: 8,
	}
	bd := cfg.OvertimePayFromAttendance(a)
	if bd.Chargeable != 10 {
		t.Fatalf("chargeable = %v, want 10", bd.Chargeable)
	}
	if bd.HolidayHours != 8 {
		t.Fatalf("holidayHours = %v, want 8", bd.HolidayHours)
	}
	want := int((10.0*1.25 + 4.0*0.25 + 8.0*1.35) * 2000)
	if bd.Pay != want {
		t.Fatalf("pay = %d, want %d", bd.Pay, want)
	}
}

func TestEstimateGrossFromAttendance(t *testing.T) {
	cfg := baseRateConfig()
	a := MonthlyAttendance{WeekdayOvertime: 60}
	bd := cfg.OvertimePayFromAttendance(a)
	got := cfg.EstimateGrossFromAttendance(a)
	want := cfg.BaseSalary + cfg.FixedOverSalary + cfg.FixedBenefits + bd.Pay
	if got != want {
		t.Fatalf("gross = %d, want %d", got, want)
	}
}
