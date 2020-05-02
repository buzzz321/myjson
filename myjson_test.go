package main

import "testing"

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

func TestParseQuotedString(t *testing.T){
	uat := Parser{
		data:    " \"hej och hopp\" ",
		currPos: 0}

	
		uat.consumeWhiteSpace()
		result :=uat.parseQuotedString()
	
	if result == "hej och hopp" {
		t.Logf("consumeWhiteSpace success, expected %v, got ->%v<-", "hej och hopp", result)
	} else {
		t.Errorf("consumeWhiteSpace failed, expected %v, got ->%v<-", "hej och hopp", result)
	}
}