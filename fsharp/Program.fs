module Program

open System.IO
open FSharp.Data

type Sample = JsonProvider<""" {"name":"x"} """>

let private printAttendance (a: Attendance.MonthlyAttendance) =
    let rows =
        [ "A 勤怠月度", a.Period
          "E 雇用区分", a.EmploymentType
          "F 固定残業条件", a.FixedOvertimeCondition
          "G 所定勤務時間", string a.ScheduledHours
          "I 休日出勤日数", string a.HolidayWorkDays
          "J 欠勤日数", string a.AbsenceDays
          "L 残業時間合計(法定外＋法定内)", string a.OvertimeTotalHours
          "M 法定外残業時間", string a.OvertimeStatutoryHours
          "O 法定外残業時間(見込み)/フレックス", string a.OvertimeProjectedHours
          "P 控除時間・時", string a.DeductionHours
          "Q 深夜時間合計", string a.LateNightHours
          "V 残業時間(換算)/通常勤務", string a.OvertimeConvertedRegular
          "W 残業時間(換算)/フレックス", string a.OvertimeConvertedFlex
          "X 法定内所定外平日フレックス残業時間", string a.WeekdayFlexInStatutoryOvertime
          "Y 法定外平日フレックス残業時間", string a.WeekdayFlexOverStatutoryOvertime
          "Z 休日フレックス残業時間", string a.HolidayFlexOvertime ]

    let labelWidth =
        rows |> List.map (fun (k, _) -> k.Length) |> List.max

    let sep = String.replicate (labelWidth + 20) "-"
    printfn "%s" sep
    for k, v in rows do
        printfn "%-*s | %s" labelWidth k v
    printfn "%s" sep
    printfn "整合性: L = Y + Z  -> %M = %M (差 %M)"
        a.OvertimeTotalHours
        (a.WeekdayFlexOverStatutoryOvertime + a.HolidayFlexOvertime)
        (a.OvertimeTotalHours - (a.WeekdayFlexOverStatutoryOvertime + a.HolidayFlexOvertime))
    printfn "整合性: V = X + Y + Z  -> %M = %M (差 %M)"
        a.OvertimeConvertedRegular
        (a.WeekdayFlexInStatutoryOvertime + a.WeekdayFlexOverStatutoryOvertime + a.HolidayFlexOvertime)
        (a.OvertimeConvertedRegular
         - (a.WeekdayFlexInStatutoryOvertime + a.WeekdayFlexOverStatutoryOvertime + a.HolidayFlexOvertime))

[<EntryPoint>]
let main argv =
    match argv with
    | [| "--import"; path |] ->
        Attendance.load path |> printAttendance
        0
    | _ ->
        Path.Combine(__SOURCE_DIRECTORY__, "testdata", "sample.json")
        |> Sample.Load
        |> _.Name
        |> printfn "%s"

        SocialInsurance.printGrades 100000m
        0
