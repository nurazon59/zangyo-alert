module Domain.Salary

// 月額の総支給を計算する純粋関数群。副作用なし。
//
// 提供する関数:
//   hourlyOvertimeRate : Config -> decimal
//     固定残業代 / 固定残業時間 で時給換算
//
//   benefitTotal : BenefitUsage list -> decimal
//     その月の福利厚生利用合計額
//
//   monthlyGross : Config -> MonthlyEntry -> decimal
//     その月の総支給額。MonthInput を match で分岐:
//       Actual amount        -> amount + benefitTotal
//       Estimate extraHours  -> BaseSalary + FixedOverSalary + FixedBenefits
//                               + extraHours * hourlyOvertimeRate
//                               + benefitTotal
//
//   average3 : decimal -> decimal -> decimal -> decimal
//     4/5/6 月の平均
