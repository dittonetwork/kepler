root = x/symbiotic/abi

.PHONY: gen-contracts
gen_contracts: $(root)/middleware_contract.go

$(root)/middleware_contract.go: $(root)/Middleware.abi.json
	cd $(root) && go generate
