module Io.Prompt

// 対話入力ヘルパ。System.Console を直に叩く副作用層。
//
// 提供する関数:
//   askYesNo  : prompt: string -> bool
//     "y"/"n" を読むまでループ。
//
//   askDecimal : prompt: string -> defaultValue: decimal option -> decimal
//     defaultValue がある場合、空入力(Enter)で採用。
//     パース失敗時は再プロンプト。
//
//   askInt    : prompt: string -> defaultValue: int option -> int
//     同上、int版。
//
// プロンプト表示例:
//   "[4月] 実績で入力しますか？ (y=実績/n=見込み): "
//   "総支給額 [候補: ¥350,000]: "
