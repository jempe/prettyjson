package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/tidwall/pretty"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Buffer([]byte(""), 10000000)

	for scanner.Scan() {
		prettyJson := pretty.Pretty(scanner.Bytes())

		fmt.Println(string(prettyJson))
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
