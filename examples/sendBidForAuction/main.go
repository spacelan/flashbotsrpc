package main

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/metachris/flashbotsrpc"
)

var privateKey, _ = crypto.GenerateKey() // creating a new private key for testing. you probably want to use an existing key.
// var privateKey, _ = crypto.HexToECDSA("YOUR_PRIVATE_KEY")

func main() {
	SendPrivTx()
}

func SendPrivTx() {
	rpc := flashbotsrpc.New("https://relay.flashbots.net")
	rpc.Debug = true

	sendBidArgs := flashbotsrpc.FlashbotsSendBidForAuctionRequest{
		AuctionTxHash: "0xf7042ecc5219488406f6ca7d7aab36c90db2bd9840bdc60d8084c264dbfa1119",
		BidRawTx:      "0xrawtxforbid",
	}

	err := rpc.FlashbotsSendBidForAuction(privateKey, sendBidArgs)
	if err != nil {
		if errors.Is(err, flashbotsrpc.ErrRelayErrorResponse) {
			// ErrRelayErrorResponse means it's a standard Flashbots relay error response, so probably a user error, rather than JSON or network error
			fmt.Println(err.Error())
		} else {
			fmt.Printf("error: %+v\n", err)
		}
		return
	}

	// Print txHash
	fmt.Println("bid sent")
}

func SendPrivTxWithAuction() {
	rpc := flashbotsrpc.New("https://relay.flashbots.net")
	rpc.Debug = true

	paymentSplit := make(map[string]uint64)
	paymentSplit["0x2f123a9454b21172115cedd81c730d8a2e89e503"] = 95
	paymentSplit["0xcbd6832ebc203e49e2b771897067fce3c58575ac"] = 6

	sendPrivTxArgs := flashbotsrpc.FlashbotsSendPrivateTransactionRequest{
		Tx:           "0xRAWTX",
		Auction:      true,
		PaymentSplit: paymentSplit,
	}

	txHash, err := rpc.FlashbotsSendPrivateTransaction(privateKey, sendPrivTxArgs)
	if err != nil {
		if errors.Is(err, flashbotsrpc.ErrRelayErrorResponse) {
			// ErrRelayErrorResponse means it's a standard Flashbots relay error response, so probably a user error, rather than JSON or network error
			fmt.Println(err.Error())
		} else {
			fmt.Printf("error: %+v\n", err)
		}
		return
	}

	// Print txHash
	fmt.Printf("txHash: %s\n", txHash)
}