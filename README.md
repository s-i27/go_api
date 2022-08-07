# 必要なファイル群の作成コマンド
touch Dockerfile docker-compose.yml main.go
$ docker-compose up -d --build
docker-compose ps
docker-compose exec golang /bin/sh
* /bin/sh は /bin/bash のシンボリックリンク

# コンテナ内で、以下のコマンドを実行する
go run main.go