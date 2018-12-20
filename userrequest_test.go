package main

import (
	"testing"
)

// assertForErrorExpectedCases is common assertion method for cases which are expected to get error
func assertForErrorExpectedCases(t *testing.T, err error, reqType RequestType, keyword string, inputText string) {
	if err == nil {
		t.Fatal("There should be error.", "input text:", inputText)
	}
	if reqType != "" {
		t.Fatal("The request type should be empty, input text:", inputText, "reqType:", reqType)
	}
	if keyword != "" {
		t.Fatal("The keyword should be empty, input text:", inputText, "keyword:", keyword)
	}
}

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
	assertForErrorExpectedCases(t, err, reqType, keyword, text)
}

func TestAnalyzeInputText_EmptyStringText_Error(t *testing.T) {
	// Arrange
	text := ""
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	assertForErrorExpectedCases(t, err, reqType, keyword, text)
}

func TestAnalyzeInputText_None2ndWordAfterAddKoto_Error(t *testing.T) {
	// Arrange
	text := "AddKoto "
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	assertForErrorExpectedCases(t, err, reqType, keyword, text)
}

func TestAnalyzeInputText_None2ndWordAferAddActivity_Error(t *testing.T) {
	// Arrange
	text := "AddActivity "
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	assertForErrorExpectedCases(t, err, reqType, keyword, text)
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

func TestAnalyzeInputText_JapaneseTextWrongHelp1_Error(t *testing.T) {
	// Arrange
	text := "使い方を教えてよ"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	assertForErrorExpectedCases(t, err, reqType, keyword, text)
}

func TestAnalyzeInputText_JapaneseTextWrongGetKotos1_Error(t *testing.T) {
	// Arrange
	text := "やること教えてよ"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	assertForErrorExpectedCases(t, err, reqType, keyword, text)
}

func TestAnalyzeInputText_JapaneseTextWrongAddKoto1_Error(t *testing.T) {
	// Arrange
	text := "空手を登録してよ"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	assertForErrorExpectedCases(t, err, reqType, keyword, text)
}

func TestAnalyzeInputText_JapaneseTextAddKotoWithoutKeyword_Error(t *testing.T) {
	// Arrange
	text := "を登録して"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	assertForErrorExpectedCases(t, err, reqType, keyword, text)
}

func TestAnalyzeInputText_JapaneseTextWrongGetActivities1_Error(t *testing.T) {
	// Arrange
	text := "履歴を教えてよ"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	assertForErrorExpectedCases(t, err, reqType, keyword, text)
}

func TestAnalyzeInputText_JapaneseTextWrongGetActivities2_Error(t *testing.T) {
	// Arrange
	text := "空手履歴を教えて"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	assertForErrorExpectedCases(t, err, reqType, keyword, text)
}

func TestAnalyzeInputText_JapaneseTextGetActivitiesWithoutKeyword1_Error(t *testing.T) {
	// Arrange
	text := "の履歴"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	assertForErrorExpectedCases(t, err, reqType, keyword, text)
}

func TestAnalyzeInputText_JapaneseTextGetActivitiesWithoutKeyword2_Error(t *testing.T) {
	// Arrange
	text := "の履歴を教えて"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	assertForErrorExpectedCases(t, err, reqType, keyword, text)
}

func TestAnalyzeInputText_JapaneseTextWrongAddActivity1_Error(t *testing.T) {
	// Arrange
	text := "英語の勉強やってないよ"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	assertForErrorExpectedCases(t, err, reqType, keyword, text)
}

func TestAnalyzeInputText_JapaneseTextAddActivityWithoutKeyword1_Error(t *testing.T) {
	// Arrange
	text := "をやったよ"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	assertForErrorExpectedCases(t, err, reqType, keyword, text)
}

func TestAnalyzeInputText_JapaneseTextAddActivityWithoutKeyword2_Error(t *testing.T) {
	// Arrange
	text := "やったよ"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	assertForErrorExpectedCases(t, err, reqType, keyword, text)
}

func TestAnalyzeInputText_JapaneseTextAddActivityWithoutKeyword3_Error(t *testing.T) {
	// Arrange
	text := "をやった"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	assertForErrorExpectedCases(t, err, reqType, keyword, text)
}

func TestAnalyzeInputText_JapaneseTextAddActivityWithoutKeyword4_Error(t *testing.T) {
	// Arrange
	text := "やった"
	// Act
	reqType, keyword, err := AnalyzeInputText(text)
	// Assert
	assertForErrorExpectedCases(t, err, reqType, keyword, text)
}
