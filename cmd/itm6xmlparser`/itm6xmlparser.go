package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/marco-ostaska/itm6go/pkg/itm6xmlparse"
)

func main() {

	fFile := flag.String("file", "", "File or directory containing files to read")
	flag.Parse()

	g, err := filepath.Glob(*fFile + "/*.xml")

	if len(g) < 1 {
		err = fmt.Errorf("No such file ", *fFile)
	}
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range g {
		sit, err := itm6xmlparse.ParseSitXML(v)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("MSL: ", sit.ROW.NODELIST)
		if sit.ROW.NODE[0] == 10 {
			fmt.Println("Nodes: ", sit.ROW.NODE[1:])
		}

	}
}
