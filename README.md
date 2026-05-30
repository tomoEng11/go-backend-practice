# go-backend-practice

Go + Echo を使ったバックエンドAPI練習用リポジトリ

## セットアップ

```bash
go mod download
```

## サーバー起動

```bash
go run main.go
```

## OpenAPI コード生成

### 必要なパッケージ

```bash
go get github.com/oapi-codegen/oapi-codegen/v2
go get github.com/oapi-codegen/runtime
```

### oapi-codegen のインストール

```bash
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
```

### コード生成の実行

```bash
oapi-codegen --config api/cfg.yaml api/openapi.yaml
```

### go generate を使う場合

`api/generate.go` に以下を記述:

```go
package api

//go:generate oapi-codegen --config cfg.yaml openapi.yaml
```

その後、以下のコマンドで生成:

```bash
go generate ./...
```
