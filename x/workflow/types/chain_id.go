package types

const (
	ChainIDEth     = "1"
	ChainIDPolygon = "137"
)

var SupportedChainIDs = chainIDs{
	ChainIDEth:     {},
	ChainIDPolygon: {},
}

type chainIDs map[string]struct{}

func (a chainIDs) IsSupported(chainID string) bool {
	_, ok := a[chainID]
	return ok
}
