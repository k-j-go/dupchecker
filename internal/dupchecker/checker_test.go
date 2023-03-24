package dupchecker

import (
	"context"
	"fmt"
	"testing"
)

func Test_Dup(t *testing.T) {
	dup := New[byte](0)
	generator := func(input string) chan byte {
		c := make(chan byte)
		go func(c chan byte, input string) {
			defer close(c)
			for _, v := range []byte(input) {
				c <- v
			}
		}(c, input)
		return c
	}
	fmt.Printf("duplicate found: %t \n", dup.fromSlice([]byte("abcdef")))
	b, _ := dup.fromChannel(context.TODO(), generator("vpyuytrwn"))
	fmt.Printf("duplicate found: %t \n", b)
	fmt.Printf("%d \n", dup.size())
}
