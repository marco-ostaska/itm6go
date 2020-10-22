package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/marco-ostaska/itm6go/pkg/parsesitxml"
)

func main() {

	fFile := flag.String("file", "", "File or directory containing files to read")
	flag.Parse()

	g, err := filepath.Glob(*fFile + "/*.xml")

	if len(g) < 1 {
		err = fmt.Errorf("No such file")
	}
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range g {
		sit, err := parsesitxml.ParseSitXML(v)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(sit.ROW.SITNAME)

	}
}
