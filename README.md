# go-todo-api
Go と Gin による RESTful な Todo API。

## 概要
- レイヤードアーキテクチャに基づき、責務分離とテスト容易性を重視した構成
- HTTP ハンドラーとユースケースを分離し、データアクセスはリポジトリ経由で実装
- 永続化は PostgreSQL を使用し、sqlc でクエリコードを生成
- In-memory 実装もあり、用途に応じて差し替え可能

## 機能
- Todo の作成、一覧取得、完了状態の更新、削除
- UUID による ID 管理
- バリデーションとエラーハンドリング

## アーキテクチャ
- 依存方向は内向きのみ: Handler → Usecase → Repository → DB
- Handler: Gin を用いた HTTP 入出力とルーティング
- Usecase: アプリケーションロジック、データ操作、エラー制御
- Repository: 永続化の抽象化と実装（PostgreSQL / In-memory）
- DB: sqlc によるクエリ生成コード

## プロジェクト構成
```
cmd/
  api/
    main.go
db/
  schema.sql
  query.sql
internal/
  domain/
    todo.go
  handler/
    routes.go
    todo_handler.go
    todo_handler_test.go
  infrastructure/
    todo_repository.go
    db/
      conn.go
      query.sql.go
  repository/
    todo_postgres.go
    todo_postgres_test.go
    test_helper.go
  usecase/
    todo_usecase.go
    todo_repo_usecase.go
    todo_inmemory.go
    todo_inmemory_test.go
```

## API エンドポイント
- `GET /health` ヘルスチェック
- `POST /todos` Todo 作成
- `GET /todos` Todo 一覧取得
- `PUT /todos/:id` 完了状態の更新
- `DELETE /todos/:id` Todo 削除

## セットアップ
### PostgreSQL を起動
```
docker-compose up -d
```

### 環境変数（任意）
```
export DATABASE_URL=postgres://todo:todo@localhost:5432/todo?sslmode=disable
```

## 起動方法
```
go run ./cmd/api
```
- サーバーは `:8080` で起動
- In-memory 実装を使う場合は `cmd/api/main.go` の `NewInMemoryTodoUsecase()` を有効化

## テスト
```
go test ./...
```
- ユースケース層のユニットテスト
- `net/http/httptest` を用いた HTTP ハンドラーのテスト
- PostgreSQL を使うリポジトリテストも含まれるため、DB が必要
