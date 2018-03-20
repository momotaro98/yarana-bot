package main

import (
	"encoding/json"
	"math/rand"
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
	ID        string    `json:"id"`
	KotoID    string    `json:"kotoId"`
	TimeStamp time.Time `json:"timestamp"`
}

// NewActivityData is constructor of ActivityData
func NewActivityData(id string, kotoID string, timeStamp time.Time) (*ActivityData, error) {
	return &ActivityData{
		ID:        id,
		KotoID:    kotoID,
		TimeStamp: timeStamp,
	}, nil
}

// DataCall is a alternative of DataCall // TODO: interface for prototype
type DataCall interface {
	GetKotosByUserID(userID string) ([]*KotoData, error)
	AddKoto(koto *KotoData) error
	GetActivitiesByKotoDataID(kotoID string) ([]*ActivityData, error)
	AddActivity(activity *ActivityData) error
}

// YaranaDataCall is a struct of DataCall for Yarana-bot
type YaranaDataCall struct {
	apiBaseURL        string
	idLen             int
	keyForAddKoto     string
	keyForAddActivity string
}

// NewYaranaDataCall is a constructor of YaranaDataCall
func NewYaranaDataCall(apiBaseURL string, keyForAddKoto string, keyForAddActivity string) (*YaranaDataCall, error) {
	return &YaranaDataCall{
		apiBaseURL:        apiBaseURL,
		idLen:             32,
		keyForAddKoto:     keyForAddKoto,
		keyForAddActivity: keyForAddActivity,
	}, nil
}

// GetKotosByUserID is a method of DataCall interface
func (c *YaranaDataCall) GetKotosByUserID(userID string) (kotos []*KotoData, err error) {
	baseURL := c.apiBaseURL + "kotos"
	url := AssembleURLWithParam(baseURL, "userId", userID)
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
	if koto.ID == "" || len(koto.ID) != c.idLen {
		koto.ID = c.GenerateUniqID()
	}
	// Create new Koto data of a user
	url := c.apiBaseURL + "koto"
	if c.keyForAddKoto != "" {
		url = AssembleURLWithParam(url, "code", c.keyForAddKoto)
	}
	jsonBytes, _ := json.Marshal(koto)
	err := HTTPPost(url, jsonBytes)
	if err != nil {
		return err
	}
	return nil
}

// GetActivitiesByKotoDataID is a method of DataCall interface
func (c *YaranaDataCall) GetActivitiesByKotoDataID(kotoID string) (activities []*ActivityData, err error) {
	baseURL := c.apiBaseURL + "activities"
	url := AssembleURLWithParam(baseURL, "kotoId", kotoID)
	body, err := HTTPGet(url)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &activities); err != nil {
		return nil, err
	}
	return activities, nil
}

// AddActivity is a method of DataCall interface
func (c *YaranaDataCall) AddActivity(activity *ActivityData) error {
	if activity.ID == "" || len(activity.ID) != c.idLen {
		activity.ID = c.GenerateUniqID()
	}
	// Create new Activity data
	url := c.apiBaseURL + "activity"
	if c.keyForAddActivity != "" {
		url = AssembleURLWithParam(url, "code", c.keyForAddActivity)
	}
	jsonBytes, _ := json.Marshal(activity)
	err := HTTPPost(url, jsonBytes)
	if err != nil {
		return err
	}
	return nil
}

// init func for rand generating
func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateUniqID generates uniq id chars
func (c *YaranaDataCall) GenerateUniqID() (id string) {
	var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, 32)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	id = string(b)
	return id
}
