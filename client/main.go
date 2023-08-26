package main

import (
	"fmt"

	appTypes "pipeCoin/app/types"

	"github.com/cometbft/cometbft/crypto/ed25519"
	cometbftHttpClient "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

func main() {
	client, err := cometbftHttpClient.New("http://localhost:26657", "/websocket")
	if err != nil {
		panic(fmt.Sprint("Something Went Wrong While Starting Client: ", err))
	}

	rootCmd := &cobra.Command{
		Use:   "pipeCoinClient",
		Short: "pipeCoin Client For Sending Transactions",
	}
	sendTransactionCmd := &cobra.Command{
		Use:   "send",
		Short: "Transfer Pipes To Receiver",
		Run: func(cmd *cobra.Command, args []string) {
			receiverPubKeyBytes, err := cmd.Flags().GetBytesHex("receiverPubKey")
			if err != nil {
				fmt.Println("Something Went Wrong While Extracting Receiver Public Key, Make Sure You Pass Valid Public Key: ", err)
				return
			}
			if len(receiverPubKeyBytes) != 32 {
				fmt.Println("Public Key Length Must Be 32")
				return
			}

			amount, err := cmd.Flags().GetFloat64("amount")
			if err != nil {
				fmt.Println("Something Went Wrong While Extracting Amount, Make Sure You Pass Valid Amount: ", err)
				return
			}
			if amount <= 0 {
				fmt.Println("Amount Should Be Greater Than Zero")
				return
			}

			userPrivateKeyBytes, err := cmd.Flags().GetBytesHex("privateKey")
			if err != nil {
				fmt.Println("Something Went Wrong While Extracting Private Key, Make Sure You Pass Valid Private Key: ", err)
				return
			}
			if len(userPrivateKeyBytes) != 64 {
				fmt.Println("Private Key Length Must Be 64")
				return
			}

			accountSequence, err := cmd.Flags().GetUint32("sequence")
			if err != nil {
				fmt.Println("Something Went Wrong While Extracting Account Sequence Number, Make Sure You Pass Valid Sequence Number: ", err)
				return
			}

			userPrivateKey := ed25519.PrivKey(userPrivateKeyBytes)
			fmt.Println("Oppp", len(userPrivateKey.PubKey().Bytes()))
			transactionPayload := appTypes.TransferPipesPayload{
				From: userPrivateKey.PubKey().Bytes(),
				To:   receiverPubKeyBytes,
				Amount: &appTypes.Pipe{
					Value: amount,
				},
				Sequence: accountSequence,
			}
			fmt.Printf("Transaction Payload:-\nFrom: %v\nTo: %v\nAmount %v\n", transactionPayload.From, transactionPayload.To, transactionPayload.Amount)
			transactionPayloadBytes, err := proto.Marshal(&transactionPayload)
			if err != nil {
				fmt.Println("Something Went Wrong While Marshaling Transaction Payload: ", err)
				return
			}
			signature, err := userPrivateKey.Sign(transactionPayloadBytes)
			if err != nil {
				fmt.Println("Something Went Wrong While Signing Transaction Payload: ", err)
				return
			}

			transactionData := appTypes.TransferPipes{
				Payload:   &transactionPayload,
				Signature: signature,
			}
			transactionDataBytes, err := proto.Marshal(&transactionData)
			if err != nil {
				fmt.Println("Something Went Wrong While Marshaling Transaction Data: ", err)
				return
			}
			fmt.Println(transactionDataBytes)
			fmt.Println(client.BroadcastTxCommit(cmd.Context(), transactionDataBytes))
		},
	}
	sendTransactionCmd.Flags().BytesHexP("receiverPubKey", "r", []byte{}, "Public Key Of Pipe Receiver")
	sendTransactionCmd.Flags().Float64P("amount", "a", 0, "Amount Of Pipes To Transfer")
	sendTransactionCmd.Flags().BytesHexP("privateKey", "k", []byte{}, "Private Key To Sign Transaction Data")
	sendTransactionCmd.Flags().Uint32P("sequence", "s", 1, "Sequence Number Of Account, It Is Used For Uniqueness Of Transaction Data")

	queryDataCmd := &cobra.Command{
		Use:   "query",
		Short: "Query User Balance",
		Run: func(cmd *cobra.Command, args []string) {
			userPubKey, err := cmd.Flags().GetBytesHex("userPubKey")
			if err != nil {
				fmt.Println("Something Went Wrong While Extracting User Public Key, Make Sure You Pass Valid Public Key: ", err)
				return
			}
			if len(userPubKey) != 32 {
				fmt.Println("Public Key Length Must Be 32")
				return
			}
			queryData := appTypes.QueryBalance{
				UserPubKey: userPubKey,
			}
			queryDataBytes, err := proto.Marshal(&queryData)
			if err != nil {
				fmt.Println("Something Went Wrong While Marshaling Query Data: ", err)
				return
			}
			resData, err := client.ABCIQuery(cmd.Context(), "", queryDataBytes)
			fmt.Println(resData, err)
			userBalance := &appTypes.Pipe{}
			proto.Unmarshal(resData.Response.Value, userBalance)
			fmt.Println("User Balance: ", userBalance.Value)
		},
	}
	queryDataCmd.Flags().BytesHexP("userPubKey", "p", []byte{}, "Public Key User")
	rootCmd.AddCommand(sendTransactionCmd, queryDataCmd)
	rootCmd.Execute()
}

/*
Master User PubKey:- 08ec2505b4b7b0ffc8bd06d33232779f2130024b71eba6066ffbfcd846a5620b (Hex)
User1 PubKey:-f01df39e0e4bed3b4527b27b50f23f42cc95aa233f3c1e90864e6ae90b167df6

go build && ./pipeCoinClient send -r=f01df39e0e4bed3b4527b27b50f23f42cc95aa233f3c1e90864e6ae90b167df6 -a=10  -k=7186f8fd52d0616a044facc94f9f61b11624b4ba745b812e8287f501598cbf5a08ec2505b4b7b0ffc8bd06d33232779f2130024b71eba6066ffbfcd846a5620b -s=1
go build && ./pipeCoinClient query -p 08ec2505b4b7b0ffc8bd06d33232779f2130024b71eba6066ffbfcd846a5620b
*/
