package main

import (
	"os"
	"path/filepath"
	"testing"
)

func writeTempConfig(t *testing.T, content string) string {
	t.Helper()
	f := filepath.Join(t.TempDir(), "config.json")
	if err := os.WriteFile(f, []byte(content), 0600); err != nil {
		t.Fatalf("write temp config: %v", err)
	}
	return f
}

func TestLoadConfig_HourlyRateExplicit(t *testing.T) {
	path := writeTempConfig(t, `{"baseSalary":400000,"hourlyRate":2500}`)
	cfg, err := LoadConfig(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.HourlyRate != 2500 {
		t.Fatalf("hourlyRate = %v, want 2500", cfg.HourlyRate)
	}
	if cfg.HourlyRateDerived {
		t.Fatal("HourlyRateDerived should be false when hourlyRate is explicit")
	}
}

func TestLoadConfig_HourlyRateDerived_DefaultHours(t *testing.T) {
	path := writeTempConfig(t, `{"baseSalary":320000}`)
	cfg, err := LoadConfig(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 320000.0 / 160.0
	if cfg.HourlyRate != want {
		t.Fatalf("hourlyRate = %v, want %v (baseSalary/160)", cfg.HourlyRate, want)
	}
	if cfg.MonthlyWorkingHours != 160.0 {
		t.Fatalf("monthlyWorkingHours = %v, want 160", cfg.MonthlyWorkingHours)
	}
	if !cfg.HourlyRateDerived {
		t.Fatal("HourlyRateDerived should be true when hourlyRate is omitted")
	}
}

func TestLoadConfig_HourlyRateDerived_CustomHours(t *testing.T) {
	path := writeTempConfig(t, `{"baseSalary":360000,"monthlyWorkingHours":180}`)
	cfg, err := LoadConfig(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 360000.0 / 180.0
	if cfg.HourlyRate != want {
		t.Fatalf("hourlyRate = %v, want %v (baseSalary/180)", cfg.HourlyRate, want)
	}
	if !cfg.HourlyRateDerived {
		t.Fatal("HourlyRateDerived should be true when hourlyRate is omitted")
	}
}
