module Attendance

open ClosedXML.Excel

type MonthlyAttendance =
    { Period: string
      EmploymentType: string
      FixedOvertimeCondition: string
      ScheduledHours: decimal
      HolidayWorkDays: decimal
      AbsenceDays: decimal
      OvertimeTotalHours: decimal
      OvertimeStatutoryHours: decimal
      OvertimeProjectedHours: decimal
      DeductionHours: decimal
      LateNightHours: decimal
      OvertimeConvertedRegular: decimal
      OvertimeConvertedFlex: decimal
      WeekdayFlexInStatutoryOvertime: decimal
      WeekdayFlexOverStatutoryOvertime: decimal
      HolidayFlexOvertime: decimal }

let private expectedHeaders =
    [ "A1", "勤怠月度"
      "E1", "雇用区分"
      "F1", "固定残業条件"
      "G1", "所定勤務時間"
      "I1", "休日出勤日数"
      "J1", "欠勤日数"
      "L1", "残業時間合計(法定外＋法定内)"
      "M1", "法定外残業時間"
      "O1", "法定外残業時間(見込み)/フレックス"
      "P1", "控除時間・時"
      "Q1", "深夜時間合計"
      "V1", "残業時間(換算)/通常勤務"
      "W1", "残業時間(換算)/フレックス"
      "X1", "法定内所定外平日フレックス残業時間"
      "Y1", "法定外平日フレックス残業時間"
      "Z1", "休日フレックス残業時間" ]

let load (path: string) : MonthlyAttendance =
    use workbook = new XLWorkbook(path)
    let sheet = workbook.Worksheet(1)

    // ヘッダ検証: 想定外フォーマットなら停止
    for cell, expected in expectedHeaders do
        let actual = sheet.Cell(cell).GetString()
        if actual <> expected then
            failwithf "Unexpected header at %s: expected '%s' but got '%s'" cell expected actual

    let str (addr: string) = sheet.Cell(addr).GetString()
    let dec (addr: string) = sheet.Cell(addr).GetValue<decimal>()

    { Period = str "A2"
      EmploymentType = str "E2"
      FixedOvertimeCondition = str "F2"
      ScheduledHours = dec "G2"
      HolidayWorkDays = dec "I2"
      AbsenceDays = dec "J2"
      OvertimeTotalHours = dec "L2"
      OvertimeStatutoryHours = dec "M2"
      OvertimeProjectedHours = dec "O2"
      DeductionHours = dec "P2"
      LateNightHours = dec "Q2"
      OvertimeConvertedRegular = dec "V2"
      OvertimeConvertedFlex = dec "W2"
      WeekdayFlexInStatutoryOvertime = dec "X2"
      WeekdayFlexOverStatutoryOvertime = dec "Y2"
      HolidayFlexOvertime = dec "Z2" }
