package main

import "time"

// YaranaDataCallMock is a test struct of DataCall for Yarana-bot
type YaranaDataCallMock struct {
}

// NewYaranaDataCallMock is a constructor of YaranaDataCallMock
func NewYaranaDataCallMock() (*YaranaDataCallMock, error) {
	return &YaranaDataCallMock{}, nil
}

// GetKotosByUserID is a method of DataCall interface
func (c *YaranaDataCallMock) GetKotosByUserID(userID string) ([]*KotoData, error) {
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
func (c *YaranaDataCallMock) AddKoto(koto *KotoData) error {
	return nil
}

// GetActivitiesByKotoDataID is a method of DataCall interface
func (c *YaranaDataCallMock) GetActivitiesByKotoDataID(kotoID string) ([]*ActivityData, error) {
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
func (c *YaranaDataCallMock) AddActivity(activity *ActivityData) error {
	return nil
}
