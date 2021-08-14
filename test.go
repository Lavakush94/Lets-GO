package main

import (
	"errors"
	"log"
)

func Divide(a float32, b float32) (float32, error) {
	var result float32
	if b == 0 {
		return result, errors.New("can not divide by 0")
	}

	result = a / b
	return result, nil
}

func main() {
	d, err := Divide(10, 0)

	if err == nil {
		log.Println(d)
	} else {
		log.Println(err)
	}
}
