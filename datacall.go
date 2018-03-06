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

// DataCall is a main interface of Yarukoto
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

// SimpleDataCall is a alternative of DataCall // TODO: interface for prototype
type SimpleDataCall interface {
	GetKotosByUserID(userID string) ([]*KotoData, error)
	AddKoto(koto *KotoData) error
	GetActivitiesByKotoDataID(kotoID string) ([]*ActivityData, error)
	AddActivity(activity *ActivityData) error
}

// YaranaDataCall is a struct of DataCall for Yarana-bot
type YaranaDataCall struct {
}

// TODO: Implement methods of YaranaDataCall

// YaranaDataCallForTest is a test struct of DataCall for Yarana-bot
type YaranaDataCallForTest struct {
}

// NewYaranaDataCallForTest is a constructor of YaranaDataCallForTest
func NewYaranaDataCallForTest() (*YaranaDataCallForTest, error) {
	return &YaranaDataCallForTest{}, nil
}

// GetKotosByUserID is a method of DataCall interface
func (c *YaranaDataCallForTest) GetKotosByUserID(userID string) ([]*KotoData, error) {
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
func (c *YaranaDataCallForTest) AddKoto(koto *KotoData) error {
	return nil
}

// GetActivitiesByKotoDataID is a method of DataCall interface
func (c *YaranaDataCallForTest) GetActivitiesByKotoDataID(kotoID string) ([]*ActivityData, error) {
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
func (c *YaranaDataCallForTest) AddActivity(activity *ActivityData) error {
	return nil
}
