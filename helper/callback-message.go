package helper

import (
	"fmt"
)

func Check(message string, err error) {
	if err != nil {
		fmt.Println(message, err)
	}
}
