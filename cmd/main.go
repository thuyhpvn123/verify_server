package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func generatePhoneNumber() string {

	phone := "09"

	for i := 0; i < 8; i++ {
		digit := rand.Intn(10) // từ 0 đến 9
		phone += fmt.Sprintf("%d", digit)
	}
	return phone
}

func main() {

	rand.Seed(time.Now().UnixNano())
	const count = 10

	file, err := os.Create("phones.txt")
	if err != nil {
		fmt.Println("error when create file: ", err)
		return
	}
	defer file.Close()

	fmt.Println("Writing into phones.txt")

	for i := 0; i < count; i++ {
		phone := generatePhoneNumber()
		_, err := file.WriteString(phone + "\n")

		if err != nil {
			fmt.Println("error when writing file: ", err)
			return
		}
	}
	fmt.Println("Finished")
}
