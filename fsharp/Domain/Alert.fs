module Domain.Alert

// 「あと何円で等級UP」「あと何時間残業可」「Benefit あと何回利用可」を計算する純粋関数。
//
// 提供する関数:
//   compute :
//       config: Config ->
//       judgement: GradeJudgement ->
//       entries: MonthlyEntry list ->
//       Alert
//
// 計算方針 (シンプル版):
//   - AmountUntilNextGrade   = NextThresholdAmount - Average  (Noneなら None)
//     ただし「平均 = (4月+5月+6月)/3」なので、合計に対しては *3 で換算する
//   - AllowedExtraOvertimeHours = AmountUntilNextGrade合計 / hourlyOvertimeRate
//   - AllowedBenefitUsages = config.Benefits 各項目について
//                            AmountUntilNextGrade合計 / UnitAmount を切り捨て
//
// 注意: 4月確定済みなら残月で割る等の制約は本実装で詰める
