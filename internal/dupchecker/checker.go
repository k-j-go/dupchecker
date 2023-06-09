package dupchecker

import (
	"context"
	"errors"
)

type DupCheckers[T comparable] map[T]struct{}

func New[T comparable](count int) *DupCheckers[T] {
	m := make(map[T]struct{}, count)
	return (*DupCheckers[T])(&m)
}

func (C *DupCheckers[T]) add(t T) error {
	if _, ok := (*C)[t]; !ok {
		(*C)[t] = struct{}{}
		return nil
	} else {
		return errors.New("dup")
	}
}

func (C *DupCheckers[T]) size() int {
	return len(*C)
}

func (C *DupCheckers[T]) fromSlice(t []T) bool {
	for _, v := range t {
		if err := C.add(v); err == nil {
			continue
		} else {
			return true
		}
	}
	return false
}

func (C *DupCheckers[T]) fromChannel(ctx context.Context, c chan T) (bool, error) {
	for {
		select {
		case <-ctx.Done():
			return true, errors.New("timeout")
		case t := <-c:
			if err := C.add(t); err != nil {
				return true, nil
			}
		}
	}
}
