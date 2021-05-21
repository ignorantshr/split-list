package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/scanner"
)

var file_path = flag.String("f", "starts.txt", "The path to the list file")
var splitor = flag.String("F", ",", "The separator between the output names")
var help = flag.Bool("h", false, "Print this and exit")

func read_line(fp string) ([]string, error) {
	f, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines := make([]string, 0)

	/* reader := bufio.NewReader(f)
	for {
		l, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}else {
				return nil, err
			}
		}
		lines = append(lines, l)
	} */

	s := scanner.Scanner{}
	s.Init(f)
	for token := s.Scan(); token != scanner.EOF; token = s.Scan() {
		text := s.TokenText()
		lines = append(lines, text)
	}
	return lines, nil
}

func main() {
	flag.Parse()
	if *help {
		flag.Usage()
		return
	}
	lines, err := read_line(*file_path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(strings.Join(lines, *splitor))
}
