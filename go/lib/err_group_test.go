package lib

import (
	"errors"
	"golang.org/x/sync/errgroup"
	"testing"
)

func TestBasicUse(t *testing.T) {
	var eg errgroup.Group

	eg.Go(func() error {
		for i := 0; i < 1e10; i++ {
			i += 100
		}
		return nil
	})

	eg.Go(func() error {
		for i := 0; i < 1e10; i++ {
			if i == 1000 {
				return errors.New("asfafsf")
			}
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		t.Error(err)
	}
}
