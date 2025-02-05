package helper

import (
	"fmt"
	"log"
)

func Check(message string, err error) {
	if err != nil {
		fmt.Println(message)
		log.Fatal(err)
	}
}
