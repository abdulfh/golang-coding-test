package services

import (
	"fmt"
	"strconv"
)

type GeneralServices interface {
	FizzBuzz() (dataOutput []string, err error)
	Multiple() (dataOutput int, err error)
	MarkPaid(billAmount string) (dataOutput []int, err error)
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

func (services *generalServices) MarkPaid(billAmount string) (dataOutput []int, err error) {
	arrayInt := []int{1000, 3000, 1000, 5000, 10000}
	convertToInt, err := strconv.Atoi(billAmount)
	if err != nil {
		return nil, err
	}

	for indexOne := 0; indexOne < len(arrayInt); indexOne++ {
		if convertToInt == arrayInt[indexOne] {
			dataOutput = append(dataOutput, indexOne)
			return dataOutput, nil
		} else {
			for indexSecond := indexOne + 1; indexSecond < len(arrayInt); indexSecond = indexSecond + 1 {
				if arrayInt[indexOne]+arrayInt[indexSecond] == convertToInt {
					dataOutput = append(dataOutput, indexOne)
					dataOutput = append(dataOutput, indexSecond)
				} else {
					for indexThird := indexSecond + 1; indexThird < len(arrayInt); indexThird = indexThird + 1 {
						if arrayInt[indexOne]+arrayInt[indexSecond]+arrayInt[indexThird] == convertToInt {
							dataOutput = append(dataOutput, indexOne)
							dataOutput = append(dataOutput, indexSecond)
							dataOutput = append(dataOutput, indexThird)
							return dataOutput, nil
						}
					}
				}
			}
		}
	}

	return dataOutput, nil
}
