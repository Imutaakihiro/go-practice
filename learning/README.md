# Go学習ノート

オライリー『実用Go言語』を読みながら、学んだことをZennスクラップ風に記録していくリポジトリ。

## 運用ルール

- トピックごとにディレクトリを切る（`basics/`, `concurrency/` など）
- 各ディレクトリに `scrap.md`（学習ログ）と `examples/`（Goコード）を置く
- `scrap.md` は時系列で追記する（`## YYYY-MM-DD HH:MM 見出し`）
- 疑問 → 調査 → 解決 の流れを残し、後で振り返れるようにする

## トピック一覧

| ディレクトリ | テーマ | ステータス |
|---|---|---|
| [background/](./background/scrap.md) | Goが生まれた背景・設計思想 | 調査済み |
| [basics/](./basics/scrap.md) | 基本文法・型・パッケージ | 未着手 |
| [lunch-bot/](./lunch-bot/scrap.md) | Slackランチシャッフルbotを作って分解 | 進行中 |

新しいトピックを始めたらここに追記する。

## Go環境

- Go 1.26.2 (darwin/arm64)
- モジュールルート: リポジトリ直下の `go.mod`
- コード実行: `go run ./learning/<topic>/examples/xxx.go`
