Go入門ハンズオン [01:基礎編]
香川Go言語勉強会
8 Sep 2018

Ryuji Iwata
Okayama, Japan
@qt_luigi

* Ryuji IWATA

.image go-intro-hands-on/img/qt_luigi250.png

Google+

.link https://plus.google.com/+RyujiIwata +RyujiIwata

Twitter

.link http://twitter.com/qt_luigi/ @qt_luigi

* 何を作るの？

* 成果物

タイピングゲームを作ろう！q@w@p

  $ ./typing 
  golang
  input : golang
  ○
  hands-on
  input : hand-on
  ×
  in
  input : in
  ○
  kagawa
  input : okayama
  ×

* 基本仕様

表示する単語をスライス（配列）で用意します。

  [ "golang", "hands-on", "in", "kagawa" ]

以下、スライスの要素の数だけ、添字の順番に繰り返します。

1. 単語を一つ表示します。

  golang

2. 単語を入力します。

  input : golang

3. 表示した単語と入力した単語が一致なら○を、不一致なら×を表示します。

  ○

* 書き方を学ぼう！

* hello.go

    // 実行させるならmainを指定
    package main

    // 外部パッケージを指定
    import (
        "fmt"
    )

    // 同一パッケージ内で使用できる変数を定義
    var (
        // 変数msg1は宣言＋初期化（値から推測してstring型）
        msg1 = "Hello"
        // 変数msg2は明示的にstring型で宣言
        msg2 string
    )

    // main()関数：エントリーポイント
    func main() {
        msg2 = "world"
        // Println()関数は、引数が複数だと半角空白文字で区切って標準出力に表示
        fmt.Println(msg1, msg2)
    }

* hello.go

- ファイル拡張子は.goで、BOMなしのUTF-8で記述。
- インデントはタブが推奨。
- 外部から呼び出せる関数名は大文字で始まる。

    // fmtパッケージのPrintln()関数は名前が大文字で始まっているので呼び出せる
    fmt.Println(msg1, msg2)

* コードの実行

初期状態

    $ ls
    hello.go

go run ファイル名

    $ go run hello.go
    Hello world

* スライス

* スライス (1/3)

- 配列への参照（ポインター）を有する型で、要素を変更できる。

    // 変数exampleをstring型スライスで宣言
    var example []string
    fmt.Println("初期値は", example)

    // 値をまとめて代入（}の前に要素がなければ最後の要素にカンマが必要！）
    example = []string{
        "Golang",
        "hands-on",
    }
    fmt.Println("2番目の値は", example[1])

    // 値を追加するにはappend()関数を使用
    example = append(example, "in", "kagawa")
    fmt.Println("追加結果は", example)

実行結果

    初期値は []
    2番目の値は hands-on
    追加結果は [Golang hands-on in kagawa]

* スライス (2/3)

- 取得したい要素の位置を開始する添字〜終了する添字+1で指定して取得。

    // 連続した値
    fmt.Println(example[1:3])
    // 指定した位置以降（以前だと[:2]）
    fmt.Println(example[2:])

実行結果

 　　 [hands-on in]
 　　 [in kagawa]

* スライス (3/3)

- 変更したい要素の位置に値を代入。

    // 4番目の値を変更
    example[3] = "okayama"
    fmt.Println(example)

実行結果

    [Golang hands-on in okayama]

* 課題１ (01_slice.go)

１−１：

hello.goをコピーして01_slice.goを作成してください。
以降は01_slice.goを編集してください。

１−２：

var文のmsg1とmsg2を消した後、次の変数を定義してください。

- 変数名：example
- 型：string型のスライス
- 値："golang", "hands-on", "in", "kagawa"

１−３：

main()関数のmsg2を削した後、exampleを標準出力してください。

* 解答１ (01_slice.go)

    package main

    import (
        "fmt"
    )

    var (
        // 値をまとめて代入（}の前に要素があるので最後の要素にカンマは不要！）
        example = []string{"golang", "hands-on", "in", "kagawa"}
    )

    func main() {
        fmt.Println(example)
    }

実行結果

    $ go run 01_slice.go 
    [golang hands-on in kagawa]


* 繰り返し

* 繰り返し文 (1/4)

- for文のみ（while文などなし！シンプル！）

    for i := 0; i < 4; i++ {
        fmt.Print(i)
    }

実行結果

  0123

* 繰り返し文 (2/4)

- 無限ループ

    // 「:=」記述：変数iは宣言＋初期化（値から推測してint型）
    i := 1
    for {
        fmt.Print(i)
        i++
    }

実行結果

  123456789101112131415（略）

* 繰り返し文 (3/4)

- rangeでスライス型を指定すると添字と値を取得できる。

    // 変数abcをstring型のスライスで宣言して、値を設定。
    var abc []string
    abc = []string{"a", "b", "c"}

    // 第1戻り値変数で添字、第2戻り値変数で値が取得できる。
    for i, v := range abc {
        fmt.Println(i, v)
    }
    // 標準出力に空行を出力
    fmt.Println()

    // 値だけを取得する場合は、第1戻り値変数を_にして第2戻り値変数を指定する。
    for _, v := range abc {
        fmt.Println(v)
    }
    fmt.Println()

    // 添字だけを取得する場合は、第1戻り値変数のみ指定する（i, _でも良い）
    for i := range abc {
        fmt.Println(i)
    }

* 繰り返し文 (4/4)

実行結果

    0 a
    1 b
    2 c

    a
    b
    c

    0
    1
    2

* 課題２ (02_for.go)

２−１：

01_slice.goをコピーして02_for.goを作成してください。
以降は02_for.goを編集してください。

２−２：

main()関数内を、以降に書き換えてください。

１．次の変数を宣言してください。

- 変数名：data
- 型：string型のスライス

２．変数dataにexampleを設定（代入）してください。

３．for文とrangeを使用してdataの値だけを順番に標準出力に表示してください。

* 解答２ (02_for.go)

    package main

    // 例）1行1パッケージの書き方
    import "fmt"

    var (
        example = []string{"golang", "hands-on", "in", "kagawa"}
    )

    func main() {
        var data []string
        data = example
        for _, v := range data {
            fmt.Println(v)
        }
    }

* 解答２ (02_for.go)

実行結果

    $ go run 02_for.go
    golang
    hands-on
    in
    kagawa

* if文

* if文

- 三項演算子はない！シンプル！
- 判定文としては、switch文もある。

    v := "GoLang"

    if v == "golang" {
        fmt.Println("○")
    } else if v == "GOLANG" {
        fmt.Println("△")
    } else {
        fmt.Println("×")
    }

    // 変数vはif文の内外からアクセスできる
    fmt.Println(v)

実行結果

    ×
    GoLang

* if文

- 条件式の前に式を記述できる

    if v := "GoLang"; v == "golang" {
        fmt.Println("○")
    } else if v == "GOLANG" {
        fmt.Println("△")
    } else {
        fmt.Println("×")
    }

    // 変数vのネームスコープがif文内だけになるので、外からアクセスするとエラーになる
    // fmt.Println(v)

* if文

実行結果（fmt.Println(v)を無効時）

    ×

実行結果（fmt.Println(v)を有効時）

    # command-line-arguments
    ./main.go:17:17: undefined: v

* 課題３ (03_if.go)

３−１：

02_for.goをコピーして03_if.goを作成してください。
以降は03_if.goを編集してください。

３−２：

main()関数のfor文内を、以降に書き換えてください。

１．fmt.Println(v)の後に、ループの値変数vの値が"in"なら"○"を、それ以外なら"×"を標準出力に表示するif文を追加してください。

* 解答３ (03_if.go)

    package main

    import "fmt"

    var (
        example = []string{"golang", "hands-on", "in", "kagawa"}
    )

    func main() {
        var data []string
        data = example
        for _, v := range data {
            fmt.Println(v)

            if v == "in" {
                fmt.Println("○")
            } else {
                fmt.Println("×")
            }
        }
    }

* 解答３ (03_if.go)

実行結果

    golang
    ×
    hands-on
    ×
    in
    ○
    kagawa
    ×

* 標準入出力

* 標準入力

    fmt.Print("入力は？> ")
    var ans string

    fmt.Scanln(&ans)

    // 変数名に&がついていないと値が書き変わらない
    // fmt.Scanln(ans)

    fmt.Printf("入力は %s です\n", ans)

- 関数の引数は「値を渡す」（引数の値をコピーする）方式になっている。
- 実変数の変数名の先頭に&をつけると、その変数の参照（ポインター）となる。
- Scanln()関数は、引数に入力された値を設定するので「参照を渡し」ている。

* 標準入力

実行結果（fmt.Scanln()関数の変数名に&あり）

    入力は？> Go
    入力は Go です

実行結果（fmt.Scanln()関数の変数名に&なし）

    入力は？> 入力は  です

* fmtパッケージ

- 標準出力、標準入力、標準エラーの関数が集約されている。

.link https://golang.org/pkg/fmt/

.image go-intro-hands-on/img/fmt.png

* fmtパッケージ

- Print系は主に標準出力

    // 改行させる場合は、改行コードを含める。
    fmt.Print("golang ")
    fmt.Print("2018\n")

    // 最後にf付きは、フォーマット形式。
    fmt.Printf("%s %d\n", "golang", 2018)

    // 最後にln付きは、文末に改行コードが付く。
    fmt.Println("golang", "2018")

    // 最初にF付きは、第1引数に出力先を指定（ここでは標準エラー）
    fmt.Fprintln(os.Stderr, "golang", "2018")

    // 最初にS付きは、文字列を返す
    s := fmt.Sprintf("%s %d", "golang", 2018)
    fmt.Println(s)

* fmtパッケージ

出力結果（上記全て）

    golang 2018

- Scan系は標準入力で、Print系と同じような命名規則。
- 標準エラーは、Errorf()関数や一部のPrint系で対応。

fmtパッケージはGoのパッケージについて学び始めるのに最適かと！q@w@p

* 課題４ (04_stdio.go)

４−１：

03_if.goをコピーして04_stdio.goを作成してください。
以降は04_stdio.goを編集してください。

４−２：

main()関数のfor文内のfmt.Println(v)の後に、次の要件を満たす処理を追加してください。

１．入力プロンプトとして"input : "を標準出力に表示する。
２．標準入力を受け取る変数はstring型で変数名はansとする。
３．ループの値変数vの値が標準入力された値なら"○"を、それ以外なら"×"を標準出力に表示する。

* 解答４ (04_stdio.go)

    package main

    import "fmt"

    // 例）1行1変数定義の書き方
    var example = []string{"golang", "hands-on", "in", "kagawa"}

    func main() {
        var data []string
        data = example
        for _, v := range data {
            fmt.Println(v)

            fmt.Print("input : ")
            var ans string
            fmt.Scanln(&ans)

            if v == ans {
                fmt.Println("○")
            } else {
                fmt.Println("×")
            }
        }
    }

* 解答４ (04_stdio.go)

実行結果

    golang
    input : golang
    ○
    hands-on
    input : hands-on
    ○
    in
    input : in
    ○
    kagawa
    input : okayama
    ×

* 実行バイナリーを作成

初期状態

    $ ls
    01_slice.go 02_for.go   03_if.go    04_stdio.go hello.go

go build ファイル名

    $ go build 04_stdio.go 
    $ ls
    01_slice.go 02_for.go   03_if.go    04_stdio    04_stdio.go hello.go

- ファイル名と同名の実行バイナリー「04_stdio」が生成される。

go build -o 実行バイナリー名 ファイル名

    $ go build -o typing 04_stdio.go 
    $ ls
    01_slice.go 02_for.go   03_if.go    04_stdio    04_stdio.go hello.go    typing

- 指定した名前の実行バイナリー「typing」が生成される。

* 完成！q@w@p

    $ ./typing
    golang
    input : golang
    ○
    hands-on
    input : hands-on
    ○
    in
    input : in
    ○
    kagawa
    input : kagawa
    ○
