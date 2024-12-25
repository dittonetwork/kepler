package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	epoch "cosmossdk.io/x/epochs/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	middlewareC "kepler/cmd/worker/contracts"
)

var (
	hostGRPC          string
	addressMiddleware string
	endpoint          string
	lastSaveEpoch     int64 = 0
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	hostGRPC = os.Getenv("HOST_GRPC")
	if hostGRPC == "" {
		log.Fatal("HOST_GRPC is not set in the environment")
	}

	addressMiddleware = os.Getenv("MIDDLEWARE_HOLESKY")
	if addressMiddleware == "" {
		log.Fatal("MIDDLEWARE_HOLESKY is not set in the environment")
	}

	endpoint = os.Getenv("ETH_ENDPOINT")
	if endpoint == "" {
		log.Fatal("ETH_ENDPOINT is not set in the environment")
	}

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				lastEpoch := getCurrentEpoch(ctx)
				if lastEpoch > lastSaveEpoch {
					lastSaveEpoch = lastEpoch
					getValidatorSet(ctx)
				}
			case <-ctx.Done():
				fmt.Println("Shutting down worker...")
				return
			}
		}
	}()

	sigReceived := <-stopCh
	fmt.Printf("Stop with signal %s", sigReceived)

	cancel()
}

func getCurrentEpoch(ctx context.Context) int64 {
	conn, err := grpc.Dial(hostGRPC, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed connect to server: %v", err)
	}
	defer conn.Close()

	c := epoch.NewQueryClient(conn)
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	newDat, err := c.CurrentEpoch(ctx, &epoch.QueryCurrentEpochRequest{Identifier: "minute"})
	if err != nil {
		log.Printf("failed get current epoch: %v", err)
	}

	log.Printf("Current epoch: %d", newDat.CurrentEpoch)

	return newDat.CurrentEpoch
}

func getValidatorSet(ctx context.Context) {
	client, err := ethclient.DialContext(ctx, endpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	instance, err := middlewareC.NewContracts(common.HexToAddress(addressMiddleware), client)
	if err != nil {
		log.Fatal(err)
	}

	validatorSet, err := instance.GetValidatorSetWithUnbonding(&bind.CallOpts{})
	if err != nil {
		log.Println(err)
		return
	}

	for _, validator := range validatorSet {
		log.Println("Validator address: ", validator.Operator)
		log.Println("Validator key: ", validator.Key)
		for _, vault := range validator.VaultData {
			log.Println("Vault:", vault.Vault.String())
			log.Println("Stake:", vault.Stake.String())
			log.Println("Power expires at:", vault.PowerExpiresAt.String())
			log.Println("Unbonding time:", vault.UnbondingTime.String())
			log.Println("Unbonding amount:", vault.UnbondingAmount.String())
		}
	}
}
