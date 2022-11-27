package _go

import (
	"fmt"
)

func Go() {
	go func() { // want `detected use of go keyword: no details provided`
		fmt.Println("go test yourself")
	}()
}
