
# vartPipeChain

### To run chain

#### (1) Initialise CometBFT consensus engine 

```
go run github.com/cometbft/cometbft/cmd/cometbft@v0.37.0 init

```
#### (2) Run CometBFT consensus engine 
```
go run github.com/cometbft/cometbft/cmd/cometbft@v0.37.0 node

```
### (3) Run ABCI Application in different terminal 
```
  go run .
```


### To interect with chain 
```
cd client 
./pipeCoinClient --help

```
For more  [info](https://github.com/cometbft/cometbft/blob/main/docs/guides/go.md)

