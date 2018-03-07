package main

import (
	"encoding/json"
	"time"
)

// KotoData is DTO of thing to do in Yarana-Bot
type KotoData struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
	Title  string `json:"title"`
}

// NewKotoData is constructor of KotoData
func NewKotoData(id string, userID string, title string) (*KotoData, error) {
	return &KotoData{
		ID:     id,
		UserID: userID,
		Title:  title,
	}, nil
}

// ActivityData is DTO of user activity for KotoData
type ActivityData struct {
	ID         string
	KotoDataID string
	TimeStamp  time.Time
}

// NewActivityData is constructor of ActivityData
func NewActivityData(id string, kotoID string, timeStamp time.Time) (*ActivityData, error) {
	return &ActivityData{
		ID:         id,
		KotoDataID: kotoID,
		TimeStamp:  timeStamp,
	}, nil
}

/*
// DataCall to be
type DataCall interface {
	GetKotoByID(id string) (*KotoData, error)
	GetKotosByUserID(userID string) ([]*KotoData, error)
	AddKoto(koto *KotoData) error
	EditKoto(id string, koto *KotoData) (*KotoData, error)
	DeleteKoto(id string) error
	GetActivityByID(id string) (*ActivityData, error)
	GetActivitiesByKotoDataID(kotoID string) ([]*ActivityData, error)
	AddActivity(activity *ActivityData) error
}
*/

// DataCall is a alternative of DataCall // TODO: interface for prototype
type DataCall interface {
	GetKotosByUserID(userID string) ([]*KotoData, error)
	AddKoto(koto *KotoData) error
	GetActivitiesByKotoDataID(kotoID string) ([]*ActivityData, error)
	AddActivity(activity *ActivityData) error
}

// YaranaDataCall is a struct of DataCall for Yarana-bot
type YaranaDataCall struct {
}

// NewYaranaDataCall is a constructor of YaranaDataCall
func NewYaranaDataCall() (*YaranaDataCall, error) {
	return &YaranaDataCall{}, nil
}

// GetKotosByUserID is a method of DataCall interface
func (c *YaranaDataCall) GetKotosByUserID(userID string) (kotos []*KotoData, err error) {
	baseURL := "https://yarana-api.azurewebsites.net/api/" + "kotos"
	url := AssembleURLWithParam(baseURL, "userId", userID) // get url like https://yarana-api.azurewebsites.net/api/kotos?userId=d59964bb713fd6f4f5ef6a7c7e029387
	body, err := HTTPGet(url)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &kotos); err != nil {
		return nil, err
	}
	return kotos, nil
}

// AddKoto is a method of DataCall interface
func (c *YaranaDataCall) AddKoto(koto *KotoData) error {
	return nil
}

// GetActivitiesByKotoDataID is a method of DataCall interface
func (c *YaranaDataCall) GetActivitiesByKotoDataID(kotoID string) ([]*ActivityData, error) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	timeStamp := time.Date(2018, 3, 3, 3, 3, 35, 0, loc)
	activity, err := NewActivityData("0123456789a", kotoID, timeStamp)
	if err != nil {
		return nil, err
	}
	loc2, _ := time.LoadLocation("Asia/Tokyo")
	timeStamp2 := time.Date(2018, 3, 3, 3, 3, 36, 0, loc2)
	activity2, err := NewActivityData("0123456789b", kotoID, timeStamp2)
	if err != nil {
		return nil, err
	}
	return []*ActivityData{activity, activity2}, nil
}

// AddActivity is a method of DataCall interface
func (c *YaranaDataCall) AddActivity(activity *ActivityData) error {
	return nil
}
