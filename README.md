# GPS-Reminder-Server

GPS を活用した、所要時間計算機能付きの予定管理アプリの **API サーバーサイド** です。

[フロントエンドはこちら](https://github.com/Outtech105k/GPS-Reminder)

なお、現在は開発版です。

Reminder application (SERVER SIDE) related GPS location.

## Getting Started

1. Docker composeの実行環境を用意します。
1. [.sample.env](/.sample.env)の名前を`.env`に変更し、内容を環境に応じて変更します。
1. Docker composeを起動します。
```bash
docker compose up -d
```

`.env`ファイルの変数内容は以下の通りです。

| Key name | Description |
| --- | --- |
| `BCRYPT_COST` | パスワードのハッシュ化に使う[BCrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)のストレッチング回数を指定します。 |
| `APP_PORT` | サーバのポート番号(クライアントに開放するポート) |
| `MYSQL_ROOTPASS` | MySQLへのrootのログインパス
| `MYSQL_DB` | アプリ内で使用するDB名 |
| `MYSQL_USER` | MySQLへの一般ユーザ名(アプリからのアクセスに使用) |
| `MYSQL_PASS` | MySQLへの一般ユーザのパス |
| `MYSQL_PORT` | MySQLへの接続ポート(アプリからはDockerネットワークを使用するので、変更可) |

## Usage

APIアクセスのエンドポイントは以下の通りです。

| Method | Endpoint | Description |
| --- | --- | --- |
| `POST` | `/signup` | アカウント登録 |
| `POST` | `/signin` | ログイン |

なお、現状ログアウト機能は存在しません。

## Author

Outtech105

[Homepage](https://outtech105.com)

[X](https://x.com/105techno)
