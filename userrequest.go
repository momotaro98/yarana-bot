package main

import (
	"fmt"
	"strings"
)

// RequstType type
type RequstType string

// RequstType constants
const (
	RequstTypeGetKotos      RequstType = "GetKotos"
	RequstTypeAddKoto       RequstType = "AddKoto"
	RequstTypeGetActivities RequstType = "GetActivities"
	RequstTypeAddActivity   RequstType = "AddActivity"
)

// UserTextRequest is struct for managing input text from user
type UserTextRequest struct {
	Type            RequstType
	VariableKeyword string
}

// NewUserTextRequest is constructor of KotoData
func NewUserTextRequest() *UserTextRequest {
	return &UserTextRequest{}
}

// AnalyzeInputText analyzes input text from user
func (r *UserTextRequest) AnalyzeInputText(text string) error {
	// TODO: Implement analyzer more rich
	words := strings.Fields(text)
	if len(words) < 1 {
		return fmt.Errorf("Can't analyze: %s", text)
	}

	if err := r.AnalyzeInputTextInCommand(text); r.Type != "" && err == nil {
		return nil
	}
	if err := r.AnalyzeInputTextInJapanese(text); r.Type != "" && err == nil {
		return nil
	}

	return fmt.Errorf("Can't analyze: %s", text)
}

// AnalyzeInputTextInCommand analyzes input text from user with Command
func (r *UserTextRequest) AnalyzeInputTextInCommand(text string) error {
	words := strings.Fields(text)
	if len(words) < 1 {
		return fmt.Errorf("Can't analyze: %s", text)
	}
	switch fWord := words[0]; fWord {
	case "GetKotos":
		r.Type = RequstTypeGetKotos
	case "AddKoto":
		if len(words) < 2 {
			return fmt.Errorf("Can't analyze: %s", text)
		}
		r.Type = RequstTypeAddKoto
		r.VariableKeyword = words[1]
	case "GetActivities":
		r.Type = RequstTypeGetActivities
		if len(words) >= 2 {
			r.VariableKeyword = words[1]
		}
	case "AddActivity":
		if len(words) < 2 {
			return fmt.Errorf("Can't analyze: %s", text)
		}
		r.Type = RequstTypeAddActivity
		r.VariableKeyword = words[1]
	default:
		return fmt.Errorf("Can't analyze: %s", text)
	}
	return nil
}

// AnalyzeInputTextInJapanese analyzes input text from user with Japanese
func (r *UserTextRequest) AnalyzeInputTextInJapanese(text string) error {
	words := strings.Fields(text)
	if len(words) < 1 {
		return fmt.Errorf("Can't analyze: %s", text)
	}
	// Japanese mode doesn't accept multi "words"
	if len(words) > 1 {
		return fmt.Errorf("Can't analyze: %s", text)
	}

	if text == "やること教えて" || text == "やることを教えて" {
		r.Type = RequstTypeGetKotos
		return nil
	}

	if n := len(text) - 3*5; n >= 0 {
		if text[n:] == "を登録して" {
			if len(text[:n]) <= 0 {
				return fmt.Errorf("Can't analyze: %s", text)
			}
			r.Type = RequstTypeAddKoto
			r.VariableKeyword = text[:n]
			return nil
		}
	}

	if text == "履歴教えて" || text == "履歴を教えて" {
		r.Type = RequstTypeGetActivities
		return nil
	} else if n := len(text) - 3*7; n >= 0 {
		if text[n:] == "の履歴を教えて" {
			if len(text[:n]) <= 0 {
				return fmt.Errorf("Can't analyze: %s", text)
			}
			r.Type = RequstTypeGetActivities
			r.VariableKeyword = text[:n]
			return nil
		}
	}

	if n := len(text) - 3*5; n >= 0 {
		if text[n:] == "をやったよ" {
			if len(text[:n]) <= 0 {
				return fmt.Errorf("Can't analyze: %s", text)
			}
			r.Type = RequstTypeAddActivity
			r.VariableKeyword = text[:n]
			return nil
		}
	}
	if n := len(text) - 3*4; n >= 0 {
		if text[n:] == "をやった" || text[n:] == "やったよ" {
			if len(text[:n]) <= 0 {
				return fmt.Errorf("Can't analyze: %s", text)
			}
			r.Type = RequstTypeAddActivity
			r.VariableKeyword = text[:n]
			return nil
		}
	}
	if n := len(text) - 3*3; n >= 0 {
		if text[n:] == "やった" {
			if len(text[:n]) <= 0 {
				return fmt.Errorf("Can't analyze: %s", text)
			}
			r.Type = RequstTypeAddActivity
			r.VariableKeyword = text[:n]
			return nil
		}
	}

	return nil
}
