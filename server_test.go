package main

import (
	"testing"
	"time"
)

const linebotAPIError401Str = "linebot: APIError 401 Authentication failed due to the following reason: invalid token. Confirm that the access token in the authorization header is valid."

// StandardYaranaDataCallMock is a test struct of DataCall for Yarana-bot
type StandardYaranaDataCallMock struct {
}

// NewStandardYaranaDataCallMock is a constructor of StandardYaranaDataCallMock
func NewStandardYaranaDataCallMock() (*StandardYaranaDataCallMock, error) {
	return &StandardYaranaDataCallMock{}, nil
}

// GetKotosByUserID is a method of DataCall interface
func (c *StandardYaranaDataCallMock) GetKotosByUserID(userID string) ([]*KotoData, error) {
	// Get Koto by userID from something
	id := "0123456789a"
	title := "TestTitle"
	koto, err := NewKotoData(id, userID, title)
	if err != nil {
		return nil, err
	}
	title2 := "TestTitle2"
	koto2, err := NewKotoData(id, userID, title2)
	if err != nil {
		return nil, err
	}
	return []*KotoData{koto, koto2}, nil
}

// AddKoto is a method of DataCall interface
func (c *StandardYaranaDataCallMock) AddKoto(koto *KotoData) error {
	return nil
}

// GetActivitiesByKotoDataID is a method of DataCall interface
func (c *StandardYaranaDataCallMock) GetActivitiesByKotoDataID(kotoID string) ([]*ActivityData, error) {
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
func (c *StandardYaranaDataCallMock) AddActivity(activity *ActivityData) error {
	return nil
}

func TestProcessGetKotosStandard(t *testing.T) {
	dataCall, _ := NewStandardYaranaDataCallMock()
	app, _ := NewYarana(
		"TEST_YARANA_API_BASE_URL",
		"TEST_YARANA_API_ADDKOTO_KEY",
		"TEST_YARANA_API_ADDACTIVITY_KEY",
		dataCall,
	)

	err := app.processGetKotos("replyToken", "userID", "TestTitle")
	if err != nil {
		if err.Error() != linebotAPIError401Str {
			t.Fatal("err should be linebot APIError 401", " Got error: ", err)
		}
	} else {
		t.Fatal("err should be linebot APIError 401")
	}
}

func TestProcessAddKotosStandard(t *testing.T) {
	dataCall, _ := NewStandardYaranaDataCallMock()
	app, _ := NewYarana(
		"TEST_YARANA_API_BASE_URL",
		"TEST_YARANA_API_ADDKOTO_KEY",
		"TEST_YARANA_API_ADDACTIVITY_KEY",
		dataCall,
	)

	err := app.processAddKoto("replyToken", "userID", "TestTitleAddKoto")
	if err != nil {
		if err.Error() != linebotAPIError401Str {
			t.Fatal("err should be linebot APIError 401", " Got error: ", err)
		}
	} else {
		t.Fatal("err should be linebot APIError 401")
	}
}

func TestProcessGetActivitiesStandard(t *testing.T) {
	dataCall, _ := NewStandardYaranaDataCallMock()
	app, _ := NewYarana(
		"TEST_YARANA_API_BASE_URL",
		"TEST_YARANA_API_ADDKOTO_KEY",
		"TEST_YARANA_API_ADDACTIVITY_KEY",
		dataCall,
	)

	err := app.processGetActivities("replyToken", "userID", "")
	if err != nil {
		if err.Error() != linebotAPIError401Str {
			t.Fatal("err should be linebot APIError 401", " Got error: ", err)
		}
	} else {
		t.Fatal("err should be linebot APIError 401")
	}
}

func TestProcessAddActivityStandard(t *testing.T) {
	dataCall, _ := NewStandardYaranaDataCallMock()
	app, _ := NewYarana(
		"TEST_YARANA_API_BASE_URL",
		"TEST_YARANA_API_ADDKOTO_KEY",
		"TEST_YARANA_API_ADDACTIVITY_KEY",
		dataCall,
	)

	err := app.processAddActivity("replyToken", "userID", "TestTitle")
	if err != nil {
		if err.Error() != linebotAPIError401Str {
			t.Fatal("err should be linebot APIError 401", " Got error: ", err)
		}
	} else {
		t.Fatal("err should be linebot APIError 401")
	}
}
