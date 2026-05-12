module Domain.Types

// ドメイン全体で共有する型を定義する。
// このモジュールは値の計算は持たず、データの形だけを表現する。
//
// 定義するもの:
//   - Benefit          : 福利厚生の定義 (Name, UnitAmount)。configから来る
//   - MonthInput       : DU。Actual(総支給額) | Estimate(追加残業時間)
//   - BenefitUsage     : Benefit と利用回数のペア
//   - MonthlyEntry     : 1ヶ月分の入力一式 (Month, Input, BenefitUsages)
//   - Config           : 設定ファイル全体を表すRecord
//   - GradeJudgement   : 等級判定結果 (平均/標準報酬月額/健保等級/厚年等級/次閾値)
//   - Alert            : アラート結果 (あと何円/何時間/何回)
