# zangyo alert

残業時間によって社会保険の等級（標準報酬月額）が上がらないように、4・5・6月の報酬実績から等級を試算し許容ラインをアラートする CLI。

## 何が分かるか

- 4・5・6月の平均報酬から定時決定の等級を試算
- 次の等級に上がるまでの **金額・許容残業時間・福利厚生の利用回数** を算出
- 勤怠 xlsx の残業時間合計から見込み総支給額を自動算出して候補値表示

## インストール

```sh
cd go
go build -o ../bin/zangyo-alert .
```

## 使い方

```sh
zangyo-alert --config ./config.json
zangyo-alert --config ./config.json --import ./勤怠実績.xlsx
zangyo-alert --config ./config.json --grades ./data/social-insurance-tokyo-2026.csv
```

| フラグ | 必須 | 説明 |
|---|---|---|
| `--config` | ◯ | 固定値（基本給・固定残業代・福利厚生定義）の JSON |
| `--import` | | 勤怠 xlsx。残業時間合計から見込み総支給額を計算し候補値として表示 |
| `--grades` | | 等級表 CSV（既定: `data/social-insurance-tokyo-2026.csv`） |

### config.json の例

```json
{
  "BaseSalary": 300000,
  "FixedOvertime": 30,
  "FixedOverSalary": 50000,
  "FixedBenefits": 40000,
  "Benefits": [
    { "Name": "住宅支援", "UnitAmount": 10000 },
    { "Name": "食事補助", "UnitAmount": 500 }
  ]
}
```

- `BaseSalary` 月額固定給
- `FixedOvertime` 固定残業時間（h）
- `FixedOverSalary` 固定残業代
- `FixedBenefits` 固定支給（住宅・通勤等）
- `Benefits` 利用回数で給与に加算される福利厚生（`UnitAmount` 円/回）

## 動き方

```
[起動]
   │
   ├─ config.json 読込
   ├─ 等級表 CSV 読込
   ├─ --import xlsx → 該当月の見込み総支給額を算出
   │
   ▼
[対話入力 4月 → 5月 → 6月]
   │  実績で入力？ (y/n)
   │   ├─ y: 総支給額  ← import済みなら候補値が[ ]で表示、Enterで採用
   │   └─ n: 追加残業時間(h) → 固定値から自動計算
   │  福利厚生: 各項目の利用回数を入力
   │
   ▼
[算定] 4月: 実績 ¥XXX,XXX  5月: 見込 ¥XXX,XXX  6月: 見込 ¥XXX,XXX
[平均] ¥XXX,XXX
[等級] 健保 N級 / 厚年 M級   標準報酬月額 ¥XXX,000

[アラート]
  - 次の等級まであと ¥X,XXX（3ヶ月合計）
  - 残り月で許容できる追加残業: 約 N 時間（時給 ¥YYY 換算）
  - 住宅支援: あと N 回まで利用可
  - 食事補助: あと N 回まで利用可
```

## ディレクトリ構成

```
go/
├── main.go              CLI エントリ
├── app/workflow.go      対話フロー
├── domain/              純粋ロジック（給与・等級・アラート計算）
├── iolayer/             I/O（config/CSV/xlsx/プロンプト/レンダラ）
└── data/                等級表 CSV
```

## スコープ

- 東京都・協会けんぽ 2026年度料率を使用
- 介護保険料・賞与・遡及支給・産育休等の特殊ケースは対象外
- 給与計算・年末調整・所得税は対象外

## ライセンス

[LICENSE](./LICENSE) を参照。
