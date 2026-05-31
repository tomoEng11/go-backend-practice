# go-backend-practice

Go + Echo を使ったバックエンドAPI練習用リポジトリ

## ディレクトリ構成

```
.
├── spec/                 # OpenAPI仕様
│   ├── public.yaml       # 公開API (認証不要)
│   └── protected.yaml    # 認証必須API
├── gen/                  # 生成コード (編集しない)
│   ├── public/
│   └── protected/
├── cmd/api/              # エントリーポイント
├── internal/
│   ├── server/           # ServerInterface 実装
│   └── middleware/       # ミドルウェア
└── Makefile
```

## セットアップ

```bash
go mod download
```

### oapi-codegen のインストール

```bash
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
```

## 使い方

```bash
# コード生成
make generate

# ビルド
make build

# 実行
make run
```

## エンドポイント

| パス | 認証 | 説明 |
|-----|------|-----|
| GET /public/health | 不要 | ヘルスチェック |
| GET /api/channels | 必要 | チャンネル一覧 |
| GET /api/users | 必要 | ユーザー一覧 |
