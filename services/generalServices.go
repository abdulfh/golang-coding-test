package services

import (
	"fmt"
	"strconv"
)

type GeneralServices interface {
	FizzBuzz() (dataOutput []string, err error)
	Multiple() (dataOutput int, err error)
}

type generalServices struct{}

func InitGeneralServices() *generalServices {
	return &generalServices{}
}

func (services *generalServices) FizzBuzz() (dataOutput []string, err error) {
	for i := 1; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			dataOutput = append(dataOutput, "FizzBuzz")
			// Jika  tidak Habis dibagi 3 atau 5 hasilkan angka
		} else if i%3 == 0 {
			// Jika Habis dibagi 3 hasilkan fizz
			dataOutput = append(dataOutput, "Fizz")
		} else if i%5 == 0 {
			// Jika Habis dibagi 3 hasilkan buzz
			dataOutput = append(dataOutput, "Buzz")
		} else {
			// Print angka
			dataOutput = append(dataOutput, strconv.Itoa(i))
		}

	}
	return dataOutput, nil
}

func (services *generalServices) Multiple() (dataOutput int, err error) {
	for i := 1; i <= 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			fmt.Println(i)
			dataOutput += i
		}
	}

	return dataOutput, nil
}
