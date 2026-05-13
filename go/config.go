package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read config %s: %w", path, err)
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("parse config %s: %w", path, err)
	}
	if cfg.Benefits == nil {
		cfg.Benefits = []Benefit{}
	}
	if cfg.HourlyRate <= 0 {
		if cfg.MonthlyWorkingHours <= 0 {
			cfg.MonthlyWorkingHours = 160.0
		}
		cfg.HourlyRate = float64(cfg.BaseSalary) / cfg.MonthlyWorkingHours
		cfg.HourlyRateDerived = true
	}
	return cfg, nil
}
