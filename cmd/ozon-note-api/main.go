package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/ozoncp/ocp-note-api/internal/utils"
)

func main() {
	fmt.Println("© Oleg Kozyrev, 2021")

	var numberOfIterations int

	fmt.Println("Enter the number of iterations ...")
	fmt.Scanf("%d", &numberOfIterations)

	for i := 0; i < numberOfIterations; i++ {

		fmt.Printf("%d iteration\n", i)

		handle := func() {
			f, err := os.Open("test.txt")
			checkError(err)

			data := make([]byte, 20)

			count, err := f.Read(data)
			checkError(err)

			fmt.Printf("%d bytes: %s\n", count, string(data[:count]))

			time.Sleep(2 * time.Second)

			defer func() {
				fmt.Println("File closed")
				f.Close()
			}()
		}

		handle()
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
