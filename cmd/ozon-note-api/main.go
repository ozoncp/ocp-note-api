package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ozoncp/ocp-note-api/core/alarmer"
	"github.com/ozoncp/ocp-note-api/core/note"
	"github.com/ozoncp/ocp-note-api/core/saver"
	_ "github.com/ozoncp/ocp-note-api/internal/utils"
)

func main() {
	fmt.Println("© Oleg Kozyrev, 2021")

	al := alarmer.New(5 * time.Second)
	sv := saver.New(5, nil, al, true)

	al.Init()
	sv.Init()

	var i int = 0
	for {
		sv.Save(note.Note{})
		time.Sleep(2 * time.Second)
		i++

		if i == 5 {
			sv.Close()
			break
		}
	}

	time.Sleep(2 * time.Second)
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
