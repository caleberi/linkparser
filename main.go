package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	lib "projects/linkparser/lib"
)

func main() {
	htmlFile := flag.String("filepath", "./examples/ex1.html", "html file to parse")
	flag.Parse()
	cwd, err := os.Getwd()
	lib.Check("Error getting current working directories ....", err)
	p := filepath.Join(cwd, *htmlFile)
	f, err := os.Open(p)
	lib.Check(fmt.Sprintf("Error opening path : %s", p), err)
	defer f.Close()
	reader := bufio.NewReader(f)
	links, err := lib.Parse(reader)
	lib.Check("Error occurred while parsing html file", err)
	linkJSON, err := json.Marshal(links)
	lib.Check("Error marshaling struct links data to JSON mapping", err)
	out, err := os.Create("./out.json")
	lib.Check("Error opening output file", err)
	out.Write(linkJSON)
	out.Close()

}
