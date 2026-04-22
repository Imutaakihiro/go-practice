# basics — 学習プラン

『実用Go言語 第2版』を読みながら、**ローカルで動く lunch-bot アプリを自分ですべて紐解く** ことを直近ゴールに据えた学習プラン。その先に Go エンジニアとして自立できる基礎体力を積み上げる。

## ゴール設定（3段階）

### 🎯 直近ゴール（このplan.mdのスコープ）

**ローカルで動く lunch-bot アプリを、自分の手ですべて紐解く**

- `learning/lunch-bot/` の全ファイル（`config.go`, `slack_client.go`, `lunch_service.go`, `shuffler.go`, `scheduler.go`, `lunch_handler.go`, `cmd/bot/main.go`）を、他人のコードではなく **自分のコードとして** 読める状態
- 各行について「なぜこう書くか」を自分の言葉で説明できる状態
- 一部を消して書き直しても、元の動作を再現できる状態

### 🏔 中間ゴール

**エンジニアの基礎となる考え方 + Go 言語の理解を深める**

- 型・抽象化・エラー処理・並行性・設計の哲学など、**言語を超えて通用する基礎体力**
- Go の標準ライブラリを武器として選べる・組み合わせられる
- 技術を「自分ごと」として学ぶ姿勢（vibe コーディング期からの完全脱却）

各ループでは「Go の設計者がなぜこう設計したか」「他言語ならどう書くか」の視点を意識する。

### 🚀 最終ゴール

**Go エンジニアになる**

- 要件から自分で設計・実装・デプロイできる
- OSS の Go コードを読み解き、自分のプロジェクトに取り入れられる
- エラー・バグを自力で推論して直せる
- interface / goroutine / エラー設計など、Go 特有の判断を状況に応じて下せる

## 進め方（1ループの構造）

各ループは以下の4ステップで進める。1ループあたり目安 **1.5時間**。

1. **読む**（〜30分） — 本の該当節を読む
2. **書く**（〜30分） — `examples/` に自分の手でコードを書く（写経ではなく、概念を使った例を自分で考える）
3. **polarisで探す**（〜15分） — polaris リポジトリで同じ概念が使われている箇所を読む
4. **まとめる**（〜15分） — `scrap.md` に自分の言葉で「わかったこと・疑問・はまったこと」を書く

### 大事な原則

- **網羅しない**: 章を全部読んでから手を動かすのは禁止。1ループで1単位を完結させる
- **AIに聞く前に5分自分で考える**: コンパイルエラーは英語のメッセージを自分でまず読む
- **穴だらけのまま動かす**: 完璧主義より、わからないまま動かして気づいた穴を埋める方が伸びる

## 本との対応関係

`plan.md` は『実用Go言語 第2版』(O'Reilly 978-4-8144-0136-9) の構成に沿っている。

- **Loop 1〜10**: 付録A「駆け足で学ぶGoの基礎」（A.1〜A.15）と一対一対応 — 本の下地編
- **Loop 11〜15**: 本編の該当章にジャンプ — テーマ別の深掘り

本編は分量が多いので、各Loopでは **「本の該当節だけピンポイント抜き読み」** する。章を全部読まない。lunch-bot に出てくる概念を優先し、それ以外は後回しで良い。

## ループ一覧

進捗チェックボックスを更新しながら進める。本の該当節は「📖」で示す。

### Part 1: 付録A 駆け足章（下地編）

- [x] **Loop 1** — Hello World + リテラル + 変数宣言 + 名前 + コメント（+ シャドーイング・カプセル化・init もカバー）
  - 📖 付録A.1〜A.4
  - ねらい: Goプログラムの見た目・命名規約・Goらしい書き方に慣れる
  - 成果物: `examples/01.go`, `examples/02_introduction.go`
- [x] **Loop 2** — 型と変換 + ゼロ値
  - 📖 付録A.5, A.7
  - ねらい: Goの型システムの哲学を理解する（暗黙の型変換をしない設計）
  - 成果物: `examples/03_bmi.go`
- [x] **Loop 3** — ポインター
  - 📖 付録A.6
  - ねらい: メモリと参照の感覚を掴む（多くの学習者が詰まる山場）
  - 成果物: `examples/04_savings.go`
- [x] **Loop 4** — スライス + マップ
  - 📖 付録A.8, A.9
  - ねらい: コレクション操作（SQLのSELECT/JOINとの類比で理解できる）
  - 成果物: `examples/05_inventory.go`
- [x] **Loop 5** — 制御構文
  - 📖 付録A.10
  - ねらい: if / for / switch のGo流儀に慣れる
  - 成果物: `examples/06_grade.go`
- [x] **Loop 6** — 関数
  - 📖 付録A.11
  - ねらい: 複数戻り値・defer など他言語と異なる設計を体験する
  - 成果物: `examples/07_stats.go`
- [ ] **Loop 7** — エラー処理
  - 📖 付録A.12（下地）+ 本編 **6.1.2「fmt.Errorf を使ったフォーマット付きエラー」** / **6.2.1「エラーのラップ」** / **6.2.2「エラーのアンラップ」**
  - ねらい: Goの哲学が最も色濃く出る部分。例外を持たない理由を体感する
  - lunch-bot連携: `config/config.go`, `repository/slack_client.go` の `if err != nil { return fmt.Errorf("...: %w", err) }` を実コード題材に
- [x] **Loop 8** — 構造体
  - 📖 付録A.13
  - ねらい: データと振る舞いの分離（OOPとは違う設計思想を理解）
  - 成果物: `examples/09_profile.go`
- [ ] **Loop 9** — ライブラリのインポート
  - 📖 付録A.14 + 本編 **7.1「プロジェクト構成の事前知識」**
  - ねらい: パッケージ管理・モジュールの仕組み
  - lunch-bot連携: `go.mod` のモジュール名と `import "github.com/androots/..."` の対応を読む

### Part 2: 本編ジャンプ（テーマ別深掘り）

- [ ] **Loop 11** — メソッド + ファクトリー関数
  - 📖 本編 **3.2.1「ファクトリー関数」** + **3.3.1「値レシーバーとポインターレシーバーのどちらを使えば良いか」**
  - ねらい: 構造体に「動詞」を生やす。`NewXxx()` パターンを理解する。Loop 3（ポインタ）と Loop 8（構造体）の統合
  - lunch-bot連携: `slack_client.go` の `NewSlackClient()` / `(c *SlackClient) PostMessage(...)`、`lunch_service.go` の `NewLunchService()` を実コード題材に
- [ ] **Loop 12** — インタフェース
  - 📖 本編 **4.1「柔軟なコードを書くインタフェースの利用法」** のみ（4.2 以降は当面スキップ）
  - ねらい: Goの設計の山場。「契約」と「依存性注入」の感覚を掴む
  - lunch-bot連携: `lunch_service.go` の `SlackRepository` インタフェースが最大の教材
- [ ] **Loop 13** — time パッケージとタイムゾーン
  - 📖 本編 **1.10「日時の取り扱い」** の 1.10.1, 1.10.2
  - ねらい: `time.Time` と `time.Duration` の違い、JSTの扱い方を身につける
  - lunch-bot連携: `scheduler.go` の `NextMorningAt` / `SleepUntil` を読む
- [ ] **Loop 14** — HTTPクライアント + JSON
  - 📖 本編 **12.1「net/httpを使ったHTTPクライアントの基本」** + **9.1.1「JSON 基本的なエンコードとデコード」**
  - ねらい: 外の世界とやりとりする定番セット。struct tag の役割を理解
  - lunch-bot連携: `slack_client.go` 全体を読み解く日
- [ ] **Loop 10** — ウェブアプリケーション（サーバー側）
  - 📖 付録A.15 → 余裕があれば本編 **11.2「ウェブプログラミングの基本」**
  - ねらい: HTTPサーバーの骨格を書く。Loop 14（クライアント）と対になる概念
  - 成果物: `examples/10_hello_web.go` (既存) を読み直す + 本編パターンで書き直し
- [ ] **Loop 15** — goroutine と channel（並行処理）
  - 📖 本編 **16.1「並行処理の基本を知る」** のみ（16.2以降は当面スキップ）
  - ねらい: Goが一番得意な領域。lunch-bot の "09:00まで待つ" を並行処理で書き直す
  - lunch-bot連携: Phase 2 の目標「並行処理の理解」に直結

## 10日走行スケジュール（実行フェーズ）

### 📍 現在地

- **開始日**: 2026-04-22（水）= Day 1
- **現在の Day**: Day 1（進行中 🟡）
- **目標完了**: Day 10 時点で「lunch-bot 全ファイルを自分の言葉で説明できる」状態に到達
- **バッファ**: 稼働外 5 日（遅れ吸収 / Phase 3 前倒し用）

> 🌅 毎朝「今日の進めるのは？」と聞いたら、plan.md のこのセクションの該当 Day を参照して答える運用。

### スケジュール全体表

| Day | Loop | テーマ | 想定時間 | 状態 |
|---|---|---|---|---|
| **Day 1** | Loop 7 | エラー処理 | 6h | 🟡 進行中 |
| Day 2 | Loop 9 | import / モジュール | 5h | ⬜ 未着手 |
| Day 3 | Loop 11 | メソッド + ファクトリー関数 | 6h | ⬜ |
| Day 4 | Loop 12 前半 | interface 概念編 | 6h | ⬜ |
| Day 5 | Loop 12 後半 | interface 実コード編 | 6h | ⬜ |
| Day 6 | Loop 13 | time パッケージ | 5h | ⬜ |
| Day 7 | Loop 14 前半 | HTTP クライアント | 6h | ⬜ |
| Day 8 | Loop 14 後半 | JSON + struct tag | 6h | ⬜ |
| Day 9 | Loop 15 | goroutine + channel | 6h | ⬜ |
| Day 10 | 統合 | 全ファイル読み通し | 5h | ⬜ |

**合計 57h / 予備 3h（+ 外側バッファ 5 日 = 30h）**

### 毎日の運用

1. **朝**: plan.md の該当 Day を開いて、今日の4ステップ（読む/読解/書く/まとめる）を確認
2. **日中**: タスクを順に消化。詰まったら仮説ベース質問で抜ける（5分/15分ルール）
3. **夜（15分）**: scrap.md に「予定通り / 遅れ / 早い」を1行記録 + 翌日 Day を確認

---

### Day 1: Loop 7 — エラー処理

**📖 読む（目安 90分）**
- 付録A.12「エラー処理」
- 本編 6.1.2「fmt.Errorf を使ったフォーマット付きエラー」
- 本編 6.2.1「エラーのラップ」
- 本編 6.2.2「エラーのアンラップ」

**🔍 実コード読解（目安 90分）**
- `learning/lunch-bot/internal/config/config.go`（27行全体）— `Load()` が 2 箇所でエラーを返すパターン
- `learning/lunch-bot/internal/repository/slack_client.go` の `PostMessage` 関数 — `%w` で何重にラップされるかを追う

**✍️ 書く（目安 90分）**
- `learning/basics/examples/11_errors.go` を新規作成
- 自作関数で `%w` ラップしたエラーを作り、呼び出し側で `errors.Is` / `errors.As` で取り出す最小例

**📝 まとめる（目安 60分）**
- `learning/basics/notes/07_errors.md` を新規作成
- 「Go のエラーは値である」「なぜ例外を持たないか」「`%w` の役割」を自分の言葉で整理

**今日の達成基準**
- [ ] `if err != nil { return fmt.Errorf("xx: %w", err) }` の意味を自分の言葉で説明できる
- [ ] `PostMessage` のエラーパスを1行ずつ追って説明できる
- [ ] `examples/11_errors.go` が動く
- [ ] `notes/07_errors.md` が書けている

---

### Day 2: Loop 9 — import / モジュール

**📖 読む（目安 60分）**
- 付録A.14「ライブラリのインポート」
- 本編 7.1「プロジェクト構成の事前知識」(7.1.1〜7.1.4 を中心に)

**🔍 実コード読解（目安 60分）**
- `go.mod`（リポジトリ直下）— 今日直した `module github.com/androots/go-practice` の意味
- `learning/lunch-bot/cmd/bot/main.go` の `import` ブロック — パッケージパスが `go.mod` の module 名とどう対応するか
- `learning/lunch-bot/internal/` の意味（`internal` は外部から import 禁止の特別ディレクトリ）

**✍️ 書く（目安 90分）**
- `learning/basics/examples/12_import/` 以下に自作パッケージを1個作る（例: `mathutil/add.go` を自作して `main.go` から呼ぶ）

**📝 まとめる（目安 60分）**
- `notes/09_import.md` — 「パッケージ / モジュール / import 解決」の関係を自分の言葉で

**今日の達成基準**
- [ ] `go.mod` の `module XXX` と `import "XXX/yyy"` が繋がっていることを説明できる
- [ ] `internal` ディレクトリの制約を言える
- [ ] 自作パッケージを1つ作って main から呼び出せた

---

### Day 3: Loop 11 — メソッド + ファクトリー関数

**📖 読む（目安 90分）**
- 本編 3.2.1「ファクトリー関数」
- 本編 3.3.1「値レシーバーとポインターレシーバーのどちらを使えば良いか」

**🔍 実コード読解（目安 90分）**
- `slack_client.go` の `NewSlackClient()`（ファクトリー関数のお手本）
- `slack_client.go` の `func (c *SlackClient) PostMessage(...)` と `GetReactionUsers(...)`（ポインタレシーバのメソッド）
- `lunch_service.go` の `NewLunchService()` と `func (s *LunchService) RunSession()`

**✍️ 書く（目安 90分）**
- `examples/13_methods.go` — 自作構造体 `Counter` を作り、`Increment()` / `Value()` メソッドを値レシーバとポインタレシーバ両方で書いて違いを比較

**📝 まとめる（目安 60分）**
- `notes/11_methods.md` — 「なぜ `NewXxx` パターンなのか」「値 vs ポインタレシーバの判断基準」

**今日の達成基準**
- [ ] `NewSlackClient()` がなぜ `*SlackClient` を返すか説明できる
- [ ] ポインタレシーバと値レシーバの使い分け基準を言える
- [ ] 自作 `Counter` で両パターンの挙動の違いを体感した

---

### Day 4: Loop 12 前半 — interface 概念編

**📖 読む（目安 150分）** ← 山場なので多め
- 本編 4.1「柔軟なコードを書くインタフェースの利用法」全体

**🔍 実コード読解（目安 60分）**
- `lunch_service.go` の `SlackRepository` インタフェースの宣言部分だけ
- 「なぜこの interface が必要か？」を自分で仮説を立てる（この時点では答え合わせしない）

**✍️ 書く（目安 90分）**
- `examples/14_interface_basic.go` — `Animal` インタフェースと `Dog` / `Cat` の最小例を自作

**📝 まとめる（目安 30分）**
- `notes/12_interface.md` に **途中まで** 書く（概念編のみ）。実コード連携は Day 5 で加筆

**今日の達成基準**
- [ ] interface が「契約」であることを自分の言葉で言える
- [ ] Go の interface が「暗黙的に満たす」仕組みを体験した
- [ ] 自作 `Animal` で動作確認できた

---

### Day 5: Loop 12 後半 — interface 実コード編

**📖 読む（目安 30分）**
- 付録 4.1 を再読（朝イチで記憶のリフレッシュ）

**🔍 実コード読解（目安 150分）** ← メインイベント
- `lunch_service.go` の `SlackRepository` インタフェース全体
- `repository/slack_client.go` の `*SlackClient` がこの interface を「暗黙的に満たしている」ことを確認
- `cmd/bot/main.go` の DI 配線（`NewLunchService(slackRepo, ...)` でなぜ `*SlackClient` を渡せるか）

**✍️ 書く（目安 90分）**
- `examples/15_interface_di.go` — `SlackRepository` のモック実装（テスト用）を書いて、`LunchService` と同じ形で差し替えられることを確認

**📝 まとめる（目安 30分）**
- `notes/12_interface.md` の続きを書く — 「なぜ DI で interface を挟むのか」

**今日の達成基準**
- [ ] `SlackRepository` と `*SlackClient` の関係を図で説明できる
- [ ] DI のメリット（テスタビリティ / 差し替え可能性）を言える
- [ ] モック実装で同じ `LunchService` を動かせた

---

### Day 6: Loop 13 — time パッケージ

**📖 読む（目安 90分）**
- 本編 1.10.1「日時の time.Time の取得」
- 本編 1.10.2「時間を表す time.Duration」

**🔍 実コード読解（目安 60分）**
- `scheduler.go` 全体（32行）— `NextMorningAt` / `SleepUntil` の実装を1行ずつ

**✍️ 書く（目安 90分）**
- `examples/16_time.go` — JST で「次の月曜 09:00」を計算する関数を自作

**📝 まとめる（目安 30分）**
- `notes/13_time.md` — 「time.Time と Duration の違い」「タイムゾーンの扱い方」

**今日の達成基準**
- [ ] `time.FixedZone("JST", 9*60*60)` が何をしているか言える
- [ ] `NextMorningAt` の分岐（`if !target.After(nowJST)`）の意図を説明できる
- [ ] 自作で「次の○曜日○時」を計算できた

---

### Day 7: Loop 14 前半 — HTTP クライアント

**📖 読む（目安 120分）**
- 本編 12.1「net/http を使った HTTP クライアントの基本」全体
  - 12.1.1 http.Get
  - 12.1.2 http.Post
  - 12.1.3 http.Client を作成してリクエスト

**🔍 実コード読解（目安 120分）**
- `slack_client.go` の `PostMessage` を構造的に読む:
  - ペイロード作成（`json.Marshal`）
  - リクエスト組み立て（`http.NewRequest`）
  - ヘッダ設定（`Authorization`）
  - 送信（`c.httpClient.Do`）
  - `defer resp.Body.Close()` の意味
  - レスポンス読み取り（`io.ReadAll`）

**✍️ 書く（目安 90分）**
- `examples/17_http_client.go` — 公開 API（例: `https://httpbin.org/post`）に JSON を POST して結果を受ける最小例

**📝 まとめる（目安 30分）**
- `notes/14_http_client.md` — 「HTTPリクエストの4段階」を自分の言葉で

**今日の達成基準**
- [ ] `http.NewRequest` → `Do` → `Body.Close` のライフサイクルを説明できる
- [ ] `defer` がどのタイミングで実行されるか言える
- [ ] 自作で公開APIにPOSTして結果を得られた

---

### Day 8: Loop 14 後半 — JSON + struct tag

**📖 読む（目安 90分）**
- 本編 9.1.1「JSON 基本的なエンコードとデコード」
- 本編 3.5「タグを使って構造体にメタデータを埋め込む」（3.5.1〜3.5.2 まで）

**🔍 実コード読解（目安 90分）**
- `slack_client.go` の `GetReactionUsers` と `reactionsGetResponse` 構造体
- ネストされた struct tag（`json:"ok"` / `json:"reactions"` など）が Slack API レスポンスとどう対応するか
- `json.Unmarshal` が struct tag を頼りにどう動くか

**✍️ 書く（目安 90分）**
- `examples/18_json.go` — 自作 struct に `json:"..."` タグを付けて Marshal / Unmarshal

**📝 まとめる（目安 30分）**
- `notes/14_http_client.md` に JSON/tag 部分を追記

**今日の達成基準**
- [ ] `json:"field_name"` が何をしているか言える
- [ ] ネストされた JSON を Go の struct に受ける方法を説明できる
- [ ] `GetReactionUsers` のレスポンス処理を1行ずつ追える

---

### Day 9: Loop 15 — goroutine + channel

**📖 読む（目安 120分）**
- 本編 16.1「並行処理の基本を知る」全体
  - 16.1.1 ゴルーチン
  - 16.1.2 チャネル

**🔍 実コード読解（目安 60分）**
- `scheduler.go` の `time.Sleep(d)` に注目
- 「このブロッキングを goroutine にすると何が変わるか？」を考える

**✍️ 書く（目安 120分）**
- `examples/19_goroutine.go` — 2つの goroutine がチャネルで値を渡し合う最小例
- `examples/20_scheduler_async.go` — lunch-bot の `SleepUntil` を goroutine + `time.AfterFunc` で書き換える実験

**📝 まとめる（目安 30分）**
- `notes/15_goroutine.md` — 「goroutine と channel の役割」「なぜ Go が並行処理を得意とするか」

**今日の達成基準**
- [ ] `go func() { ... }()` の意味を言える
- [ ] チャネルが「同期装置」として働く仕組みを説明できる
- [ ] lunch-bot の `SleepUntil` を goroutine 化した実装を書けた

---

### Day 10: 統合読み通し — lunch-bot 全ファイル

**🔍 全ファイル通し読み（目安 180分）**
順番に開いて、1行ずつ自分の言葉で説明できるか確認:
1. `cmd/bot/main.go`（DI 配線）
2. `internal/config/config.go`
3. `internal/handler/lunch_handler.go`
4. `internal/service/lunch_service.go`
5. `internal/service/shuffler.go`
6. `internal/service/scheduler.go`
7. `internal/repository/slack_client.go`

**📝 最終まとめ（目安 120分）**
- `learning/lunch-bot/scrap.md` に「lunch-bot を紐解いた10日の総括」を追記
- 各ファイルの1行サマリを書く（例: 「`slack_client.go` = 外部の Slack API を叩く窓口。HTTP + JSON の定番パターン」）
- 「もう一度ゼロから書いたら、どの順で書くか？」を自分で設計してみる

**今日の達成基準**
- [ ] 7 つのファイル全部について、他人に3分で説明できる
- [ ] 一部を消して書き直したら元通り動かせると思える自信がある
- [ ] 🎯 **直近ゴール達成！** lunch-bot が自分のコードになった

---

## lunch-bot 紐解き進捗マップ

各 Loop が完走すると、lunch-bot のどこが読めるようになるか:

| ファイル | 読めるようになる Loop |
|---|---|
| `shuffler.go` | Loop 4（スライス）+ Loop 6（関数）← ✅ 既に読める |
| `config/config.go` | Loop 7（エラー処理） |
| `lunch_handler.go` | Loop 8（構造体）+ Loop 11（メソッド）|
| `cmd/bot/main.go` | Loop 9（import）+ Loop 11（メソッド）|
| `scheduler.go` | Loop 13（time パッケージ）|
| `slack_client.go` | Loop 14（HTTP + JSON）|
| `lunch_service.go` | Loop 12（インタフェース）← **山場** |
| "09:00 まで待つ" の並行処理化 | Loop 15（goroutine/channel）|

**Loop 7, 9, 11, 12, 13, 14, 15 の 7 本を倒せば、lunch-bot 全体が自分のコードになる**。これが直近ゴールの具体的な到達点。

## ファイル構成

```
learning/basics/
├── plan.md          ← この計画書（静的ロードマップ）
├── scrap.md         ← 時系列の学習ログ（疑問→調査→解決の流れを残す）
├── notes/           ← 本の内容を自分の言葉でまとめたノート（ファインマン・テクニック）
│   ├── 01_hello.md
│   ├── 02_types.md
│   └── ...
└── examples/        ← 各ループで書いたGoコード
    ├── 01_hello.go
    ├── 02_types.go
    └── ...
```

### 各ファイルの役割の違い

| ファイル | 内容 | いつ書く |
|---|---|---|
| `notes/XX_xxx.md` | 本の内容を自分の言葉で **整理・要約** | 「読む」ステップで本を閉じた後 |
| `examples/XX_xxx.go` | 概念を使って **動かしたコード** | 「書く」ステップ |
| `scrap.md` | 疑問・気づき・はまった話を **時系列** で | ループ中いつでも、思いついたら |
