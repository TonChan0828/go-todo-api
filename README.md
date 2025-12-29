# go-todo-api
Go と Gin による RESTful な Todo API。

## Overview
- レイヤードアーキテクチャに基づき、責務分離とテスト容易性を重視した構成
- フレームワーク非依存のユースケース層と、HTTP ハンドラー層を分離
- 永続化はインメモリ実装で、差し替えを想定した設計

## Features
- Todo の作成、一覧取得、完了状態の更新、削除
- UUID による ID 管理
- バリデーションとエラーハンドリング

## Architecture
- 依存方向は内向きのみ: Handler → Usecase → Domain
- Handler: Gin を用いた HTTP 入出力とルーティング
- Usecase: アプリケーションロジック、データ操作、エラー制御
- Domain: ビジネスモデル

## Project Structure
```
cmd/
  api/
    main.go
internal/
  domain/
    todo.go
  handler/
    routes.go
    todo_handler.go
    todo_handler_test.go
  usecase/
    todo_usecase.go
    todo_inmemory.go
    todo_inmemory_test.go
  repository/
  infrastructure/
    db/
```

## API Endpoints
- `GET /health` ヘルスチェック
- `POST /todos` Todo 作成
- `GET /todos` Todo 一覧取得
- `PUT /todos/:id` 完了状態の更新
- `DELETE /todos/:id` Todo 削除

## How to Run
```
go run ./cmd/api
```
- サーバーは `:8080` で起動

## Testing
```
go test ./...
```
- ユースケース層のユニットテスト
- `net/http/httptest` を用いた HTTP ハンドラーのテスト
