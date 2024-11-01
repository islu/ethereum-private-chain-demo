
# Devlog #4
## Idea

**交易資料同步進資料庫**

初步的研究，文件中似乎沒有提供篩選指定地址交易紀錄的 function

但是有可以取得所有交易紀錄的 funcion

那這個時候有幾個選擇可以做
1. 將所有交易紀錄存進資料庫
2. 只將指定地址交易紀錄進資料庫

這個地方就看需求了，有可能未來會需要提供每一個地址的交易紀錄，另外進資料庫篩選功能會比較好實現

我這邊先以簡單的方式實現只將指定地址的交易紀錄進資料庫


**Cron job 的實現**

這方面使用任何語言應該都可以實現，這邊想要做的是功能由後端這邊提供，由另外一個服務去做排程，當然也可以將功能與排程都包在同一個服務中，主要是因為涉及會去存取資料庫，在專案中不希望太過複雜造成管理上的困難

## Survey & Impl

**[問題]** `tx.AsMessage` 沒有這個方法

**[解答]** [type *types.Transaction has no field or method AsMessage)](https://ethereum.stackexchange.com/questions/149220/type-types-transaction-has-no-field-or-method-asmessage)

```go
// tx types.Transactions
    if from, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx); err == nil {
		fmt.Println("Tx From: ", from.Hex())
	}
```


**[問題]**

```zsh
failed to connect to `user=postgres database=public`: [::1]:5432 (localhost): server error: FATAL: database "public" does not exist (SQLSTATE 3D000)
```

**[解答]**

在我的情境下預期的是本地程式能夠讀取容器建立的資料庫，但是同時有本地的資料庫及容器的資料庫，有 port 是一樣的，本地程式執行時是讀到本地的資料庫，所以出現上面的錯誤

解決方式，可以使用不同的 port 或是先停掉本地的資料庫

## Summery


## References


[輕鬆看懂 Etherscan](https://medium.com/taipei-ethereum-meetup/viewing-blocks-and-transactions-on-etherscan-3bb5b3685ba7)


[從 0 認識 Blockchain - Transaction 以及你該知道的一切](https://ambersun1234.github.io/blockchain/blockchain-transaction/)

- Block 包含零或多個 Transaction

[公鑰、私鑰、地址傻傻分不清！？](https://medium.com/taipei-ethereum-meetup/%E5%85%AC%E9%91%B0-%E7%A7%81%E9%91%B0-%E5%9C%B0%E5%9D%80%E5%82%BB%E5%82%BB%E5%88%86%E4%B8%8D%E6%B8%85-f83b8fe424e)

- 地址長度 160 bits
