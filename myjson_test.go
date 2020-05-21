package main

import (
	"strconv"
	"testing"
)

func TestConsume(t *testing.T) {
	uat := Parser{
		data:    "    .",
		currPos: 0}

	uat.consumeWhiteSpace()

	result := uat.peek()
	if result == "." {
		t.Logf("consumeWhiteSpace success, expected %v, got ->%v<-", ".", result)
	} else {
		t.Errorf("consumeWhiteSpace failed, expected %v, got ->%v<-", ".", result)
	}
}

func TestConsumeMulti(t *testing.T) {
	uat := Parser{
		data:    "   \n *",
		currPos: 0}

	uat.consumeWhiteSpace()

	result := uat.peek()
	if result == "*" {
		t.Logf("consumeWhiteSpace success, expected %v, got ->%v<-", "*", result)
	} else {
		t.Errorf("consumeWhiteSpace failed, expected %v, got ->%v<-", "*", result)
	}
}

func TestParseQuotedString(t *testing.T) {
	uat := Parser{
		data:    " \"hej och hopp\" ",
		currPos: 0}

	uat.consumeWhiteSpace()
	result := uat.parseQuotedString()

	if result == "hej och hopp" {
		t.Logf("Y success, expected %v, got ->%v<-", "hej och hopp", result)
	} else {
		t.Errorf("Y failed, expected %v, got ->%v<-", "hej och hopp", result)
	}
}

func TestParseNumbers(t *testing.T) {
	uat := Parser{
		data:    " 123.4 ",
		currPos: 0}

	result := uat.parseNumber()

	if result == "123.4" {
		t.Logf("TestParseNumbers success, expected %v, got ->%v<-", "123.4", result)
	} else {
		t.Errorf("TestParseNumbers failed, expected %v, got ->%v<-", "123.4", result)
	}
}

func TestParseArray(t *testing.T) {
	uat := Parser{
		data:    " [1,2,3,4,5,6] ",
		currPos: 0}

	result := uat.parseArray()

	//t.Errorf("%v", result)
	data := 1
	for _, actual := range result.arr {
		i, _ := strconv.Atoi(actual.value)
		if i == data {
			t.Logf("TestParseArray success, expected %v, got ->%v<-", data, i)
		} else {
			t.Errorf("TestParseArray failed, expected %v, got ->%v<-", data, i)
		}
		data++
	}

}

func TestParseArrayInObject(t *testing.T) {
	uat := Parser{
		data:    " {\"plura\": [1,2,3,4,5,6] }",
		currPos: 0}
	//måste ju se till att spara ner plura med någonstans
	result := uat.parseObject()

	//t.Errorf("%v", result)

	values, ok := result.get("plura")
	if ok {
		t.Logf("TestParseArray success, expected %v, got ->%v<-", true, ok)
	} else {
		t.Errorf("TestParseArray failed, expected %v, got ->%v<-", true, ok)
	}

	data := 1
	for _, actual := range values.arr {
		i, _ := strconv.Atoi(actual.value)
		if i == data {
			t.Logf("TestParseArray success, expected %v, got ->%v<-", data, i)
		} else {
			t.Errorf("TestParseArray failed, expected %v, got ->%v<-", data, i)
		}
		data++
	}

}

func TestObject(t *testing.T) {
	uat := Parser{
		data: `{
		"red": 14772,
		"blue": 16523,
		"green": 16614,
		"type": "Center"} `,
		currPos: 0}

	result := uat.parseObject()

	//t.Errorf("TestObject, expected %v, got ->%v<-", "*", result)
	if result.objects["red"].jtype == JNumber {
		t.Logf("TestObject success, expected %v, got ->%v<-", JNumber, result.objects["red"].jtype)
	} else {
		t.Errorf("TestObject failed, expected %v, got ->%v<-", JNumber, result.objects["red"].jtype)
	}

	if result.objects["red"].value == "14772" {
		t.Logf("TestObject success, expected %v, got ->%v<-", "14772", result.objects["red"].value)
	} else {
		t.Errorf("TestObject failed, expected %v, got ->%v<-", "14772", result.objects["red"].value)
	}
}
