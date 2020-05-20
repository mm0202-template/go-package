# 動作確認
## ローカル
### コマンドチェック
```shell script
docker-compose run local 
```

### テスト
```shell script
docker-compose run test 
```

### Dockerイメージのビルド
```shell script
docker build . 
```

### ビルドイメージのコマンドチェック
```shell script
docker-compose run bundle 
```

### ビルドイメージのテスト
```shell script
docker-compose run bundle-test
```

## リモート
リモートへプッシュした後、以下のコマンドで`Hello World!!`と表示されれば成功。
```shell script
account_name=[アカウント名]
repository_name=[リポジトリ名]

docker run golang bash -c "go get github.com/$account_name/$repository_name && go build github.com/$account_name/$repository_name && $repository_name"
```