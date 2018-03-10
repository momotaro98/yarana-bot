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

	switch fWord := words[0]; fWord {
	case "GetKotos":
		r.Type = RequstTypeGetKotos
	case "AddKoto":
		if len(words) < 2 {
			return fmt.Errorf("Can't analyze: %s", text)
		}
		r.Type = RequstTypeAddKoto
		r.VariableKeyword = words[1]
	default:
		return fmt.Errorf("Can't analyze: %s", text)
	}

	return nil
}
