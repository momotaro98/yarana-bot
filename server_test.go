package main

import (
	"fmt"
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

// GetUsers is a method of DataCall interface
func (c *StandardYaranaDataCallMock) GetUsers() ([]*User, error) {
	user1, err := NewUser("0123456789a")
	if err != nil {
		return nil, err
	}
	user2, err := NewUser("0123456789b")
	if err != nil {
		return nil, err
	}
	return []*User{user1, user2}, nil
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
		"TEST_CHANNEL_SECRET",
		"TEST_CHANNEL_TOKEN",
		"TEST_ENDPOINT_BASE",
		dataCall,
	)

	// Standard case
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
		"TEST_CHANNEL_SECRET",
		"TEST_CHANNEL_TOKEN",
		"TEST_ENDPOINT_BASE",
		dataCall,
	)

	// Standard case
	err := app.processAddKoto("replyToken", "userID", "NewTestTitle")
	if err != nil {
		if err.Error() != linebotAPIError401Str {
			t.Fatal("err should be linebot APIError 401", " Got error: ", err)
		}
	} else {
		t.Fatal("err should be linebot APIError 401")
	}

	// duplicate Koto title case
	err = app.processAddKoto("replyToken", "userID", "TestTitle")
	if err != nil {
		if err.Error() != "User was going to add duplicate Koto." {
			t.Fatal("err should be yarana-bot duplicate Koto error.", " Got error: ", err)
		}
	} else {
		t.Fatal("err should have error")
	}
}

func TestProcessGetActivitiesStandard(t *testing.T) {
	dataCall, _ := NewStandardYaranaDataCallMock()
	app, _ := NewYarana(
		"TEST_CHANNEL_SECRET",
		"TEST_CHANNEL_TOKEN",
		"TEST_ENDPOINT_BASE",
		dataCall,
	)

	// Standard case
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
		"TEST_CHANNEL_SECRET",
		"TEST_CHANNEL_TOKEN",
		"TEST_ENDPOINT_BASE",
		dataCall,
	)

	// Standard case
	err := app.processAddActivity("replyToken", "userID", "TestTitle")
	if err != nil {
		if err.Error() != linebotAPIError401Str {
			t.Fatal("err should be linebot APIError 401", " Got error: ", err)
		}
	} else {
		t.Fatal("err should be linebot APIError 401")
	}

	// non exist title case
	NonExistTestTitle := "NonExistTestTitle"
	err = app.processAddActivity("replyToken", "userID", NonExistTestTitle)
	if err != nil {
		if err.Error() != fmt.Sprintf("Not found \"%s\" in the user", NonExistTestTitle) {
			t.Fatal("err should be yarana-bot non-exist Koto error.", " Got error: ", err)
		}
	} else {
		t.Fatal("err should have error")
	}
}

// NoKotoYaranaDataCallMock is a test struct of DataCall for Yarana-bot
type NoKotoYaranaDataCallMock struct {
}

// NewNoKotoYaranaDataCallMock is a constructor of NoKotoYaranaDataCallMock
func NewNoKotoYaranaDataCallMock() (*NoKotoYaranaDataCallMock, error) {
	return &NoKotoYaranaDataCallMock{}, nil
}

// GetUsers is a method of DataCall interface
func (c *NoKotoYaranaDataCallMock) GetUsers() ([]*User, error) {
	return []*User{}, nil
}

// GetKotosByUserID is a method of DataCall interface
func (c *NoKotoYaranaDataCallMock) GetKotosByUserID(userID string) ([]*KotoData, error) {
	return []*KotoData{}, nil
}

// AddKoto is a method of DataCall interface
func (c *NoKotoYaranaDataCallMock) AddKoto(koto *KotoData) error {
	return nil
}

// GetActivitiesByKotoDataID is a method of DataCall interface
func (c *NoKotoYaranaDataCallMock) GetActivitiesByKotoDataID(kotoID string) ([]*ActivityData, error) {
	return nil, nil
}

// AddActivity is a method of DataCall interface
func (c *NoKotoYaranaDataCallMock) AddActivity(activity *ActivityData) error {
	return nil
}

func TestProcessGetKotosNoKoto(t *testing.T) {
	dataCall, _ := NewNoKotoYaranaDataCallMock()
	app, _ := NewYarana(
		"TEST_CHANNEL_SECRET",
		"TEST_CHANNEL_TOKEN",
		"TEST_ENDPOINT_BASE",
		dataCall,
	)

	err := app.processGetKotos("replyToken", "userID", "TestTitle")
	if err != nil {
		if err.Error() != "No Koto data in the user" {
			t.Fatal("err should be \"No koto data in the user\"", " Got error: ", err)
		}
	} else {
		t.Fatal("err should have error")
	}
}

func TestProcessAddKotoNoKoto(t *testing.T) {
	dataCall, _ := NewNoKotoYaranaDataCallMock()
	app, _ := NewYarana(
		"TEST_CHANNEL_SECRET",
		"TEST_CHANNEL_TOKEN",
		"TEST_ENDPOINT_BASE",
		dataCall,
	)

	// Standard case
	err := app.processAddKoto("replyToken", "userID", "NewTestTitle")
	if err != nil {
		if err.Error() != linebotAPIError401Str {
			t.Fatal("err should be linebot APIError 401", " Got error: ", err)
		}
	} else {
		t.Fatal("err should be linebot APIError 401")
	}
}

func TestProcessGetActivitiesNoKoto(t *testing.T) {
	dataCall, _ := NewNoKotoYaranaDataCallMock()
	app, _ := NewYarana(
		"TEST_CHANNEL_SECRET",
		"TEST_CHANNEL_TOKEN",
		"TEST_ENDPOINT_BASE",
		dataCall,
	)

	err := app.processGetActivities("replyToken", "userID", "TestTitle")
	if err != nil {
		if err.Error() != "No Koto data in the user" {
			t.Fatal("err should be \"No koto data in the user\"", " Got error: ", err)
		}
	} else {
		t.Fatal("err should have error")
	}
}

func TestProcessAddActivityNoKoto(t *testing.T) {
	dataCall, _ := NewNoKotoYaranaDataCallMock()
	app, _ := NewYarana(
		"TEST_CHANNEL_SECRET",
		"TEST_CHANNEL_TOKEN",
		"TEST_ENDPOINT_BASE",
		dataCall,
	)

	err := app.processAddActivity("replyToken", "userID", "TestTitle")
	if err != nil {
		if err.Error() != "No Koto data in the user" {
			t.Fatal("err should be \"No koto data in the user\"", " Got error: ", err)
		}
	} else {
		t.Fatal("err should have error")
	}
}

// NoActivityYaranaDataCallMock is a test struct of DataCall for Yarana-bot
type NoActivityYaranaDataCallMock struct {
}

// NewNoActivityYaranaDataCallMock is a constructor of NoActivityYaranaDataCallMock
func NewNoActivityYaranaDataCallMock() (*NoActivityYaranaDataCallMock, error) {
	return &NoActivityYaranaDataCallMock{}, nil
}

// GetUsers is a method of DataCall interface
func (c *NoActivityYaranaDataCallMock) GetUsers() ([]*User, error) {
	user1, err := NewUser("0123456789a")
	if err != nil {
		return nil, err
	}
	user2, err := NewUser("0123456789b")
	if err != nil {
		return nil, err
	}
	return []*User{user1, user2}, nil
}

// GetKotosByUserID is a method of DataCall interface
func (c *NoActivityYaranaDataCallMock) GetKotosByUserID(userID string) ([]*KotoData, error) {
	// Get Koto by userID from something
	id := "0123456789a"
	title := "TestTitle"
	koto, err := NewKotoData(id, userID, title)
	if err != nil {
		return nil, err
	}
	return []*KotoData{koto}, nil
}

// AddKoto is a method of DataCall interface
func (c *NoActivityYaranaDataCallMock) AddKoto(koto *KotoData) error {
	return nil
}

// GetActivitiesByKotoDataID is a method of DataCall interface
func (c *NoActivityYaranaDataCallMock) GetActivitiesByKotoDataID(kotoID string) ([]*ActivityData, error) {
	return []*ActivityData{}, nil
}

// AddActivity is a method of DataCall interface
func (c *NoActivityYaranaDataCallMock) AddActivity(activity *ActivityData) error {
	return nil
}

func TestProcessGetActivitiesNoActivity(t *testing.T) {
	dataCall, _ := NewNoActivityYaranaDataCallMock()
	app, _ := NewYarana(
		"TEST_CHANNEL_SECRET",
		"TEST_CHANNEL_TOKEN",
		"TEST_ENDPOINT_BASE",
		dataCall,
	)

	err := app.processGetActivities("replyToken", "userID", "TestTitle")
	if err != nil {
		if err.Error() != "No activity data in the user" {
			t.Fatal("err should be \"No activity data in the user\"", " Got error: ", err)
		}
	} else {
		t.Fatal("err should have error")
	}
}

func TestProcessAddActivityNoActivity(t *testing.T) {
	dataCall, _ := NewNoActivityYaranaDataCallMock()
	app, _ := NewYarana(
		"TEST_CHANNEL_SECRET",
		"TEST_CHANNEL_TOKEN",
		"TEST_ENDPOINT_BASE",
		dataCall,
	)

	// Standard case
	err := app.processAddActivity("replyToken", "userID", "TestTitle")
	if err != nil {
		if err.Error() != linebotAPIError401Str {
			t.Fatal("err should be linebot APIError 401", " Got error: ", err)
		}
	} else {
		t.Fatal("err should be linebot APIError 401")
	}
}
