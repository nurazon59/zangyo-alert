package main

const WeekdayOvertimeMultiplier = 1.25

type OvertimePayBreakdown struct {
	WeekdayHours   float64
	Chargeable     float64
	LateNightHours float64
	HolidayHours   float64
	Pay            int
}

func BenefitTotal(usages []BenefitUsage) int {
	total := 0
	for _, u := range usages {
		total += u.Benefit.UnitAmount * u.Count
	}
	return total
}

func (c Config) MonthlyGross(entry MonthlyEntry) int {
	bt := BenefitTotal(entry.BenefitUsages)
	if entry.Input.IsActual {
		return entry.Input.Amount + bt
	}
	base := c.BaseSalary + c.FixedOverSalary + c.FixedBenefits
	extra := int(entry.Input.ExtraHours * WeekdayOvertimeMultiplier * c.HourlyRate)
	return base + extra + bt
}

// Average は0除算をガードする整数平均。要素0件のときは0を返す。
func Average(months ...int) int {
	if len(months) == 0 {
		return 0
	}
	sum := 0
	for _, m := range months {
		sum += m
	}
	return sum / len(months)
}

// OvertimePayFromAttendance は勤怠xlsxの Y2/Q2/Z2 から残業代を算出する。
// Y2 は法定外平日フレックス残業時間(生h)、固定残業を控除してから1.25倍。
// 深夜は追加割増分のみ(例:0.25)、休日は全体倍率(例:1.35)を掛ける。
func (c Config) OvertimePayFromAttendance(a MonthlyAttendance) OvertimePayBreakdown {
	chargeable := a.WeekdayOvertime - c.FixedOvertime
	if chargeable < 0 {
		chargeable = 0
	}

	weekdayPay := chargeable * WeekdayOvertimeMultiplier
	lateNightPay := a.LateNightHours * c.OvertimeRates.LateNight
	holidayPay := a.HolidayOvertime * c.OvertimeRates.Holiday

	pay := int((weekdayPay + lateNightPay + holidayPay) * c.HourlyRate)
	return OvertimePayBreakdown{
		WeekdayHours:   a.WeekdayOvertime,
		Chargeable:     chargeable,
		LateNightHours: a.LateNightHours,
		HolidayHours:   a.HolidayOvertime,
		Pay:            pay,
	}
}

func (c Config) EstimateGrossFromAttendance(a MonthlyAttendance) int {
	bd := c.OvertimePayFromAttendance(a)
	return c.BaseSalary + c.FixedOverSalary + c.FixedBenefits + bd.Pay
}
