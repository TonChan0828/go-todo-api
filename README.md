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
- Handler: Gin を用いた HTTP 入出力とルーティング（`/v1` 配下）
- Usecase: アプリケーションロジック、データ操作、エラー制御
- Repository: 永続化の抽象化と実装（PostgreSQL / In-memory）
- DB: sqlc によるクエリ生成コード

## プロジェクト構成
```
cmd/
  api/
    main.go
docs/
  BEGINNER_GUIDE.md
  docs.go
  swagger.json
  swagger.yaml
db/
  schema.sql
  query.sql
internal/
  domain/
    todo.go
  handler/
    v1/
      routes.go
      todo_handler.go
      todo_mapper.go
      todo_handler_test.go
      dto/
        todo_request.go
        todo_response.go
        error_response.go
  infrastructure/
    todo_repository.go
    db/
      conn.go
      db.go
      query.sql.go
      models.go
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
- `GET /swagger/index.html` Swagger UI
- `POST /v1/todos` Todo 作成
- `GET /v1/todos` Todo 一覧取得
- `PUT /v1/todos/:id` 完了状態の更新
- `DELETE /v1/todos/:id` Todo 削除

## ドキュメント
- 初学者向けガイド: `docs/BEGINNER_GUIDE.md`

## セットアップ
### PostgreSQL を起動
```
docker-compose up -d
```

### 環境変数（任意）
```
export DATABASE_URL=postgres://todo:todo@localhost:5432/todo?sslmode=disable
```
未設定の場合は上記のデフォルト値が使用されます。

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
