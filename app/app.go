package app

import (
	"fmt"

	"github.com/cometbft/cometbft/crypto/ed25519"

	appTypes "pipeCoin/app/types"

	"github.com/cometbft/cometbft/abci/types"
	db "github.com/cosmos/cosmos-db"
	"google.golang.org/protobuf/proto"
)

type Pipe uint64

type App struct {
	types.BaseApplication
	db *db.GoLevelDB
}

func NewApp(db *db.GoLevelDB) *App {
	return &App{
		db: db,
	}
}

func (app *App) InitChain(req types.RequestInitChain) (res types.ResponseInitChain) {
	fmt.Println("Chain Is Initializing...")
	maxPipes, err := proto.Marshal(&MaxPipes)
	if err != nil {
		panic(fmt.Sprint("Something Went While Marshaling Pipes: ", err))
	}
	app.db.Set(MastersPublicKey, maxPipes)
	fmt.Print("Total Pipes: ", MaxPipes.Value, "\n", "All Pipes Issued To ", ed25519.PubKey(MastersPublicKey).String())
	return
}

func (app *App) CheckTx(req types.RequestCheckTx) (res types.ResponseCheckTx) {
	transferPipesTxData := appTypes.TransferPipes{}
	proto.Unmarshal(req.Tx, &transferPipesTxData)
	if transferPipesTxData.Payload.Amount.Value <= 0 {
		res.Code = 1
		res.Info = "Amount Must Be Greater Than Zero"
		res.Log = "Transaction Failed, Invalid Data"
		return
	}
	if len(transferPipesTxData.Payload.From) != 32 {
		res.Code = 2
		res.Info = "Invalid Sender Public Key"
		res.Log = "Transaction Failed, Invalid Data"
		return
	}
	if len(transferPipesTxData.Payload.To) != 32 {
		res.Code = 3
		res.Info = "Invalid Receiver Public Key"
		res.Log = "Transaction Failed, Invalid Data"
		return
	}
	if len(transferPipesTxData.Signature) != 64 {
		res.Code = 3
		res.Info = "Invalid Signature"
		res.Log = "Transaction Failed, Invalid Data"
		return
	}

	signatureMsgData, err := proto.Marshal(transferPipesTxData.Payload)
	if err != nil {
		res.Code = 4
		res.Info = "Invalid Transaction Payload"
		res.Log = "Transaction Failed, Invalid Data"
		return
	}
	if !ed25519.PubKey(transferPipesTxData.Payload.From).VerifySignature(signatureMsgData, transferPipesTxData.Signature) {
		res.Code = 4
		res.Info = "Signature Verification Failed"
		res.Log = "Transaction Failed, Signature Verification Failed"
		return
	}
	res.Code = 0
	res.Info = "Transaction Verification Done"
	fmt.Println("Transaction Is Successfully Verified")
	return
}

func (app *App) DeliverTx(req types.RequestDeliverTx) (res types.ResponseDeliverTx) {
	transferPipesTxData := appTypes.TransferPipes{}
	proto.Unmarshal(req.Tx, &transferPipesTxData)
	senderBalanceBytes, err := app.db.Get(transferPipesTxData.Payload.From)
	if err != nil {
		res.Code = 5
		res.Info = fmt.Sprint("Something Went Wrong While Fetching Sender Balance: ", err)
		return
	}
	senderBalance := appTypes.Pipe{}
	if len(senderBalanceBytes) == 0 {
		senderBalance.Value = 0
		senderBalanceBytes, err = proto.Marshal(&senderBalance)
		if err == nil {
			app.db.Set(transferPipesTxData.Payload.From, senderBalanceBytes)
		}
		res.Code = 6
		res.Info = "Sender Have Insufficient Balance"
		return
	}
	if err = proto.Unmarshal(senderBalanceBytes, &senderBalance); err != nil {
		res.Code = 7
		res.Info = fmt.Sprint("Something Went Wrong While UnMarshaling Sender Balance: ", err)
		return
	}

	receiverBalanceBytes, err := app.db.Get(transferPipesTxData.Payload.To)
	if err != nil {
		res.Code = 8
		res.Info = fmt.Sprint("Something Went Wrong While Fetching Receiver Balance: ", err)
		return
	}
	receiverBalance := appTypes.Pipe{}
	if len(receiverBalanceBytes) == 0 {
		receiverBalance.Value = 0
	} else {
		if err = proto.Unmarshal(receiverBalanceBytes, &receiverBalance); err != nil {
			res.Code = 9
			res.Info = fmt.Sprint("Something Went Wrong While UnMarshaling Receiver Balance: ", err)
			return
		}
	}

	if senderBalance.Value >= transferPipesTxData.Payload.Amount.Value {
		senderBalance.Value -= transferPipesTxData.Payload.Amount.Value
		receiverBalance.Value += transferPipesTxData.Payload.Amount.Value
	} else {
		res.Code = 10
		res.Info = "Sender Have Insufficient Balance"
		return
	}
	senderBalanceBytes, _ = proto.Marshal(&senderBalance)
	receiverBalanceBytes, _ = proto.Marshal(&receiverBalance)
	app.db.Set(transferPipesTxData.Payload.From, senderBalanceBytes)
	app.db.Set(transferPipesTxData.Payload.To, receiverBalanceBytes)

	fmt.Println("Transaction Is Delivered")
	return
}

func (app *App) Query(req types.RequestQuery) (res types.ResponseQuery) {
	queryData := appTypes.QueryBalance{}
	if err := proto.Unmarshal(req.Data, &queryData); err != nil {
		res.Code = 11
		res.Info = "Something Went Wrong While UnMarshaling Sender Balance"
		return
	}
	userBalance, err := app.db.Get(queryData.UserPubKey)
	if err != nil {
		res.Code = 12
		res.Info = fmt.Sprint("Something Went Wrong While Fetching User Balance: ", err)
		return
	}
	if len(userBalance) == 0 {
		userBalanceBytes, err := proto.Marshal(&appTypes.Pipe{Value: 0})
		if err == nil {
			app.db.Set(queryData.UserPubKey, userBalanceBytes)
		}
	}
	res.Code = 0
	res.Value = userBalance
	return
}
