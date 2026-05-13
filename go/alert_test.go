package main

import "testing"

func TestComputeAlert_TopGradeNoAlert(t *testing.T) {
	cfg := Config{RawConfig: RawConfig{Benefits: []Benefit{{Name: "住宅", UnitAmount: 10000}}}}
	j := GradeJudgement{StandardAmount: 1390000, Pension: 50, NextThreshold: nil}
	a := cfg.ComputeAlert(j, nil, 1500000)
	if a.AmountUntil != 0 || a.AllowedHours != 0 {
		t.Fatalf("top grade should give zero alert: %+v", a)
	}
	if a.BenefitLimits["住宅"] != 0 {
		t.Fatalf("top grade benefit limit = %d", a.BenefitLimits["住宅"])
	}
}

func TestComputeAlert_RemainingMonths(t *testing.T) {
	cfg := Config{RawConfig: RawConfig{
		HourlyRate: 2000,
		Benefits: []Benefit{
			{Name: "住宅", UnitAmount: 10000},
			{Name: "食事", UnitAmount: 500},
		},
	}}
	threshold := 310000
	j := GradeJudgement{StandardAmount: 300000, Pension: 22, NextThreshold: &threshold}
	entries := []MonthlyEntry{
		{Month: 4, Input: MonthInput{IsActual: true, Amount: 300000}},
		{Month: 5, Input: MonthInput{}},
		{Month: 6, Input: MonthInput{}},
	}
	a := cfg.ComputeAlert(j, entries, 300000)
	if a.AmountUntil != 30000 {
		t.Fatalf("amount until = %d, want 30000", a.AmountUntil)
	}
	if a.AllowedHours != 6.0 {
		t.Fatalf("allowed hours = %v, want 6.0", a.AllowedHours)
	}
	if a.BenefitLimits["住宅"] != 3 {
		t.Fatalf("住宅 limit = %d, want 3", a.BenefitLimits["住宅"])
	}
	if a.BenefitLimits["食事"] != 60 {
		t.Fatalf("食事 limit = %d, want 60", a.BenefitLimits["食事"])
	}
}
