package main

import (
	"fmt"
	"io/ioutil"
)

// JType ..
type JType int8

// JString JNumber JArray JObject ..
const (
	JString JType = iota
	JNumber
	JArray
	JObject
)

func (jtype JType) String() string {
	switch jtype {
	case JString:
		return "JString"
	case JNumber:
		return "JNumber"
	case JArray:
		return "JArray"
	case JObject:
		return "JObject"
	default:
		return fmt.Sprintf("%d", int(jtype))
	}
}

/*
{ key: [1,2,3,4,5,6] }
Value
{
	strValue string
	number string
	obj  object
	arr Array
}
Object
{
	objects map[string]string
}
Array
{
	arr [] Value
}
*/

// JSONValue Need to be exported to be printed
type JSONValue struct {
	jtype JType
	value string
	arr   []JSONValue
	obj   JSObject
}

// JSONType Need to be exported to be printed
type JSObject struct {
	objects map[string]JSONValue
}

func (o *JSObject) add(key string, value JSONValue) {
	o.objects[key] = value
}

func (o *JSObject) get(key string) (JSONValue, bool) {
	value, ok := o.objects[key]
	return value, ok
}

// Parser for json parsing ;)
type Parser struct {
	data    string
	currPos int
}

func (p Parser) isDigit(ch string) bool {
	r := rune(ch[0])
	isnum := int(r - '0')
	if isnum <= 9 && isnum >= 0 {
		return true
	}

	return false
}

func (p *Parser) consume(token string) bool {
	if p.data[p.currPos] != token[0] {
		return false
	}
	p.currPos += len(token)

	if p.currPos > len(p.data) {
		p.currPos = len(p.data) - 1
	}

	return true
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

func (p *Parser) parseQuotedString() string {
	p.consumeWhiteSpace()
	if !p.consume("\"") {
		return ""
	}
	startpos := p.currPos
	endFound := false

	for p.currPos < len(p.data) {
		if p.data[p.currPos] == '"' {
			endFound = true
			break
		}
		p.currPos++
	}
	if endFound == true {
		if !p.consume("\"") {
			panic("No end of string \" not found")
		}
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
			endFound = true
			break
		}
		p.currPos++
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

func (p *Parser) parseObject() JSObject {
	var retVal JSObject

	retVal.objects = make(map[string]JSONValue)

	p.consumeWhiteSpace()
	if !p.consume("{") {
		return retVal //no object to parse..
	}

	for true {
		p.consumeWhiteSpace()

		if p.consume("}") {
			break
		}

		key := p.parseQuotedString()

		if key == "" {
			fmt.Println("Empty object")
			continue // empty object
		} else {
			//fmt.Println(retVal.key)
		}
		p.consumeWhiteSpace()
		if !p.consume(":") {
			panic("no colon found")
		}
		p.consumeWhiteSpace()
		value := p.parseValue()
		retVal.add(key, value)

		p.consumeWhiteSpace()
		if !p.consume(",") {
			break
		}
	}

	return retVal
}

func (p *Parser) parseArray() JSONValue {
	var retVal JSONValue

	p.consumeWhiteSpace()
	if !p.consume("[") {
		fmt.Println("Error no array?!?")
		return retVal // no array to parse
	}

	for true {
		p.consumeWhiteSpace()
		if p.consume("]") {
			break
		}
		value := p.parseValue()

		retVal.arr = append(retVal.arr, value)

		p.consumeWhiteSpace()
		p.consume(",")
	}
	//fmt.Println("-->", retVal)
	return retVal
}

func (p *Parser) parseValue() JSONValue {
	p.consumeWhiteSpace()

	ch := p.peek()

	//fmt.Printf("--> will try to parse value for ch=%v\n", ch)
	if ch == "{" {
		return JSONValue{JObject, "", nil, p.parseObject()}
	} else if ch == "[" {
		return JSONValue{JArray, "", p.parseArray().arr, JSObject{}}
	} else if p.isDigit(ch) {
		return JSONValue{JNumber, p.parseNumber(), nil, JSObject{}}
	} else if ch == "\"" {
		return JSONValue{JString, p.parseQuotedString(), nil, JSObject{}}
	} else {
		fmt.Printf("unknown type %v \n", ch)
		panic("bailing out")
	}
}

func main() {
	dat, err := ioutil.ReadFile("simple.json")

	if err != nil {
		panic(err)
	}
	str := string(dat[:len(dat)])
	fmt.Println("fsp = ", str)
}
