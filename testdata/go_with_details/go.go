package _go

import (
	"fmt"
)

func Go() {
	go func() { // want `detected use of go keyword: details details details`
		fmt.Println("go test yourself")
	}()
}
