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
		data:    "{ \"barr\": [1,2,3,4,5,6] }",
		currPos: 0}

	var result jsonType
	result = uat.parseObject()

	t.Logf("%v", result)
	if result.jtype == JObject {
		t.Logf("TestParseArray success, expected %v, got ->%v<-", "JObject", result)
	} else {
		t.Errorf("TestParseArray failed, expected %v, got ->%v<-", "JObject", result)
	}

	if result.key == "barr" {
		t.Logf("TestParseArray success, expected %v, got ->%v<-", "barr", result)
	} else {
		t.Errorf("TestParseArray failed, expected %v, got ->%v<-", "barr", result)
	}

	if len(result.objects) == 1 {
		t.Logf("TestParseArray success, expected %v, got ->%v<-", "1", result)
	} else {
		t.Errorf("TestParseArray failed, expected %v, got ->%v<-", "1", result)
	}

	if result.objects[0].jtype == JArray {
		t.Logf("TestParseArray success, expected %v, got ->%v<-", "JArray", result)
	} else {
		t.Errorf("TestParseArray failed, expected %v, got ->%v<-", "JArray", result)
	}

	if len(result.objects[0].objects) == 6 {
		t.Logf("TestParseArray success, expected %v, got ->%v<-", "6", result)
	} else {
		t.Errorf("TestParseArray failed, expected %v, got ->%v<-", "6", result)
	}

	data := 1

	for _, actual := range result.objects[0].objects {
		i, _ := strconv.Atoi(actual.value)
		if i == data {
			t.Logf("TestParseArray success, expected %v, got ->%v<-", data, i)
		} else {
			t.Errorf("TestParseArray failed, expected %v, got ->%v<-", data, i)
		}
		data++
	}
}
