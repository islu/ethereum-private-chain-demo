

## 部署流程

1. 安裝 Docker 和 Docker Compose

2. 啟動 Docker

3. 使用 `git clone` 該專案

```zsh
git clone https://github.com/islu/ethereum-private-chain-demo.git
```

4. 切換到專案目錄

```zsh
cd ethereum-private-chain-demo
```

5. 啟動服務

```zsh
docker-compose up --build -d
```

6. 為了同步資料需要重啟容器。待確認是否有更好的方式

```zsh
docker restart ethereum-private-chain-demo-geth-rpc-endpoint-1
```

7. 開啟 http://localhost:8080/swagger/index.html


8. 當不再需要服務時，使用以下命令

```zsh
docker-compose down
```
