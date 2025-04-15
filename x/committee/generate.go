package committee

//go:generate mockgen --source=types/expected_keepers.go  --destination=testutil/expected_keepers_mocks.go --package=testutil
//go:generate mockgen --source=types/repository.go --destination=testutil/repository_mocks.go --package=testutil
