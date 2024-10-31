package keeper

import "time"

func Retry(operation func() error, maxRetries int, delay time.Duration) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		err = operation()
		if err == nil {
			return nil
		}
		time.Sleep(delay)
	}
	return err
}
