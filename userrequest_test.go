package main

import (
		"testing"
	)

func TestAnalyzeInputText_SingleHelpCommand_Valid(t *testing.T) {
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

func TestAnalyzeInputText_SingleGetKotosCommand_Valid(t *testing.T) {
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

func TestAnalyzeInputText_AddKotoCommandWithKeyword_Valid(t *testing.T) {
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

func TestAnalyzeInputText_SingleGetActivitiesCommand_Valid(t *testing.T) {
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

func TestAnalyzeInputText_GetActivitiesCommandWithKeyword_Valid(t *testing.T) {
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

func TestAnalyzeInputText_AddActivityCommandWithKeyword_Valid(t *testing.T) {
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

func TestAnalyzeInputText_InvalidText_Error(t *testing.T) {
	// Arrange
	text := "abcdefghijklmnopqrstuvwxyz"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	if err == nil {
		t.Fatal("There should be error.", "input text:", text)
	}
	if reqType != "" {
		t.Fatal("The request type should be empty, input text:", text, "RequestType:", reqType)
	}
	if keyword != "" {
		t.Fatal("The keyword should be empty, input text:", text, "keyword:", keyword)
	}
}

func TestAnalyzeInputText_EmptyStringText_Error(t *testing.T) {
	// Arrange
	text := ""
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	if err == nil {
		t.Fatal("There should be error.", "input text:", text)
	}
	if reqType != "" {
		t.Fatal("The request type should be empty, input text:", text, "RequestType:", reqType)
	}
	if keyword != "" {
		t.Fatal("The keyword should be empty, input text:", text, "keyword:", keyword)
	}
}

func TestAnalyzeInputText_None2ndWordAfterAddKoto_Error(t *testing.T) {
	// Arrange
	text := "AddKoto "
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	if err == nil {
		t.Fatal("There should be error.", "input text:", text)
	}
	if reqType != "" {
		t.Fatal("The request type should be empty, input text:", text, "RequestType:", reqType)
	}
	if keyword != "" {
		t.Fatal("The keyword should be empty, input text:", text, "keyword:", keyword)
	}
}

func TestAnalyzeInputText_None2ndWordAferAddActivity_Error(t *testing.T) {
	// Arrange
	text := "AddActivity "
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	if err == nil {
		t.Fatal("There should be error.", "input text:", text)
	}
	if reqType != "" {
		t.Fatal("The request type should be empty, input text:", text, "RequestType:", reqType)
	}
	if keyword != "" {
		t.Fatal("The keyword should be empty, input text:", text, "keyword:", keyword)
	}
}

func TestAnalyzeInputText_JapaneseTextHelp_Valid(t *testing.T) {
	// Arrange
	text := "使い方"
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

func TestAnalyzeInputText_JapaneseTextHelp2_Valid(t *testing.T) {
	// Arrange
	text := "使い方教えて"
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

func TestAnalyzeInputText_JapaneseTextHelp3_Valid(t *testing.T) {
	// Arrange
	text := "使い方を教えて"
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

func TestAnalyzeInputText_JapaneseTextGetKotos_Valid(t *testing.T) {
	// Arrange
	text := "やること"
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

func TestAnalyzeInputText_JapaneseTextGetKotos2_Valid(t *testing.T) {
	// Arrange
	text := "やること教えて"
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

func TestAnalyzeInputText_JapaneseTextGetKotos3_Valid(t *testing.T) {
	// Arrange
	text := "やることを教えて"
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

func TestAnalyzeInputText_JapaneseTextAddKotoWithKeyword1_Valid(t *testing.T) {
	// Arrange
	text := "空手を登録して"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	if err != nil {
		t.Fatal("There should not be error", " Got error: ", err, "input text: ", text)
	}
	if reqType != RequestTypeAddKoto {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if keyword != "空手" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
}

func TestAnalyzeInputText_JapaneseTextAddKotoWithKeyword2_Valid(t *testing.T) {
	// Arrange
	text := "英語の勉強を登録して"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	if err != nil {
		t.Fatal("There should not be error", " Got error: ", err, "input text: ", text)
	}
	if reqType != RequestTypeAddKoto {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if keyword != "英語の勉強" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
}

func TestAnalyzeInputText_JapaneseTextGetActivities1_Valid(t *testing.T) {
	// Arrange
	text := "履歴"
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

func TestAnalyzeInputText_JapaneseTextGetActivities2_Valid(t *testing.T) {
	// Arrange
	text := "履歴を教えて"
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

func TestAnalyzeInputText_JapaneseTextGetActivitiesWithKeyword_Valid(t *testing.T) {
	// Arrange
	text := "英語の勉強の履歴"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	if err != nil {
		t.Fatal("There should not be error", " Got error: ", err, "input text: ", text)
	}
	if reqType != RequestTypeGetActivities {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if keyword != "英語の勉強" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
}

func TestAnalyzeInputText_JapaneseTextGetActivitiesWithKeyword2_Valid(t *testing.T) {
	// Arrange
	text := "英語の勉強の履歴を教えて"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	if err != nil {
		t.Fatal("There should not be error", " Got error: ", err, "input text: ", text)
	}
	if reqType != RequestTypeGetActivities {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if keyword != "英語の勉強" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
}

func TestAnalyzeInputText_JapaneseTextAddActivityWithKeyword_Valid(t *testing.T) {
	// Arrange
	text := "空手やったよ"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	if err != nil {
		t.Fatal("There should not be error", " Got error: ", err, "input text: ", text)
	}
	if reqType != RequestTypeAddActivity {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if keyword != "空手" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
}

func TestAnalyzeInputText_JapaneseTextAddActivityWithKeyword2_Valid(t *testing.T) {
	// Arrange
	text := "英語の勉強やった"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	if err != nil {
		t.Fatal("There should not be error", " Got error: ", err, "input text: ", text)
	}
	if reqType != RequestTypeAddActivity {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if keyword != "英語の勉強" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
}

func TestAnalyzeInputText_JapaneseTextAddActivityWithKeyword3_Valid(t *testing.T) {
	// Arrange
	text := "空手をやったよ"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	if err != nil {
		t.Fatal("There should not be error", " Got error: ", err, "input text: ", text)
	}
	if reqType != RequestTypeAddActivity {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if keyword != "空手" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
}

func TestAnalyzeInputText_JapaneseTextAddActivityWithKeyword4_Valid(t *testing.T) {
	// Arrange
	text := "英語の勉強をやった"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	if err != nil {
		t.Fatal("There should not be error", " Got error: ", err, "input text: ", text)
	}
	if reqType != RequestTypeAddActivity {
		t.Fatal("The request type is not correct, input text: ", text)
	}
	if keyword != "英語の勉強" {
		t.Fatal("The keyword is not correct, input text: ", text)
	}
}

/*
func TestAnalyzeInputTextCaseInvalid(t *testing.T) {

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
