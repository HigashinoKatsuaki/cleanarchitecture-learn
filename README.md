# cleanarchitecture-learn
goでクリーンアーキテクチャ書いてみた

## 動かし方
### mysql起動
```
docker-compose up
```

### マイグレーション
gooseでスキーマ管理しているので、gooseをインストールしていない場合以下を実行
```
go get -u github.com/pressly/goose/cmd/goose
```

migrate
```
goose up
```

### バックエンドを起動
```
go run src/main.go
```

### フロントエンドを起動
```
yarn run next -p 9002
```
※ ポートは好きなところでOK
