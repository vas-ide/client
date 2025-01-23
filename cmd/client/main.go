package main

import (
	"fmt"
	"log"

	"github.com/vas-ide/storage/pkg/storage"
)

func main() {

	st := storage.NewStorage()

	file, err := st.Upload("VetGramm", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", file)
}
