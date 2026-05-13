package main

type Benefit struct {
	Name       string `json:"name"`
	UnitAmount int    `json:"unitAmount"`
}

type MonthInput struct {
	IsActual   bool
	Amount     int
	ExtraHours float64
}

type BenefitUsage struct {
	Benefit Benefit
	Count   int
}

type MonthlyEntry struct {
	Month         int
	Input         MonthInput
	BenefitUsages []BenefitUsage
}

type OvertimeRates struct {
	LateNight float64 `json:"lateNight"`
	Holiday   float64 `json:"holiday"`
}

type RawConfig struct {
	BaseSalary          int           `json:"baseSalary"`
	FixedOvertime       float64       `json:"fixedOvertime"`
	FixedOverSalary     int           `json:"fixedOverSalary"`
	FixedBenefits       int           `json:"fixedBenefits"`
	HourlyRate          float64       `json:"hourlyRate"`
	MonthlyWorkingHours float64       `json:"monthlyWorkingHours"`
	OvertimeRates       OvertimeRates `json:"overtimeRates"`
	Benefits            []Benefit     `json:"benefits"`
}

type Config struct {
	RawConfig
	HourlyRateDerived bool
}

type GradeRow struct {
	Grade      int
	KenpoGrade *int
	Amount     int
	Lower      *int
	Upper      *int
}

type GradeJudgement struct {
	StandardAmount int
	KenpoGrade     *int
	Pension        int
	NextThreshold  *int
}

type Alert struct {
	AmountUntil   int
	AllowedHours  float64
	BenefitLimits map[string]int
}
