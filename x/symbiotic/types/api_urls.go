package types

import "kepler/pkg/common"

var defaultRpcUrls = []string{
	"https://rpc.ankr.com/eth_sepolia",
	"https://ethereum-sepolia.blockpi.network/v1/rpc/public",
	"https://eth-sepolia.public.blastapi.io",
	"https://sepolia.gateway.tenderly.co",
}

func NewRpcUrls() *common.ApiUrls {
	return common.NewApiUrls(
		common.GetUrlsFromEnv("ETH_RPC_URLS", defaultRpcUrls),
	)
}
