# Flashbots RPC client

Fork of [ethrpc](https://github.com/onrik/ethrpc) with additional [Flashbots RPC methods](https://docs.flashbots.net/flashbots-auction/searchers/advanced/rpc-endpoint):

* `FlashbotsCallBundle` ([`eth_callBundle`](https://docs.flashbots.net/flashbots-auction/searchers/advanced/rpc-endpoint/#eth_callbundle))
* `FlashbotsSendBundle` ([`eth_sendBundle`](https://docs.flashbots.net/flashbots-auction/searchers/advanced/rpc-endpoint/#eth_sendbundle))
* `FlashbotsGetUserStats` ([`flashbots_getUserStats`](https://docs.flashbots.net/flashbots-auction/searchers/advanced/rpc-endpoint/#flashbots_getuserstats))
* `FlashbotsSendPrivateTransaction` (`eth_sendPrivateTransaction`)
* `FlashbotsCancelPrivateTransaction` (`eth_cancelPrivateTransaction`)
* `FlashbotsSimulateBlock`: simulate a full block

## Usage

Add library to your project:

`go get github.com/spacelan/flashbotsrpc`

Create a new private key here for testing (you probably want to use an existing one):

```go
privateKey, _ := crypto.GenerateKey()
```

#### Simulate transactions with `eth_callBundle`:

```go
callBundleArgs := flashbotsrpc.FlashbotsCallBundleParam{
    Txs:              []string{"YOUR_RAW_TX"},
    BlockNumber:      fmt.Sprintf("0x%x", 13281018),
    StateBlockNumber: "latest",
}

result, err := rpc.FlashbotsCallBundle(privateKey, callBundleArgs)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", result)
```

#### Get Flashbots user stats:

```go
rpc := flashbotsrpc.New("https://relay.flashbots.net")
result, err := rpc.FlashbotsGetUserStats(privateKey, 13281018)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", result)
```

#### Send a transaction bundle to Flashbots with `eth_sendBundle`:

```go
sendBundleArgs := flashbotsrpc.FlashbotsSendBundleRequest{
    Txs:         []string{"YOUR_RAW_TX"},
    BlockNumber: fmt.Sprintf("0x%x", 13281018),
}

result, err := rpc.FlashbotsSendBundle(privateKey, sendBundleArgs)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", result)
```

#### More examples

You can find example code in the [`/examples/` directory](https://github.com/spacelan/flashbotsrpc/tree/master/examples).
