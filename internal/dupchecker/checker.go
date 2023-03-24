package dupchecker

import (
	"context"
	"errors"
)

type DupChecker[T comparable] map[T]struct{}

func New[T comparable](count int) *DupChecker[T] {
	m := make(map[T]struct{}, count)
	return (*DupChecker[T])(&m)
}

func (C *DupChecker[T]) add(t T) error {
	if _, ok := (*C)[t]; !ok {
		(*C)[t] = struct{}{}
		return nil
	} else {
		return errors.New("dup")
	}
}

func (C *DupChecker[T]) size() int {
	return len(*C)
}

func (C *DupChecker[T]) fromSlice(t []T) bool {
	for _, v := range t {
		if err := C.add(v); err == nil {
			continue
		} else {
			return true
		}
	}
	return false
}

func (C *DupChecker[T]) fromChannel(ctx context.Context, c chan T) (bool, error) {
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
