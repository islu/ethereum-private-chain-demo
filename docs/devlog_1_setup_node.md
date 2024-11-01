# Devlog #0

**初步構想**

1. Devlog 會是比較隨意的紀錄，所以會比較碎片化，主要記錄當時的想法及問題
2. 專案說明跟部署流程會放在 `README.md` 並提供簡單的架構圖
3. 任何的做法都是從想法出發，所以會盡量的描述為什麼這麼做
4. 每個人當下的背景知識不一樣， 所以有些地方會因為我已經知道而不會寫出來

在文件中的某些段落我會做一些標記，方便理解我想表達什麼，如：

- **[問題]**：遇到的問題
- **[解答]**：解決的方式
- **[待辦]**：時間上可能來不及有機會再研究
- **[想法]**：為什麼樣這樣做

# Devlog #1

## Idea

我是第一次自行建立區塊鏈，有基本的概念但是沒有實際動手過，有聽過 `Geth` 所以會先以此工具這個為主。

研究流程會是：
1. 本地架設，用來理解架構及原理
2. 容器化，統一環境方便部署

**[待辦]**
- `Geth` 與 `Nethermind` 有什麼區別

## Survey & Impl

先安裝 [Installing Geth](https://geth.ethereum.org/docs/getting-started/installing-geth)

Mac 安裝
```zsh
brew tap ethereum/ethereum
brew install ethereum
```

也有提供 Docker 版本
```
docker pull ethereum/client-go
docker run -it -p 30303:30303 ethereum/client-go
```

確認版本
```zsh
geth -v

geth version 1.14.11-stable
```

接著參考別人的架設流程
1. [建立以太坊私有鏈(一) 節點部署](https://notes.andywu.tw/2018/%E5%BB%BA%E7%AB%8B%E4%BB%A5%E5%A4%AA%E5%9D%8A%E7%A7%81%E6%9C%89%E9%8F%88%E4%B8%80/)
2. [How to build an Ethereum private blockchain using Geth](https://hemantkgupta.medium.com/how-to-build-an-ethereum-private-blockchain-network-using-geth-3133e23f729d)
3. [Ethereum 開發筆記 2–2：Geth 基礎用法及架設 Muti-Nodes 私有鏈](https://blog.fukuball.com/ethereum-%E9%96%8B%E7%99%BC%E7%AD%86%E8%A8%98-22geth-%E5%9F%BA%E7%A4%8E%E7%94%A8%E6%B3%95%E5%8F%8A%E6%9E%B6%E8%A8%AD-muti-nodes-%E7%A7%81%E6%9C%89%E9%8F%88/)

會發現有些參數跟文件不太一樣，因此需要再確認可不可行。另外以太坊也更換成 PoS，所以這個工具也改版了，這可能使問題更加複雜。

不過還是先選擇參考教學的部署流程，了解架構之後有機會再轉換。

---

**[問題]** 該使用 puppeth 嗎

很多教學文件有提到這個工具可以方便產生相關設定檔

**[想法]** 我這裡選擇直接建立 `genesis.json`

官方有提到 [Puppeth](https://geth.ethereum.org/docs/tools/puppeth) 已經不包含在 `Geth` 裡面了，所以這裡就不使用了
> Note 
> Puppeth was [removed from Geth](https://github.com/ethereum/go-ethereum/pull/26581) in January 2023.

---

最後其實是參考官方的部署流程 [Private Networks](https://geth.ethereum.org/docs/fundamentals/private-network) 而完成，還是以官方文件為主。詳細如何做可以直接看 `End-to-end example` 段落

實作過程中有遇到一些問題

**[問題]**

```
Fatal: Failed to register the Ethereum service: only PoS networks are supported, please transition old ones with Geth v1.13.x
```

**[解答]**

[Fatal: Failed to register the Ethereum service: only PoS networks are supported, please transition old ones with Geth v1.13.x #30120](https://github.com/ethereum/go-ethereum/issues/30120)

```json
{
  "config": {
    "chainId": 8888,
    "homesteadBlock": 0,
    "eip150Block": 0,
    "eip155Block": 0,
    "eip158Block": 0,
    "byzantiumBlock": 0,
    "constantinopleBlock": 0,
    "petersburgBlock": 0,
    "istanbulBlock": 0,
    "berlinBlock": 0,
    "clique": {
      "period": 5,
      "epoch": 30000
    },
    "terminalTotalDifficultyPassed": true
  },
  // ...
}
```

多增加 `"terminalTotalDifficultyPassed": true` 去略過，可以成功運行

但是這裡有另一個問題，節點無法進行打包區塊的動作，即使啟用 `--mine` 一樣不行

```zsh
    --mine                              (default: false)                   ($GETH_MINE)
    Enable mining

```

所以最後決定還是降版本至 1.13 後看正不正常，這裡決定使用容器化方式處理

這裡先用最簡單的方式將本地的建置好的環境複製到容器，先試看看能不能正常運行，測試結果是可以的

---


**[問題]** 啟用 JSON RPC

**[解答]** 

剛開始很困惑啟用 http 到底是不是啟用，http 是其中一種方式

文件有提到 [JSON-RPC Server](https://geth.ethereum.org/docs/interacting-with-geth/rpc)

> JSON-RPC is provided on multiple transports. Geth supports JSON-RPC over HTTP, WebSocket and Unix Domain Sockets. Transports must be enabled through command-line flags.

啟用方式
```zsh
geth --http
```


---

**[問題]** 可以使用 MetaMask 連接到私有鏈嗎

**[解答]** 可以，只要新增網路及 ID 就可以連上


## Summery


## References

[Command-line Options](https://geth.ethereum.org/docs/fundamentals/command-line-options)

[Private Networks](https://geth.ethereum.org/docs/fundamentals/private-network)

[JSON-RPC Server](https://geth.ethereum.org/docs/interacting-with-geth/rpc)

[以太坊引導節點介紹](https://ethereum.org/zh-tw/developers/docs/nodes-and-clients/bootnodes/)
