package main

import (
	"time"
)

// KotoData is DTO of thing to do in Yarana-Bot
type KotoData struct {
	ID     string
	UserID string
	Title  string
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

// TODO: Implement methods of YaranaDataCall

// NewYaranaDataCall is a constructor of YaranaDataCall
func NewYaranaDataCall() (*YaranaDataCall, error) {
	return &YaranaDataCall{}, nil
}

// GetKotosByUserID is a method of DataCall interface
func (c *YaranaDataCall) GetKotosByUserID(userID string) ([]*KotoData, error) {
	// Get Koto by userID from something
	id := "0123456789a"
	title := "Test Title"
	koto, err := NewKotoData(id, userID, title)
	if err != nil {
		return nil, err
	}
	title2 := "Test Title 2"
	koto2, err := NewKotoData(id, userID, title2)
	if err != nil {
		return nil, err
	}
	return []*KotoData{koto, koto2}, nil
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
