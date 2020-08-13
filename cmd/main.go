package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Chris-Ju/Identify-WebScanTools/identify"
)

func main() {
	buf, err := ioutil.ReadFile("flow")
	if err != nil {
		log.Fatal(err)
	}
	tool, err := identify.Run(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tool)

}
