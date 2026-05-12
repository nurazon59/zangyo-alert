module Io.ConfigLoader

// configファイル(JSON)を読み込み Domain.Types.Config に詰め替える。
// JsonProvider はサンプルJSONから型を生成するので、testdata/sample.json を更新しておく。
//
// 提供する関数:
//   load : path: string -> Config
//     失敗時は例外をそのまま投げる(CLI入口で拾う方針)
//
// JsonProvider のサンプル形式 (requirement.md 準拠):
//   {
//     "BaseSalary": 300000,
//     "FixedOvertime": 30,
//     "FixedOverSalary": 50000,
//     "FixedBenefits": 40000,
//     "Benefits": [
//       { "Name": "住宅支援", "UnitAmount": 10000 },
//       { "Name": "食事補助", "UnitAmount": 500 }
//     ]
//   }
//
// 既存 Config.fs のロジックを引き継ぎつつ、Benefits 配列の詰め替えを追加する
