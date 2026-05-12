module Io.Renderer

// 結果整形のみ。出力(printfn)は呼び出し側(Workflow)で行うため、ここは string を返す。
//
// 提供する関数:
//   renderResult :
//       entries: MonthlyEntry list ->
//       judgement: GradeJudgement ->
//       alert: Alert ->
//       string
//
// 出力レイアウト (requirement.md 準拠):
//   [算定]  4月: 実績 ¥XXX,XXX  5月: 見込 ¥XXX,XXX  6月: 見込 ¥XXX,XXX
//   [平均]  ¥XXX,XXX
//   [等級]  健保 N級 / 厚年 M級   標準報酬月額 ¥XXX,000
//
//   [アラート]
//     - 次の等級まであと ¥X,XXX
//     - 残り月で許容できる追加残業: 約 N 時間
//     - 住宅支援: あと N 回まで利用可
//     - 食事補助: あと N 回まで利用可
//
// 金額は3桁区切り、"実績"/"見込" はラベル
