# cleanarchitecture-learn
goでクリーンアーキテクチャ書いてみた

## 動かし方
### マイグレーション
gooseでスキーマ管理しているので、gooseをインストールしていない場合以下を実行
```
go get -u github.com/pressly/goose/cmd/goose
```

migrate
```
goose up
```

### アプリケーション起動
```
docker-compose up
```
localhost:9002がフロントエンドです
