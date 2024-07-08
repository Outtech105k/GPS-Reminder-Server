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
