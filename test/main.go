package main

import (
	"fmt"
	"github.com/gerbenjacobs/docserv"
	"os"
)

func main() {
	ds := docserv.NewDocServ([]string{"README.md", "test/TEST.md"})
	err := ds.Run()
	if err != nil {
		fmt.Println("Got error?:", err)
		os.Exit(1)
	}
}
