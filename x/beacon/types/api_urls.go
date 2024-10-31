package types

import "kepler/pkg/common"

var defaultBeaconUrls = []string{
	"https://eth-sepolia-beacon.public.blastapi.io",
	"http://unstable.sepolia.beacon-api.nimbus.team",
	"https://ethereum-sepolia-beacon-api.publicnode.com",
}

func NewBeaconApiUrls() *common.ApiUrls {
	return common.NewApiUrls(
		common.GetUrlsFromEnv("BEACON_API_URLS", defaultBeaconUrls),
	)
}
