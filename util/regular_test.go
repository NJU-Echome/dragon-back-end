package util

import (
	"fmt"
	"testing"
)

func TestMatchPhone(t *testing.T) {
	if MatchPhoneNum("1595057277") {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
