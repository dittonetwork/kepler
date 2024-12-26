package lockup

import (
	"cosmossdk.io/x/accounts/accountstd"
)

func ProvideAllLockupAccounts() []accountstd.DepinjectAccount {
	return []accountstd.DepinjectAccount{
		ProvidePeriodicLockingAccount(),
		ProvideContinuousLockingAccount(),
		ProvidePermanentLockingAccount(),
		ProvideDelayedLockingAccount(),
	}
}

func ProvideContinuousLockingAccount() accountstd.DepinjectAccount {
	return accountstd.DIAccount(CONTINUOUS_LOCKING_ACCOUNT, NewContinuousLockingAccount)
}

func ProvidePeriodicLockingAccount() accountstd.DepinjectAccount {
	return accountstd.DIAccount(PERIODIC_LOCKING_ACCOUNT, NewPeriodicLockingAccount)
}

func ProvideDelayedLockingAccount() accountstd.DepinjectAccount {
	return accountstd.DIAccount(DELAYED_LOCKING_ACCOUNT, NewDelayedLockingAccount)
}

func ProvidePermanentLockingAccount() accountstd.DepinjectAccount {
	return accountstd.DIAccount(PERMANENT_LOCKING_ACCOUNT, NewPermanentLockingAccount)
}
