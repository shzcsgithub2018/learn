package err

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"testing"
	"time"
)

func TestErrGroup(t *testing.T) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 12*time.Second)
	defer cancelFunc()
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		time.Sleep(10 * time.Second)
		t.Log("10s")
		return errors.New("AAAAA")
	})

	eg.Go(func() error {
		time.Sleep(1 * time.Second)
		t.Log("1s")
		return nil
	})

	if goErr := eg.Wait(); goErr != nil {
		t.Error(goErr)
	}
}
