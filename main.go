package main

import (
	"fmt"
	"os"
	"os/signal"
	"pipeCoin/app"
	"syscall"

	abciServer "github.com/cometbft/cometbft/abci/server"
	db "github.com/cosmos/cosmos-db"
)

func main() {
	fmt.Println("Chian Is Starting...")
	dbIns, err := db.NewGoLevelDB("pipeCoin", "pipeCoinData", nil)
	if err != nil {
		panic(fmt.Sprint("Something Went Wrong While Starting Database: ", err))
	}
	app := app.NewApp(dbIns)
	server := abciServer.NewSocketServer(":26658", app)
	server.Start()
	defer server.Stop()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
