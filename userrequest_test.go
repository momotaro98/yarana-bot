package main

import (
		"testing"
	)

func TestAnalyzeInputText_SingleCommand_HelpType(t *testing.T) {
	// Arrange
	text := "Help"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	if err != nil {
		t.Fatal("There should not be error", " Got error: ", err, "input text: ", text)
	}
	if reqType != RequestTypeHelp {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if keyword != "" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
}

func TestAnalyzeInputText_SingleCommand_GetKotosType(t *testing.T) {
	// Arrange
	text := "GetKotos"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	if err != nil {
		t.Fatal("There should not be error", " Got error: ", err, "input text: ", text)
	}
	if reqType != RequestTypeGetKotos {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if keyword != "" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
}

func TestAnalyzeInputText_CommandWithKeyword_AddKotoType(t *testing.T) {
	// Arrange
	text := "AddKoto training"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	if err != nil {
		t.Fatal("There should not be error", " Got error: ", err, "input text: ", text)
	}
	if reqType != RequestTypeAddKoto {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if keyword != "training" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
}

func TestAnalyzeInputText_SingleCommand_GetActivitiesType(t *testing.T) {
	// Arrange
	text := "GetActivities"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	if err != nil {
		t.Fatal("There should not be error", " Got error: ", err, "input text: ", text)
	}
	if reqType != RequestTypeGetActivities {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if keyword != "" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
}

func TestAnalyzeInputText_CommandWithKeyword_GetActivitiesAndKeyword(t *testing.T) {
	// Arrange
	text := "GetActivities training"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	if err != nil {
		t.Fatal("There should not be error", " Got error: ", err, "input text: ", text)
	}
	if reqType != RequestTypeGetActivities {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if keyword != "training" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
}

func TestAnalyzeInputText_CommandWithKeyword_AddActivityType(t *testing.T) {
	// Arrange
	text := "AddActivity training"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	if err != nil {
		t.Fatal("There should not be error", " Got error: ", err, "input text: ", text)
	}
	if reqType != RequestTypeAddActivity {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if keyword != "training" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
}

/*
func TestAnalyzeInputTextCaseStandard(t *testing.T) {
	// Command case
	// RequestTypeHelp case
	text := "Help"

	// Japanese case
	// RequestTypeHelp case
	userReq = NewUserTextRequest()
	text = "使い方"
	if err := userReq.AnalyzeInputText(text); err != nil {
		t.Fatal("There shoud not be error", " Got error: ", err, "input text: ", text)
	}
	if userReq.Type != "Help" {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
	// RequestTypeHelp case
	userReq = NewUserTextRequest()
	text = "使い方教えて"
	if err := userReq.AnalyzeInputText(text); err != nil {
		t.Fatal("There shoud not be error", " Got error: ", err, "input text: ", text)
	}
	if userReq.Type != "Help" {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
	// RequestTypeHelp case
	userReq = NewUserTextRequest()
	text = "使い方を教えて"
	if err := userReq.AnalyzeInputText(text); err != nil {
		t.Fatal("There shoud not be error", " Got error: ", err, "input text: ", text)
	}
	if userReq.Type != "Help" {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
	// RequestTypeGetKotos case
	userReq = NewUserTextRequest()
	text = "やること"
	if err := userReq.AnalyzeInputText(text); err != nil {
		t.Fatal("There shoud not be error", " Got error: ", err, "input text: ", text)
	}
	if userReq.Type != "GetKotos" {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
	// RequestTypeGetKotos case
	userReq = NewUserTextRequest()
	text = "やること教えて"
	if err := userReq.AnalyzeInputText(text); err != nil {
		t.Fatal("There shoud not be error", " Got error: ", err, "input text: ", text)
	}
	if userReq.Type != "GetKotos" {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
	// RequestTypeGetKotos case
	userReq = NewUserTextRequest()
	text = "やることを教えて"
	if err := userReq.AnalyzeInputText(text); err != nil {
		t.Fatal("There shoud not be error", " Got error: ", err, "input text: ", text)
	}
	if userReq.Type != "GetKotos" {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
	// RequestTypeAddKoto case
	userReq = NewUserTextRequest()
	text = "空手を登録して"
	if err := userReq.AnalyzeInputText(text); err != nil {
		t.Fatal("There shoud not be error", " Got error: ", err, "input text: ", text)
	}
	if userReq.Type != "AddKoto" {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if userReq.VariableKeyword != "空手" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
	// RequestTypeAddKoto case 2
	userReq = NewUserTextRequest()
	text = "英語の勉強を登録して"
	if err := userReq.AnalyzeInputText(text); err != nil {
		t.Fatal("There shoud not be error", " Got error: ", err, "input text: ", text)
	}
	if userReq.Type != "AddKoto" {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if userReq.VariableKeyword != "英語の勉強" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
	// RequestTypeGetActivities case
	userReq = NewUserTextRequest()
	text = "履歴"
	if err := userReq.AnalyzeInputText(text); err != nil {
		t.Fatal("There shoud not be error", " Got error: ", err, "input text: ", text)
	}
	if userReq.Type != "GetActivities" {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
	// RequestTypeGetActivities case
	userReq = NewUserTextRequest()
	text = "履歴を教えて"
	if err := userReq.AnalyzeInputText(text); err != nil {
		t.Fatal("There shoud not be error", " Got error: ", err, "input text: ", text)
	}
	if userReq.Type != "GetActivities" {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
	// RequestTypeGetActivities case 2
	userReq = NewUserTextRequest()
	text = "英語の勉強の履歴"
	if err := userReq.AnalyzeInputText(text); err != nil {
		t.Fatal("There shoud not be error", " Got error: ", err, "input text: ", text)
	}
	if userReq.Type != "GetActivities" {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if userReq.VariableKeyword != "英語の勉強" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
	// RequestTypeGetActivities case 2
	userReq = NewUserTextRequest()
	text = "英語の勉強の履歴を教えて"
	if err := userReq.AnalyzeInputText(text); err != nil {
		t.Fatal("There shoud not be error", " Got error: ", err, "input text: ", text)
	}
	if userReq.Type != "GetActivities" {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if userReq.VariableKeyword != "英語の勉強" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
	// RequestTypeAddActivity case
	userReq = NewUserTextRequest()
	text = "空手やったよ"
	if err := userReq.AnalyzeInputText(text); err != nil {
		t.Fatal("There shoud not be error", " Got error: ", err, "input text: ", text)
	}
	if userReq.Type != "AddActivity" {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if userReq.VariableKeyword != "空手" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
	// RequestTypeAddActivity case 2
	userReq = NewUserTextRequest()
	text = "英語の勉強やった"
	if err := userReq.AnalyzeInputText(text); err != nil {
		t.Fatal("There shoud not be error", " Got error: ", err, "input text: ", text)
	}
	if userReq.Type != "AddActivity" {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if userReq.VariableKeyword != "英語の勉強" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
	// RequestTypeAddActivity case 3
	userReq = NewUserTextRequest()
	text = "空手をやったよ"
	if err := userReq.AnalyzeInputText(text); err != nil {
		t.Fatal("There shoud not be error", " Got error: ", err, "input text: ", text)
	}
	if userReq.Type != "AddActivity" {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if userReq.VariableKeyword != "空手" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
	// RequestTypeAddActivity case 4
	userReq = NewUserTextRequest()
	text = "英語の勉強をやった"
	if err := userReq.AnalyzeInputText(text); err != nil {
		t.Fatal("There shoud not be error", " Got error: ", err, "input text: ", text)
	}
	if userReq.Type != "AddActivity" {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if userReq.VariableKeyword != "英語の勉強" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
}

func TestAnalyzeInputTextCaseInvalid(t *testing.T) {
	// Command case
	// Invalid text
	userReq := NewUserTextRequest()
	text := "abcdefghijklmnopqrstuvwxyz"
	if err := userReq.AnalyzeInputText(text); err == nil {
		t.Fatal("There shoud be error.", "input text:", text)
	} else if err.Error() != fmt.Sprintf("Can't analyze: %s", text) {
		t.Fatal("Error message is not correct.", "Got error:", err)
	}
	if userReq.Type != "" {
		t.Fatal("The request type should be empty, input text:", text, "user.Type:", userReq.Type)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword should be empty, input text:", text, "user.VariableKeyword:", userReq.VariableKeyword)
	}
	// No text string case
	userReq = NewUserTextRequest()
	text = ""
	if err := userReq.AnalyzeInputText(text); err == nil {
		t.Fatal("There shoud be error.", "input text:", text)
	} else if err.Error() != fmt.Sprintf("Can't analyze: %s", text) {
		t.Fatal("Error message is not correct.", "Got error:", err)
	}
	if userReq.Type != "" {
		t.Fatal("The request type should be empty, input text:", text, "user.Type:", userReq.Type)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword should be empty, input text:", text, "user.VariableKeyword:", userReq.VariableKeyword)
	}
	// No 2nd word after AddKoto
	userReq = NewUserTextRequest()
	text = "AddKoto "
	if err := userReq.AnalyzeInputText(text); err == nil {
		t.Fatal("There shoud be error.", "input text:", text)
	} else if err.Error() != fmt.Sprintf("Can't analyze: %s", text) {
		t.Fatal("Error message is not correct.", "Got error:", err)
	}
	if userReq.Type != "" {
		t.Fatal("The request type should be empty, input text:", text, "user.Type:", userReq.Type)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword should be empty, input text:", text, "user.VariableKeyword:", userReq.VariableKeyword)
	}
	// No 2nd word after AddActivity
	userReq = NewUserTextRequest()
	text = "AddActivity "
	if err := userReq.AnalyzeInputText(text); err == nil {
		t.Fatal("There shoud be error.", "input text:", text)
	} else if err.Error() != fmt.Sprintf("Can't analyze: %s", text) {
		t.Fatal("Error message is not correct.", "Got error:", err)
	}
	if userReq.Type != "" {
		t.Fatal("The request type should be empty, input text:", text, "user.Type:", userReq.Type)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword should be empty, input text:", text, "user.VariableKeyword:", userReq.VariableKeyword)
	}

	// Japanese case
	// invalid case
	userReq = NewUserTextRequest()
	text = "使い方を教えてよ"
	if err := userReq.AnalyzeInputText(text); err == nil {
		t.Fatal("There shoud be error.", "input text:", text)
	} else if err.Error() != fmt.Sprintf("Can't analyze: %s", text) {
		t.Fatal("Error message is not correct.", "Got error:", err)
	}
	if userReq.Type != "" {
		t.Fatal("The request type should be empty, input text:", text, "user.Type:", userReq.Type)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword should be empty, input text:", text, "user.VariableKeyword:", userReq.VariableKeyword)
	}
	// invalid case
	userReq = NewUserTextRequest()
	text = "やること教えてよ"
	if err := userReq.AnalyzeInputText(text); err == nil {
		t.Fatal("There shoud be error.", "input text:", text)
	} else if err.Error() != fmt.Sprintf("Can't analyze: %s", text) {
		t.Fatal("Error message is not correct.", "Got error:", err)
	}
	if userReq.Type != "" {
		t.Fatal("The request type should be empty, input text:", text, "user.Type:", userReq.Type)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword should be empty, input text:", text, "user.VariableKeyword:", userReq.VariableKeyword)
	}
	// Invalid text
	userReq = NewUserTextRequest()
	text = "空手を登録してよ"
	if err := userReq.AnalyzeInputText(text); err == nil {
		t.Fatal("There shoud be error.", "input text:", text)
	} else if err.Error() != fmt.Sprintf("Can't analyze: %s", text) {
		t.Fatal("Error message is not correct.", "Got error:", err)
	}
	if userReq.Type != "" {
		t.Fatal("The request type should be empty, input text:", text, "user.Type:", userReq.Type)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword should be empty, input text:", text, "user.VariableKeyword:", userReq.VariableKeyword)
	}
	// empty keyword text
	userReq = NewUserTextRequest()
	text = "を登録して"
	if err := userReq.AnalyzeInputText(text); err == nil {
		t.Fatal("There shoud be error.", "input text:", text)
	} else if err.Error() != fmt.Sprintf("Can't analyze: %s", text) {
		t.Fatal("Error message is not correct.", "Got error:", err)
	}
	if userReq.Type != "" {
		t.Fatal("The request type should be empty, input text:", text, "user.Type:", userReq.Type)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword should be empty, input text:", text, "user.VariableKeyword:", userReq.VariableKeyword)
	}
	// invalid
	userReq = NewUserTextRequest()
	text = "履歴を教えてよ"
	if err := userReq.AnalyzeInputText(text); err == nil {
		t.Fatal("There shoud be error.", "input text:", text)
	} else if err.Error() != fmt.Sprintf("Can't analyze: %s", text) {
		t.Fatal("Error message is not correct.", "Got error:", err)
	}
	if userReq.Type != "" {
		t.Fatal("The request type should be empty, input text:", text, "user.Type:", userReq.Type)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword should be empty, input text:", text, "user.VariableKeyword:", userReq.VariableKeyword)
	}
	// Invalid text
	userReq = NewUserTextRequest()
	text = "空手履歴を教えて"
	if err := userReq.AnalyzeInputText(text); err == nil {
		t.Fatal("There shoud be error.", "input text:", text)
	} else if err.Error() != fmt.Sprintf("Can't analyze: %s", text) {
		t.Fatal("Error message is not correct.", "Got error:", err)
	}
	if userReq.Type != "" {
		t.Fatal("The request type should be empty, input text:", text, "user.Type:", userReq.Type)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword should be empty, input text:", text, "user.VariableKeyword:", userReq.VariableKeyword)
	}
	// empty keyword
	userReq = NewUserTextRequest()
	text = "の履歴"
	if err := userReq.AnalyzeInputText(text); err == nil {
		t.Fatal("There shoud be error.", "input text:", text)
	} else if err.Error() != fmt.Sprintf("Can't analyze: %s", text) {
		t.Fatal("Error message is not correct.", "Got error:", err)
	}
	if userReq.Type != "" {
		t.Fatal("The request type should be empty, input text:", text, "user.Type:", userReq.Type)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword should be empty, input text:", text, "user.VariableKeyword:", userReq.VariableKeyword)
	}
	// empty keyword
	userReq = NewUserTextRequest()
	text = "の履歴を教えて"
	if err := userReq.AnalyzeInputText(text); err == nil {
		t.Fatal("There shoud be error.", "input text:", text)
	} else if err.Error() != fmt.Sprintf("Can't analyze: %s", text) {
		t.Fatal("Error message is not correct.", "Got error:", err)
	}
	if userReq.Type != "" {
		t.Fatal("The request type should be empty, input text:", text, "user.Type:", userReq.Type)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword should be empty, input text:", text, "user.VariableKeyword:", userReq.VariableKeyword)
	}
	// invalid keyword text
	userReq = NewUserTextRequest()
	text = "英語の勉強やってないよ"
	if err := userReq.AnalyzeInputText(text); err == nil {
		t.Fatal("There shoud be error.", "input text:", text)
	} else if err.Error() != fmt.Sprintf("Can't analyze: %s", text) {
		t.Fatal("Error message is not correct.", "Got error:", err)
	}
	if userReq.Type != "" {
		t.Fatal("The request type should be empty, input text:", text, "user.Type:", userReq.Type)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword should be empty, input text:", text, "user.VariableKeyword:", userReq.VariableKeyword)
	}
	// invalid keyword text
	userReq = NewUserTextRequest()
	text = "英語の勉強やったよー"
	if err := userReq.AnalyzeInputText(text); err == nil {
		t.Fatal("There shoud be error.", "input text:", text)
	} else if err.Error() != fmt.Sprintf("Can't analyze: %s", text) {
		t.Fatal("Error message is not correct.", "Got error:", err)
	}
	if userReq.Type != "" {
		t.Fatal("The request type should be empty, input text:", text, "user.Type:", userReq.Type)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword should be empty, input text:", text, "user.VariableKeyword:", userReq.VariableKeyword)
	}
	// invalid keyword text
	userReq = NewUserTextRequest()
	text = "英語の勉強をやったよん"
	if err := userReq.AnalyzeInputText(text); err == nil {
		t.Fatal("There shoud be error.", "input text:", text)
	} else if err.Error() != fmt.Sprintf("Can't analyze: %s", text) {
		t.Fatal("Error message is not correct.", "Got error:", err)
	}
	if userReq.Type != "" {
		t.Fatal("The request type should be empty, input text:", text, "user.Type:", userReq.Type)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword should be empty, input text:", text, "user.VariableKeyword:", userReq.VariableKeyword)
	}
	// empty case
	userReq = NewUserTextRequest()
	text = "をやったよ"
	if err := userReq.AnalyzeInputText(text); err == nil {
		t.Fatal("There shoud be error.", "input text:", text)
	} else if err.Error() != fmt.Sprintf("Can't analyze: %s", text) {
		t.Fatal("Error message is not correct.", "Got error:", err)
	}
	if userReq.Type != "" {
		t.Fatal("The request type should be empty, input text:", text, "user.Type:", userReq.Type)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword should be empty, input text:", text, "user.VariableKeyword:", userReq.VariableKeyword)
	}
	// empty case
	userReq = NewUserTextRequest()
	text = "やったよ"
	if err := userReq.AnalyzeInputText(text); err == nil {
		t.Fatal("There shoud be error.", "input text:", text)
	} else if err.Error() != fmt.Sprintf("Can't analyze: %s", text) {
		t.Fatal("Error message is not correct.", "Got error:", err)
	}
	if userReq.Type != "" {
		t.Fatal("The request type should be empty, input text:", text, "user.Type:", userReq.Type)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword should be empty, input text:", text, "user.VariableKeyword:", userReq.VariableKeyword)
	}
	// empty case
	userReq = NewUserTextRequest()
	text = "をやった"
	if err := userReq.AnalyzeInputText(text); err == nil {
		t.Fatal("There shoud be error.", "input text:", text)
	} else if err.Error() != fmt.Sprintf("Can't analyze: %s", text) {
		t.Fatal("Error message is not correct.", "Got error:", err)
	}
	if userReq.Type != "" {
		t.Fatal("The request type should be empty, input text:", text, "user.Type:", userReq.Type)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword should be empty, input text:", text, "user.VariableKeyword:", userReq.VariableKeyword)
	}
	// empty case
	userReq = NewUserTextRequest()
	text = "やった"
	if err := userReq.AnalyzeInputText(text); err == nil {
		t.Fatal("There shoud be error.", "input text:", text)
	} else if err.Error() != fmt.Sprintf("Can't analyze: %s", text) {
		t.Fatal("Error message is not correct.", "Got error:", err)
	}
	if userReq.Type != "" {
		t.Fatal("The request type should be empty, input text:", text, "user.Type:", userReq.Type)
	}
	if userReq.VariableKeyword != "" {
		t.Fatal("The keyword should be empty, input text:", text, "user.VariableKeyword:", userReq.VariableKeyword)
	}
}
*/
