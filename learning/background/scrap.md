# background — Goが生まれた背景と設計思想

『実用Go言語』を読み始める前に、「なぜGoが作られたか」を一次情報ベースで整理。
言語仕様の取捨選択の理由が分かると、学習中に出会う「なぜこんな書き方なんだ？」が腑に落ちる。

---

## 2026-04-20 調査まとめ

### TL;DR（3行で）

- **2007年9月21日**、Google社内で Robert Griesemer / Rob Pike / Ken Thompson の3人がホワイトボードに目標を書き始めたのが起点。
- 動機は「Googleの巨大C++コードベースのビルドが**45分**かかる」「依存関係の膨張」「マルチコア時代に追いつけない既存言語」への不満。
- 一貫テーマは **"Language Design in the Service of Software Engineering"**（ソフトウェアエンジニアリングのための言語設計）。研究のための言語ではない。

---

### なぜGoが生まれたのか（核心）

#### 当時のGoogleが抱えていた具体的な痛み

Rob Pike の SPLASH 2012 講演「Go at Google」より：

> "When builds are slow, there is time to think. The origin myth for Go states that it was during one of those 45 minute builds that Go was conceived."
> （ビルドが遅いと、考える時間がある。Goの起源神話は、その45分ビルドの最中に構想されたというものだ）

具体的な数字：

- 2007年当時、Google の C++ バイナリ1本のビルドに **45分**
- ソースコード 4.2MB → `#include` 展開後 **8GB** がコンパイラに渡される（**2000倍** の膨張）
- ヘッダファイルが1つのバイナリビルドで何万回も開かれる

#### 既存言語の「困っていたこと」

Go公式FAQから：

> "One had to choose either efficient compilation, efficient execution, or ease of programming; all three were not available in the same mainstream language."
> （効率的なコンパイル・効率的な実行・書きやすさ、メインストリーム言語ではこの3つが同時に成り立たなかった）

| 言語 | 当時のGoogleでの問題 |
|---|---|
| **C++** | ビルド遅い、ヘッダ依存爆発、並行処理が言語レベルで弱い |
| **Java** | 起動遅い、冗長、並行処理が難しい |
| **Python** | 実行が遅い、型がなく大規模化で破綻しやすい |

→ 「**C++くらい速く、Pythonくらい書きやすく、Javaくらいスケールする**」言語が欲しかった。

#### 「ソフトウェアエンジニアリング」の問題リスト

Pike が挙げた、Goが解決を狙った課題：

- ビルドが遅い
- 依存が制御できない
- プログラマごとに使う言語サブセットが違う
- プログラム理解が難しい
- 重複作業
- アップデートとバージョン乖離のコスト
- 自動化ツールが書きにくい
- クロス言語ビルドの複雑さ

**ポイント**：これは「言語機能」ではなく「**チームで大規模に開発するときの困りごと**」。Goは最初から"個人の生産性"より"組織の生産性"を見ている。

---

### 3人の設計者と持ち寄ったもの

| 人物 | バックグラウンド | Goに持ち込んだもの |
|---|---|---|
| **Ken Thompson** | UNIX / B言語 / C言語 / Plan 9 の設計者。チューリング賞受賞 | 言語の素朴さ、UNIX的な道具志向 |
| **Rob Pike** | Plan 9 / UTF-8共同設計者 / Newsqueak / Limbo（並行処理言語） | CSPベースの並行処理モデル（goroutine, channel） |
| **Robert Griesemer** | V8 JavaScriptエンジン / Java HotSpot のオプティマイザ | 言語設計とコンパイラの実装力 |

> "We were not designing the language by committee."
> （委員会で設計していない）

3人が全員納得した機能だけを入れた、というのが Go の文法が小さい理由のひとつ。

---

### 影響を受けた言語

公式FAQの一文：

> "Go is mostly in the C family (basic syntax), with significant input from the Pascal/Modula/Oberon family (declarations, packages), plus some ideas from languages inspired by Tony Hoare's CSP, such as Newsqueak and Limbo (concurrency)."

| 系統 | Goに残っているもの |
|---|---|
| **C** | 基本構文、ポインタ、ブロック構造 |
| **Pascal / Modula / Oberon** | 宣言の書き方（`var x int`）、パッケージ、可視性ルール |
| **Plan 9 C** | `#include` を許さない依存モデル |
| **Newsqueak / Limbo（CSP系）** | goroutine と channel |
| **Smalltalk的** | インターフェイスの暗黙実装（duck typing寄り） |

「`int x` ではなく `x int`」という宣言順は、**Pascal由来**で「変数名が主役」という思想。

---

### 設計思想（学習中に思い出すべき指針）

#### 1. **less is exponentially more（少ないほど指数的に多くを得る）**

Pike のブログタイトル。機能を足すと言語の複雑さが**乗算的に**増えていく。だから入れない。

#### 2. **キーワードは25個だけ**

C99: 37個、C++11: 84個以上 → Go: **25個**。文法は型情報なしでパースできるほど単純。

#### 3. **明示性 > 暗黙性**

- 未使用の import / 変数は **コンパイルエラー**（警告じゃなく）
- エラーは戻り値で返す（例外なし）
- 型変換は明示（`int(x)`）

> "Explicit error checking forces the programmer to think about errors—and deal with them—when they arise."

#### 4. **継承ではなく合成（composition over inheritance）**

クラス継承は無い。インターフェイスは**暗黙実装**（implements キーワード不要）。

> "Type hierarchies result in brittle code. The hierarchy must be designed early..."

#### 5. **並行処理は言語の一級市民**

`go func()` で goroutine、`chan` でメッセージパッシング。マルチコア前提の設計。

#### 6. **大文字始まり = エクスポート**

`Reader` は public、`reader` は package private。**識別子そのものに可視性が書かれている**ので、grep するだけでAPI境界がわかる。

#### 7. **ツールが書きやすいことを言語設計で担保**

`gofmt`（自動整形）、`go doc`、`go fix`、`go test`、`go mod` … 言語仕様自体がツールフレンドリー。

> "gofmt is often cited by users as one of Go's best features even though it is not part of the language."

#### 8. **GCあり、でもメモリレイアウトは握れる**

Java と違って、struct のフィールドは**インラインに配置**できる。`buf [256]byte` を入れたら追加アロケーションは起きない。

---

### 主要バージョンの進化（ざっくり年表）

| 年 | 出来事 |
|---|---|
| 2007.09 | 3人がホワイトボードで設計開始 |
| 2009.11 | Google が公式に Go を発表（オープンソース） |
| 2012.03 | **Go 1.0** リリース（後方互換の約束ここから） |
| 2018.08 | Go 1.11 で `go modules`（依存管理の刷新） |
| 2022.03 | Go 1.18 で **ジェネリクス** 導入 |
| 2024.08 | Go 1.23（range over func など） |
| 2026.04 | 現在の最新は Go 1.26.2 |

---

### この背景を知ると、学習中に何が変わるか

『実用Go言語』を読むときに、以下の "違和感" を感じても **それは設計通り** と分かる：

- 「例外がない」 → 例外はエラーを隠して握りつぶす危険があるから明示的にした
- 「三項演算子がない」 → `if-else` で書けばいいから機能を増やさない（less is more）
- 「クラスがない」 → 継承の脆さを避けるため、合成とインターフェイスで構成
- 「未使用変数がエラー」 → 大規模コードを clean に保つための硬い意志
- 「ジェネリクスがずっとなかった」 → シンプルさを優先、本当に必要と判断するまで10年我慢した
- 「`gofmt` がスタイルを強制」 → スタイル論争の時間を消すため

→ Goは「**何を入れるか**」より「**何を入れないか**」で設計された言語。
→ なので学習中は「機能が足りない！」と感じる場面で、**"なぜ無いのか"** を考えると思想が腹落ちする。

---

### 一次情報・参考リンク

- [Go at Google: Language Design in the Service of Software Engineering](https://go.dev/talks/2012/splash.article) — Rob Pike, SPLASH 2012（必読）
- [Go FAQ](https://go.dev/doc/faq) — 公式FAQ
- [The first Go program](https://go.dev/blog/first-go-program) — Goで書かれた最初のプログラム
- [Half a decade with Go](https://go.dev/blog/5years) — 5年間の振り返り
- [Less is exponentially more](https://commandcenter.blogspot.com/2012/06/less-is-exponentially-more.html) — Rob Pike のブログ
- [Go: A Documentary](https://golang.design/history/) — コミュニティ製の年表
- [Wikipedia: Go (programming language)](https://en.wikipedia.org/wiki/Go_(programming_language))

---

## 2026-04-20 09:52 Python → Go 移行の視点で優しく読み直す

上の調査まとめは情報密度が高いので、**物語として** もう一度たどり直す。会社が Python → Go に移行している背景を理解する切り口で整理する。

### 🏢 2007年、Google社内で起きていたこと

Google のエンジニアが、コードを1行直しただけで **ビルド完了まで45分** 待たされる世界。コーヒーを淹れて、トイレに行って、メールを返しても、まだ終わらない。

3人のベテラン — **Ken Thompson**（C言語と UNIX を作った人）、**Rob Pike**、**Robert Griesemer** — が、このビルドを待ちながら考えた。

> 「この言語、そもそも設計が古くない？もっとマシな言語、自分たちで作れないかな」

これが Go の始まり。**「学者が研究のために作った」言語ではなく、「現場のエンジニアが困って作った」言語** という点が、他の言語との大きな違い。

### 😣 当時の3大言語、それぞれの痛み

| 言語 | 書きやすい？ | 速い？ | 大規模に耐える？ |
|---|---|---|---|
| **C++** | ❌ 難しい | ✅ 速い | ⚠️ ビルドが地獄 |
| **Java** | ⚠️ 冗長 | ⚠️ 起動遅い | ✅ 耐える |
| **Python** | ✅ 書きやすい | ❌ 遅い | ❌ 型がなくて破綻 |

Pike の有名な一言：

> 「効率的なコンパイル・効率的な実行・書きやすさ — どれか2つは選べるけど、3つ同時に満たす言語がなかった」

**Go はこの3つを全部欲しがった** 言語。

### 🐍 Python → Go 移行が起きる理由（ここが本題）

会社が Python から Go に移っているのは、Google が2007年に感じた痛みと同じ構造。

#### Python の「書きやすさ」は最高。でも…

- **① 型がない** → プロジェクトが10万行を超えると「この変数、何の型だっけ？」で迷子。mypy を後付けしても限界がある
- **② 遅い** → Web サーバーや大量データ処理で CPU がボトルネック
- **③ デプロイが複雑** → Python のバージョン、venv、依存ライブラリ…本番で動かすまでが大変
- **④ 並行処理が弱い** → GIL（1スレッドしか動けない制約）でマルチコアを活かしにくい

#### Go がそれに対して出した答え

- **① 静的型付け** → コンパイラが間違いを教えてくれる（書いた瞬間にエラー）
- **② 速い** → C++ に近い速度で動く
- **③ 1ファイルバイナリ** → `go build` すると実行ファイル1個ができて、サーバに置くだけで動く。環境構築地獄なし
- **④ goroutine** → `go func()` と書くだけで並行処理。GIL なし

**BigQuery / dbt で扱うような「大量データ × 並行処理 × 長くメンテする」世界は、実は Go が一番得意な領域。** Google がまさにそこで生まれたから。Python が「書いて捨てるスクリプト」で最強なのと対照的に、Go は「何年も動かす業務システム」向けに最適化されている、と捉えると移行の意図が腑に落ちる。

### 🎨 Go の「優しい頑固さ」— 設計思想

Go を学んでいて「なんでこんな書き方なの？」と思う瞬間のほとんどは、**意図的な不便さ**。

#### 例1: 未使用の変数がエラーになる

```go
x := 10   // これだけ書いて使わないと、コンパイル自体が失敗する
```

他の言語だと「警告」だが、Go は**エラー**。理由：

> 「大規模コードベースを綺麗に保ちたいから、妥協しない」

#### 例2: 例外（try/catch）がない

Python だとこう：
```python
try:
    do_something()
except Exception:
    pass   # ← これで例外を握りつぶせてしまう
```

Go はこう：
```go
result, err := do_something()
if err != nil {
    return err   // ← 必ずエラーを意識する
}
```

> 「エラーを握りつぶせる言語より、毎回考えさせる言語の方が安全」

#### 例3: クラス継承がない

Python だと `class Dog(Animal):` と書けるが、Go には継承が**ない**。代わりに「合成（composition）」を使う。

> 「継承は最初の設計が間違うと全部やり直しになる。だから最初から入れない」

### 📏 究極の一言：「less is exponentially more」

Pike のブログタイトル。**「機能が少ないほど、得られるものは指数的に増える」**。

- C++ のキーワード: **84個以上**
- Go のキーワード: **たった25個**

少ないから：
- 1週間で文法を覚えきれる
- チームメンバー全員が同じ書き方をする
- `gofmt` が整形を強制するので、コードレビューで「スタイル論争」が消える

Python で「PEP8/black/isort…」と揉めた議論が、Go では**最初から存在しない**。

### 🧭 学習中、心の支えにする3つの姿勢

1. **機能が無いのは忘れたからじゃなく、入れなかったから**
   → 「なぜ無いのか？」を考えると設計思想が見える
2. **Go は個人の生産性より、チームの生産性を優先する言語**
   → 書く時の手間より、読む時の明快さを選んでいる
3. **Python で快適だった書き方が Go では非推奨になることがある**
   → それは劣化ではなく、別の軸で最適化されているだけ

**「Python は1人で素早く書くのに最強」「Go はチームで長く保守するのに最強」と割り切ると、移行の悩みが減る。** 両方向の優劣ではなく、"目的が違う道具" として捉えるのが健全。BigQuery/dbt の知識は、Go でデータパイプラインや ETL を書くときにそのまま活きる。
