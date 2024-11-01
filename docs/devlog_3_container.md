
# Devlog #3

## Idea

主要是把 Devlog 1 再本地建立以容器化

首先會找看看有沒有範本可以參考

## Survey & Impl

主要參考 
- [Deploying Private Ethereum Blockchain with Geth](https://web3coda.com/labcontent/?name=ethereum-blockchain-deployment)
- [How to build an Ethereum private blockchain network using geth and Docker](https://hemantkgupta.medium.com/how-to-build-an-ethereum-private-blockchain-network-using-geth-and-docker-41f2ce8d6f6e)

有提供程式碼 [web3coda / ethereum-blockchain-setup](https://github.com/web3coda/ethereum-blockchain-setup)

所以基本上會以上面的範本做修改

---

**[問題]** Node 節點無法同步資料

```
2024-10-29 13:13:16 INFO [10-29|05:13:16.009] Block synchronisation started
2024-10-29 13:13:16 INFO [10-29|05:13:16.025] Imported new chain segment               number=3 hash=2318af..0a86f8 blocks=3 txs=0 mgas=0.000 elapsed=4.531ms     mgasps=0.000 triedirty=0.00B
2024-10-29 13:13:16 INFO [10-29|05:13:16.025] Syncing: chain download in progress      synced=+Inf% chain=18.00B headers=3@6.00B bodies=3@6.00B receipts=3@6.00B eta=-5.812ms
2024-10-29 13:13:16 WARN [10-29|05:13:16.026] Synchronisation failed, retrying         err="sync cancelled"
2024-10-29 13:13:16 INFO [10-29|05:13:16.030] Indexed transactions                     blocks=4 txs=0 tail=0 elapsed=5.073ms
2024-10-29 13:13:20 INFO [10-29|05:13:20.875] Looking for peers                        peercount=1 tried=1 static=0
2024-10-29 13:13:31 WARN [10-29|05:13:31.012] Syncing, discarded propagated block      number=4 hash=23e2f9..d96d68
2024-10-29 13:13:46 INFO [10-29|05:13:46.020] Syncing: chain download in progress      synced=100.00% chain=18.00B headers=5@6.00B bodies=3@6.00B receipts=3@6.00B eta=0s
2024-10-29 13:13:54 INFO [10-29|05:13:54.049] Syncing: chain download in progress      synced=100.00% chain=18.00B headers=5@6.00B bodies=3@6.00B receipts=3@6.00B eta=0s
2024-10-29 13:14:02 INFO [10-29|05:14:02.067] Syncing: chain download in progress      synced=100.00% chain=18.00B headers=6@6.00B bodies=3@6.00B receipts=3@6.00B eta=0s
2024-10-29 13:14:10 INFO [10-29|05:14:10.096] Syncing: chain download in progress      synced=100.00% chain=18.00B headers=6@6.00B bodies=3@6.00B receipts=3@6.00B eta=0s
2024-10-29 13:14:18 INFO [10-29|05:14:18.131] Syncing: chain download in progress      synced=100.00% chain=18.00B headers=7@6.00B bodies=3@6.00B receipts=3@6.00B eta=0s
2024-10-29 13:14:26 INFO [10-29|05:14:26.165] Syncing: chain download in progress      synced=100.00% chain=18.00B headers=7@6.00B bodies=3@6.00B receipts=3@6.00B eta=0s
2024-10-29 13:14:34 INFO [10-29|05:14:34.200] Syncing: chain download in progress      synced=100.00% chain=18.00B headers=8@6.00B bodies=3@6.00B receipts=3@6.00B eta=0s
```


**[解答]？** 

本來想說是需要同步的時間，但是會上面的 log 會卡在 100 %

別人似乎也有
[Geth stuck on 100%: Syncing: chain download in progress, Validator down since Dencun :( #29417](https://github.com/ethereum/go-ethereum/issues/29417)

但是似乎無法解決我現在的問題，目前的解決方式，重新啟動該節點

有看到這段 log 就代表同步成功了
```
Imported new chain segment               number=5 hash=1bb2fe..677131 blocks=3 txs=0 mgas=0.000 elapsed=7.184ms mgasps=0.000 triedirty=0.00B ignored=1
```

---

**[問題]** 啟用的服務容器呼叫開放 JSON RPC 的端點無法呼叫

```zsh
403 Forbidden: invalid host specified
```

**[解答]**

[Ethereum client-go RPC response 403 “invalid host specified #16526”](https://github.com/ethereum/go-ethereum/issues/16526)

[Command-line Options](https://geth.ethereum.org/docs/fundamentals/command-line-options) v1.13.15
```zsh
--http.vhosts value                 (default: "localhost")             ($GETH_HTTP_VHOSTS)
    
    Comma separated list of virtual hostnames from which to accept requests (server
    enforced). Accepts '*' wildcard.
```

可以設定成以下

```zsh
--http.vhosts='*'
```

實務上可能不太適合開放全部，但是測試開發先設定成這樣

---

**[問題]**

```
Fatal: Failed to unlock account 0x8De0c53FC169BA09F111aA4170697e8CF42CCbBe (no key for given address or file)
```

**[解答]** 

沒有提供相對應的帳戶私鑰，需要了解一下容器中的目錄結構，看有沒有將帳戶私鑰放到對得位置

---

**[問題]** 容器中服務連接 Boot Node

**[解答]** 

encode 格式 `enode://ENODE-value@IP:PORT`

在 Docker 環境 `IP` 需要替換成 `hostname`，又因為可以自行去命名名稱，所以不會是 localhost 需要特別注意


## Summery



## References

