package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

func readFile(path string, query string) {
	file, ok := os.Open(path)
	if ok != nil {
		fmt.Println("Failed to open:", path)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for i := 1; scanner.Scan(); i++ {
		if strings.Contains(scanner.Text(), query) {
			fmt.Printf("%s:%d: %s\n", path, i, scanner.Text())
		}
	}

}

type Config struct {
	Component struct {
		Name     string
		Files    []string
		Patterns []string
	}
}

func main() {
	flag.Parse()
	filename := flag.Arg(0)
	var config []Config
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		panic(err)
	}

	for _, value := range config {
		fmt.Println("---", value.Component.Name, "---")
		for _, filename := range value.Component.Files {
			for _, pattern := range value.Component.Patterns {
				readFile(filename, pattern)
			}
		}
	}
}
