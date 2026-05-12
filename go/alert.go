package main

func (c Config) ComputeAlert(j GradeJudgement, entries []MonthlyEntry, avg int) Alert {
	limits := map[string]int{}
	for _, b := range c.Benefits {
		limits[b.Name] = 0
	}
	if j.NextThreshold == nil {
		return Alert{AmountUntil: 0, AllowedHours: 0, BenefitLimits: limits}
	}

	amountUntilTotal := (*j.NextThreshold - avg) * 3
	if amountUntilTotal < 0 {
		amountUntilTotal = 0
	}

	remaining := 0
	for _, e := range entries {
		if !e.Input.IsActual {
			remaining++
		}
	}

	rate := c.HourlyRate * WeekdayOvertimeMultiplier
	allowedHours := 0.0
	if remaining > 0 && rate > 0 {
		allowedHours = float64(amountUntilTotal) / rate / float64(remaining)
	}

	for _, b := range c.Benefits {
		if b.UnitAmount > 0 {
			limits[b.Name] = amountUntilTotal / b.UnitAmount
		}
	}

	return Alert{
		AmountUntil:   amountUntilTotal,
		AllowedHours:  allowedHours,
		BenefitLimits: limits,
	}
}
