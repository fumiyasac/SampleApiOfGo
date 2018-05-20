# GolangでのAPI作成のプラクティス用

一応まだ作成中ではありますが、Go言語で(RESTっぽい)APIを書いて見る練習がしたかったので作ってみました。おそらくGo言語っぽさがありましたらPull Requestを頂ければ幸いです。

### このリポジトリの構成

- `constants`    ... どこでも使うような定数値やenumのような形で使いたいものを格納する。
- `controllers`  ... エンドポイントからのリクエスト/レスポンスに関わるものを格納する。
- `database`     ... 接続するデータベース(ORMの接続設定やMigraion関連)に関わるものを格納する。
  - `engine.go   (ORMを利用するための処理)`
  - `dbconf.yaml (マイグレーション実行用)`
  - `migrations/ (マイグレーションファイル置き場)`
- `docs`         ... Swaggerで管理するドキュメントに関わるものを格納する。
- `entities`     ... 接続しているテーブルのDBスキーマ定義と構造体を格納する。
- `factories`    ... APIレスポンス形式の雛形となる構造体を格納する。
- `repositories` ... APIレスポンスを構築するためのロジックに関わるものを格納する。
- `static`       ... HTMLやCSS等の静的ファイルに関わるものを格納する。
- `validators`   ... エンドポイントからControllerに渡された値の妥当性検証に関わるものを格納する。
- `main.go`      ... 主にここではエンドポイントの管理を担当している。

大雑把な処理の流れは下記のような流れになります。

1. controllerでリクエストを受け取ってvalidatorで妥当性を検証する。
2. 1.が妥当であればcontrollerに対応するrepositoriesの処理(Interfaceを定義すること!)を実行する。
3. 2.ではMySQLからentityで定義した形のデータを取得する等のロジックに関わる処理を行いfactoryで定義したデータを作成する。
4. 3.の処理結果に応じてfactoryで定義した形のレスポンスを返却する。

### Go言語標準以外のパッケージ

現在は下記のGo言語のパッケージを用いて開発しています。
ロジックの構築に加えて運用保守に関わるもの等、便利そうなものは随時導入していく予定です。

__1. API構築に関わるもの__

| パッケージ名 | 用途 |
|:---|:---|
|[Gin](https://github.com/gin-gonic/gin) |APIのルーティング管理 |
|[xorm](https://github.com/go-xorm/xorm) |MySQLでのデータ取得用のORM |
|[bcrypt](https://gowebexamples.com/password-hashing/) |パスワードのハッシュ化 |

__2. 運用保守に関わるもの__

| パッケージ名 | 用途 |
|:---|:---|
|[goose](https://github.com/pressly/goose) |MySQLのDBマイグレーションツール |
|[swag](https://github.com/swaggo/swag)<br>[gin-swagger](https://github.com/swaggo/gin-swagger)<br>[gin-swagger/swaggerFiles](https://github.com/swaggo/gin-swagger/swaggerFiles)<br>[go.uuid](go.uuid) |Ginを利用した構成でGodocからSwaggerによるAPI定義書の書き出し |

### 事前準備と補足

※ 以降は少なくともGo言語とMySQLがインストールされていることを前提としています。

__1. パッケージの取得:__

導入していないパッケージがある場合は下記のコマンドにて導入をしておく。

```
$ go get -u (パッケージ名)
```

__2. DB定義のマイグレーション:__

DB構造の変更や追加がある際は下記のコマンドをGooseのREADME等を参考にし、実行するSQLを作成した後にマイグレーションを実行すること。
※1. イメージとしてはRailsの`rake db:migrate`と似たようなことをする。
※2. `-path "database"`としているのは`dbconf.yaml`の場所がデフォルトと異なるため。

```
$ goose -path "database" (実行したい命令)
```

参考:

+ [Golang製DBマイグレーションツールgoose + MySQLを試してみた](https://qiita.com/K_ichi/items/b9362e3a3c5688e494e2)
+ [gooseによるテーブルマイグレーション管理](http://engineering.enish.jp/?p=994&doing_wp_cron=1526793213.5329029560089111328125)
+ [golang製のDBマイグレーションツールgooseをMySQLで使ってみる](http://shusatoo.net/programming/golang/goose-mysql-migration/)

__3. API定義書の自動書き出し:__

APIの処理を新たに作成した場合にはコメントにAPI定義を記載しておく。
記載が終わったら下記のコマンドを実行し、API定義書を更新する。

参考:

+ [[swaggo]GoのGoDocを書いたら、Swaggerを出せるやばいやつ](https://qiita.com/pei0804/items/3a0b481d1e47e5a72078)
+ [swaggoが思いの外、素晴らしかった](https://syossan.hateblo.jp/entry/2018/05/15/175653)

```
$ swag init
$ go run main.go
```

起動ができたら下記のURLから、SwaggerUIと照らし合わせて、定義書と挙動が問題ないことを確認しておく。

+ [チェック!](http://localhost:8080/swagger/index.html)

