package types

import (
	"os"
	"strings"
)

type ApiUrls struct {
	beaconApiUrls   []string
	ethApiUrls      []string
	currentBeaconId int
	currentEthId    int
}

func NewApiUrls() ApiUrls {
	// USE ONLY YOUR LOCAL BEACON CLIENT FOR SAFETY!!!
	beaconApiUrls := strings.Split(os.Getenv("BEACON_API_URLS"), ",")
	if len(beaconApiUrls) == 1 && beaconApiUrls[0] == "" {
		beaconApiUrls[0] = "https://eth-sepolia-beacon.public.blastapi.io"
		beaconApiUrls = append(beaconApiUrls, "http://unstable.sepolia.beacon-api.nimbus.team")
		beaconApiUrls = append(beaconApiUrls, "https://ethereum-sepolia-beacon-api.publicnode.com")
	}

	ethApiUrls := strings.Split(os.Getenv("ETH_API_URLS"), ",")

	if len(ethApiUrls) == 1 && ethApiUrls[0] == "" {
		ethApiUrls[0] = "https://rpc.ankr.com/eth_sepolia"
		ethApiUrls = append(ethApiUrls, "https://ethereum-sepolia.blockpi.network/v1/rpc/public")
		ethApiUrls = append(ethApiUrls, "https://eth-sepolia.public.blastapi.io")
		ethApiUrls = append(ethApiUrls, "https://sepolia.gateway.tenderly.co")
	}

	return ApiUrls{beaconApiUrls: beaconApiUrls, ethApiUrls: ethApiUrls}
}

func (au ApiUrls) GetEthApiUrl() string {
	return au.ethApiUrls[au.currentEthId]
}

func (au ApiUrls) GetBeaconApiUrl() string {
	return au.beaconApiUrls[au.currentBeaconId]
}

func (au *ApiUrls) RotateEthUrl() {
	au.currentEthId++
	if au.currentEthId == len(au.ethApiUrls) {
		au.currentEthId = 0
	}
}

func (au *ApiUrls) RotateBeaconUrl() {
	au.currentBeaconId++
	if au.currentBeaconId == len(au.beaconApiUrls) {
		au.currentBeaconId = 0
	}
}
