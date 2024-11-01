
# Devlog #2

## Idea

後端服務參考 [chatbotgang / go-clean-arch](https://github.com/chatbotgang/go-clean-arch)當作基礎 

這邊就會是比較熟悉的領域，主要使用 Gin 框架實現

沒接觸過 JSON PRC，只有實作過 gPRC 但看起來是是不一樣的東西，所以重點會是在處理 JSON RPC

## Survey & Impl


- [使用 Golnag 打造 Web 應用程式 - RPC](https://willh.gitbook.io/build-web-application-with-golang-zhtw/08.0/08.4)
- [RPC JSON-RPC 簡介](https://note.pcwu.net/2017/06/05/json-rpc-intro/)
- [Beginner's Guide to RPC in Golang: Understanding the Basics](https://dev.to/atanda0x/a-beginners-guide-to-rpc-in-golang-understanding-the-basics-4eeb)

JSON PRC 通常主要是以 POST 並以 Json 格式當作參數，透過參數指定需要呼叫 function

---

官方建議開發者的文件 [Go API](https://geth.ethereum.org/docs/developers/dapp-developer/native)，有提到這個套件 [ethereum / go-ethereum](https://github.com/ethereum/go-ethereum) 因此決定使用

另外，應該是可以根據 [Ethereum JSON-RPC Specification](https://ethereum.github.io/execution-apis/api-documentation) 提供的 Spec 一樣做出來，不一定需要使用套件吧

時間上的關係依然決定使用套件，如果有需要使用到未實作的 function 時，再來實作


---

**[問題]** 產生出來的可以匯入 MetaMask 嗎

**[解答]** 可以，可以轉成私鑰再匯入

---

**[問題]** VS Code 開發過程出現警告

```
packages.Load error: err: exit status 1: stderr: go: updates to go.mod needed; to update it:  
go mod tidy
```

**[解答]？**

 [x/tools/gopls: Error loading workspace: go: updates to go.mod needed, disabled by -mod=readonly : packages.Load error #44085](https://github.com/golang/go/issues/44085)

自己解決的方式 `Go mod tidy` 和升級套件就可以

## Summery


## References


[JSON-RPC 2.0 Specification](https://www.jsonrpc.org/specification)

> JSON-RPC is a stateless, light-weight remote procedure call (RPC) protocol. Primarily this specification defines several data structures and the rules around their processing. It is transport agnostic in that the concepts can be used within the same process, over sockets, over http, or in many various message passing environments. It uses [JSON](http://www.json.org/) ([RFC 4627](http://www.ietf.org/rfc/rfc4627.txt)) as data format.

