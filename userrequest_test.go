package main

import "testing"

func TestAnalyzeInputTextCaseStandard(t *testing.T) {
	// RequstTypeGetKotos case
	userReq := NewUserTextRequest()
	text := "GetKotos"
	if err := userReq.AnalyzeInputText(text); err != nil {
		t.Fatal("There shoud not be error", " Got error: ", err, "input text: ", text)
	}
	if userReq.Type != "GetKotos" {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
	// RequstTypeAddKoto case
	userReq = NewUserTextRequest()
	text = "AddKoto training"
	if err := userReq.AnalyzeInputText(text); err != nil {
		t.Fatal("There shoud not be error", " Got error: ", err, "input text: ", text)
	}
	if userReq.Type != "AddKoto" {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if userReq.VariableKeyword != "training" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
	// RequstTypeGetActivities case
	userReq = NewUserTextRequest()
	text = "GetActivities"
	if err := userReq.AnalyzeInputText(text); err != nil {
		t.Fatal("There shoud not be error", " Got error: ", err, "input text: ", text)
	}
	if userReq.Type != "GetActivities" {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
	// RequstTypeGetActivities case 2
	userReq = NewUserTextRequest()
	text = "GetActivities training"
	if err := userReq.AnalyzeInputText(text); err != nil {
		t.Fatal("There shoud not be error", " Got error: ", err, "input text: ", text)
	}
	if userReq.Type != "GetActivities" {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if userReq.VariableKeyword != "training" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
	// RequstTypeAddActivity case
	userReq = NewUserTextRequest()
	text = "AddActivity training"
	if err := userReq.AnalyzeInputText(text); err != nil {
		t.Fatal("There shoud not be error", " Got error: ", err, "input text: ", text)
	}
	if userReq.Type != "AddActivity" {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if userReq.VariableKeyword != "training" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
}
