package testutil

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
)

type testFn func() (bool, error)

const (
	baseWait = 1 * time.Millisecond
	maxWait  = 100 * time.Millisecond
)

func WaitForResult(try testFn) error {
	var err error
	wait := baseWait
	for retries := 100; retries > 0; retries-- {
		var success bool
		success, err = try()
		if success {
			time.Sleep(25 * time.Millisecond)
			return nil
		}

		time.Sleep(wait)
		wait *= 2
		if wait > maxWait {
			wait = maxWait
		}
	}
	if err != nil {
		return errors.Wrap(err, "timed out with error")
	}
	return fmt.Errorf("timed out")
}
