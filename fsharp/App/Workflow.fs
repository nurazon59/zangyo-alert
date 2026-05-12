module App.Workflow

// 対話フロー全体。Pure core を順に呼ぶ impure shell。
//
// 型:
//   Args = { ConfigPath: string; ImportXlsxPath: string option }
//     ImportXlsxPath は今回スコープ外。受け取るだけで未使用
//
// 提供する関数:
//   run : Args -> int
//     プロセス終了コードを返す (成功時 0)
//
// 処理の流れ:
//   1. config       <- ConfigLoader.load args.ConfigPath
//   2. gradeTable   <- 等級表CSVをロード (Io層、Domain.Grades 内のloaderか別ファイル)
//   3. for month in [4; 5; 6]:
//        a. Prompt.askYesNo で 実績/見込み を聞く
//        b. 実績なら Prompt.askDecimal で総支給額、見込みなら追加残業時間
//        c. config.Benefits 各項目について Prompt.askInt で利用回数
//        d. MonthlyEntry を組み立てて蓄積
//   4. 各月の monthlyGross を Salary で計算 -> average3 で平均
//   5. judgement <- Grades.judge gradeTable avg
//   6. alert     <- Alert.compute config judgement entries
//   7. printfn "%s" (Renderer.renderResult entries judgement alert)
