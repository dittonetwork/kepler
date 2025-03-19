package types

const (
	EthereumChainID = "1"
	PolygonChainID  = "137"
)

var SupportedChainIDs = chainIDs{
	EthereumChainID: {},
	PolygonChainID:  {},
}

type chainIDs map[string]struct{}

func (a chainIDs) IsSupported(chainID string) bool {
	_, ok := a[chainID]
	return ok
}
