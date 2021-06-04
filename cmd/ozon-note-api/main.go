package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ozoncp/ocp-note-api/core/alarmer"
	"github.com/ozoncp/ocp-note-api/core/saver"
	_ "github.com/ozoncp/ocp-note-api/internal/utils"
)

func main() {
	fmt.Println("© Oleg Kozyrev, 2021")

	al := alarmer.New(2 * time.Second)
	sv := saver.New(al)

	al.Init()
	sv.Init()

	time.Sleep(100 * time.Second)

}

func repeatOpenFile() {
	var numberOfIterations int

	fmt.Println("Enter the number of iterations ...")
	fmt.Scanf("%d", &numberOfIterations)

	for i := 0; i < numberOfIterations; i++ {

		fmt.Printf("%d iteration\n", i)

		func() {
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
		}()
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
