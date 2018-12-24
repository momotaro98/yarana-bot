package main

import (
	"fmt"
	"strings"
)

// RequestType type
type RequestType string

// RequestType constants
const (
	RequestTypeHelp          RequestType = "Help"
	RequestTypeGetKotos      RequestType = "GetKotos"
	RequestTypeAddKoto       RequestType = "AddKoto"
	RequestTypeGetActivities RequestType = "GetActivities"
	RequestTypeAddActivity   RequestType = "AddActivity"
)

// AnalyzeInputText analyzes input text from user
// and returns RequestType and keyword of the type, and error
func AnalyzeInputText(text string) (RequestType, string, error) {
	words := strings.Fields(text)
	if len(words) < 1 {
		return "", "", fmt.Errorf("can't analyze: %s", text)
	}
	if reqType, keyword, err := analyzeInputTextInCommand(text); reqType != "" && err == nil {
		return reqType, keyword, err
	}
	if reqType, keyword, err := analyzeInputTextInJapanese(text); reqType != "" && err == nil {
		return reqType, keyword, err
	}

	return "", "", fmt.Errorf("can't analyze: %s", text)
}

// AnalyzeInputTextInCommand analyzes input text from user with Command
func analyzeInputTextInCommand(text string) (RequestType, string, error) {
	words := strings.Fields(text)
	if len(words) < 1 {
		return "", "", fmt.Errorf("can't analyze: %s", text)
	}
	switch fWord := words[0]; fWord {
	case "Help":
		if len(words) != 1 {
			return "", "", fmt.Errorf("can't analyze: %s", text)
		}
		return RequestTypeHelp, "", nil
	case "GetKotos":
		if len(words) != 1 {
			return "", "", fmt.Errorf("can't analyze: %s", text)
		}
		return RequestTypeGetKotos, "", nil
	case "AddKoto":
		if len(words) < 2 {
			return "", "", fmt.Errorf("can't analyze: %s", text)
		}
		return RequestTypeAddKoto, words[1], nil
	case "GetActivities":
		if len(words) >= 2 {
			return RequestTypeGetActivities, words[1], nil
		}
		return RequestTypeGetActivities, "", nil
	case "AddActivity":
		if len(words) < 2 {
			return "", "", fmt.Errorf("can't analyze: %s", text)
		}
		return RequestTypeAddActivity, words[1], nil
	}
	
	return "", "", fmt.Errorf("can't analyze: %s", text)
}

// analyzeInputTextInJapanese analyzes input text from user with Japanese
func analyzeInputTextInJapanese(text string) (RequestType, string, error) {
	words := strings.Fields(text)
	if len(words) < 1 {
		return "", "", fmt.Errorf("can't analyze: %s", text)
	}
	// Japanese mode doesn't accept multi "words"
	if len(words) > 1 {
		return "", "", fmt.Errorf("can't analyze: %s", text)
	}

	// Help
	if text == "使い方" || text == "使い方教えて" || text == "使い方を教えて" ||
		text == "使いかた" || text == "使いかた教えて" || text == "使いかたを教えて" ||
		text == "つかい方" || text == "つかい方教えて" || text == "つかい方を教えて" ||
		text == "つかいかた" || text == "つかいかた教えて" || text == "つかいかたを教えて" {
		return RequestTypeHelp, "", nil
	}

	// GetKotos
	if text == "やること" || text == "やること教えて" || text == "やることを教えて" {
		return RequestTypeGetKotos, "", nil
	}

	// AddKoto
	if n := len(text) - 3*5; n >= 0 {
		if text[n:] == "を登録して" {
			if len(text[:n]) <= 0 {
				return "", "", fmt.Errorf("can't analyze: %s", text)
			}
			return RequestTypeAddKoto, text[:n], nil
		}
	}

	// GetActivities
	if text == "履歴" || text == "履歴教えて" || text == "履歴を教えて" {
		return RequestTypeGetActivities, "", nil
	}
	if n := len(text) - 3*7; n >= 0 {
		if text[n:] == "の履歴を教えて" {
			if len(text[:n]) <= 0 {
				return "", "", fmt.Errorf("can't analyze: %s", text)
			}
			return RequestTypeGetActivities, text[:n], nil
		}
	}
	if n := len(text) - 3*3; n >= 0 {
		if text[n:] == "の履歴" {
			if len(text[:n]) <= 0 {
				return "", "", fmt.Errorf("can't analyze: %s", text)
			}
			return RequestTypeGetActivities, text[:n], nil
		}
	}

	// AddActivity
	if n := len(text) - 3*5; n >= 0 {
		if text[n:] == "をやったよ" {
			if len(text[:n]) <= 0 {
				return "", "", fmt.Errorf("can't analyze: %s", text)
			}
			return RequestTypeAddActivity, text[:n], nil
		}
	}
	if n := len(text) - 3*4; n >= 0 {
		if text[n:] == "をやった" || text[n:] == "やったよ" {
			if len(text[:n]) <= 0 {
				return "", "", fmt.Errorf("can't analyze: %s", text)
			}
			return RequestTypeAddActivity, text[:n], nil
		}
	}
	if n := len(text) - 3*3; n >= 0 {
		if text[n:] == "やった" {
			if len(text[:n]) <= 0 {
				return "", "", fmt.Errorf("can't analyze: %s", text)
			}
			return RequestTypeAddActivity, text[:n], nil
		}
	}

	return "", "", fmt.Errorf("can't analyze: %s", text)
}
