#!/bin/bash

apt-get update && apt-get install -y unzip

curl -OL https://github.com/google/protobuf/releases/download/v3.19.4/protoc-3.19.4-linux-x86_64.zip
unzip protoc-3.19.4-linux-x86_64.zip -d protoc3
sudo mv protoc3/bin/* /usr/local/bin/
sudo mv protoc3/include/* /usr/local/include/

# protoc-gen-docのインストール
go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest

# docディレクトリがなければ作る
if [ ! -d ./docs ]; then
  mkdir ./docs
fi

# ドキュメント作成(markdownとhtmlの2種類を生成)
protoc \
  --doc_out=./docs \
  --doc_opt=html,index.html \
  ./grpc/*.proto

# ごみ削除
if [ -d ./protoc3 ]; then
  rm -r protoc3
fi

if [ -e protoc-3.19.4-linux-x86_64.zip ]; then
  rm protoc-3.19.4-linux-x86_64.zip
fi