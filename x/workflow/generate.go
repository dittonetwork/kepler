package job

//go:generate  mockgen --source=keeper/keeper.go --destination=types/mock/keeper.go --package=mock
//go:generate  mockgen --source=types/expected_keepers.go --destination=types/mock/expected_keepers.go --package=mock
