package main

import (
	"fmt"
	"io/ioutil"
)

// Parser for json parsing ;)
type Parser struct {
	data    string
	currPos int
}

func (p Parser) isDigit(ch string) bool {
	r := rune(ch[0])
	if int(r-'0') <= 9 {
		return true
	}

	return false
}

func (p *Parser) consume(token string) {

	p.currPos += len(token)

	if p.currPos > len(p.data) {
		p.currPos = len(p.data) - 1
	}
}

func (p *Parser) consumeWhiteSpace() {
	for p.currPos < len(p.data) {
		ch := p.data[p.currPos]
		if ch != ' ' && ch != '\n' &&
			ch != '\t' && ch != '\v' &&
			ch != '\r' {
			break
		}
		p.currPos++
	}
}

func (p *Parser) parseQuotedString() string{
	p.consumeWhiteSpace()
	p.consume("\"")
	startpos := p.currPos
	endFound := false	

	for p.currPos < len(p.data) {
		if p.data[p.currPos] == '"'{
			endFound = true
			break
		}
		p.currPos++
	}
	if endFound == true {
		p.consume("\"")
		return string([]rune(p.data)[startpos : p.currPos-1])
	}

	return ""
}

func (p *Parser) parseNumber() string {
	p.consumeWhiteSpace()

	startpos := p.currPos
	endFound := false

	for p.currPos < len(p.data) {
		ch := p.data[p.currPos]
		if ch == '.' || ch == '-' {
			p.currPos++
			continue
		}
		if ch < '0' || ch > '9' {
			break
		}
		p.currPos++
		endFound = true
	}

	if endFound == true {
		return string([]rune(p.data)[startpos:p.currPos])
	}
	return ""
}

func (p Parser) peek() string {
	return string([]rune(p.data)[p.currPos])
}

func (p Parser) peekNext() string {
	if p.currPos+1 > len(p.data) {
		return ""
	}
	return string([]rune(p.data)[p.currPos+1])
}

func (p *Parser) parseObject() {
	p.consume("{")
	p.consumeWhiteSpace()
	//p.consumeQuotedString()
	//p.consume(":")
	p.consumeWhiteSpace()
}

func (p *Parser) parse() map[string]string {
	var retVal = make(map[string]string)
	p.consumeWhiteSpace()

	switch ch := p.peek(); ch {
	case "{":
		p.parseObject()
	default:
		fmt.Printf("unknown type %v \n", ch)
	}
	return retVal
}

func main() {
	dat, err := ioutil.ReadFile("simple.json")

	if err != nil {
		panic(err)
	}
	str := string(dat[:len(dat)])
	fmt.Println("fsp = ", str)
}
