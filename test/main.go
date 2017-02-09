package main

import (
	"fmt"
	"os"

	"github.com/gerbenjacobs/docserv"
	"github.com/gerbenjacobs/docserv/test/static" // generated by running `./build.sh run`
)

func main() {
	data := map[string][]byte{
		"README.md":    static.MustAsset("README.md"),
		"test/TEST.md": static.MustAsset("test/TEST.md"),
	}
	ds := docserv.NewStaticDocServ(data)
	err := ds.Run()
	if err != nil {
		fmt.Println("Got error?:", err)
		os.Exit(1)
	}
}
