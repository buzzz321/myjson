package main

import (
	"fmt"
	"io/ioutil"
)

// Parser for json parsing ;)
type Parser struct {
	data    string
	currPos uint64
}

func (p Parser) parse() map[string]string {
	var retVal = make(map[string]string)
	return retVal
}

func (p Parser) consumeWhiteSpace() {

}

func (p Parser) peek() string {
	return ""
}

func main() {
	dat, err := ioutil.ReadFile("simple.json")

	if err != nil {
		panic(err)
	}
	str := string(dat[:len(dat)])
	fmt.Println("fsp = ", str)
}
