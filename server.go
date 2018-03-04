package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

// KotoData is DTO of thing to do in Yarana-Bot
type KotoData struct {
	ID    string
	Title string
}

// NewKotoData is constructor of KotoData
func NewKotoData(id string, title string) (*KotoData, error) {
	return &KotoData{
		ID:    id,
		Title: title,
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
	GetKotoByTitle(title string) (*KotoData, error)
	AddKoto(koto *KotoData) error
	EditKoto(id string, koto *KotoData) (*KotoData, error)
	DeleteKoto(id string) error
	GetActivityByID(id string) (*ActivityData, error)
	GetActivitiesByKotoDataID(kotoID string) ([]*ActivityData, error)
	AddActivity(activity *ActivityData) error
}

// SimpleDataCall is a alternative of DataCall // TODO: interface for prototype
type SimpleDataCall interface {
	GetKotoByTitle(title string) (*KotoData, error)
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

// GetKotoByTitle is a method of DataCall interface
func (c *YaranaDataCallForTest) GetKotoByTitle(title string) (*KotoData, error) {
	koto, err := NewKotoData("0123456789a", title)
	if err != nil {
		return nil, err
	}
	return koto, nil
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

func main() {
	dataCall, err := NewYaranaDataCallForTest()
	if err != nil {
		log.Fatal(err)
	}
	app, err := NewYarana(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
		os.Getenv("APP_BASE_URL"),
		dataCall,
	)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/callback", app.Callback)
	// TODO: Support HTTPS by using `ListenAdnServeTLS`, reverse proxy or etc.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}

// Yarana app
type Yarana struct {
	bot        *linebot.Client
	appBaseURL string
	dataCall   SimpleDataCall
}

// NewYarana creates Yarana struct
func NewYarana(channelSecret, channelToken, appBaseURL string, dataCall SimpleDataCall) (*Yarana, error) {
	apiEndpointBase := os.Getenv("ENDPOINT_BASE")
	if apiEndpointBase == "" {
		apiEndpointBase = linebot.APIEndpointBase
	}
	bot, err := linebot.New(
		channelSecret,
		channelToken,
		linebot.WithEndpointBase(apiEndpointBase), // Usually you omit this.
	)
	if err != nil {
		return nil, err
	}
	return &Yarana{
		bot:        bot,
		appBaseURL: appBaseURL,
		dataCall:   dataCall,
	}, nil
}

// Callback function for http server
func (app *Yarana) Callback(w http.ResponseWriter, r *http.Request) {
	events, err := app.bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		log.Printf("Got event %v", event)
		log.Printf("EventSource Type: %v", event.Source.Type)
		log.Printf("EventSource UserID: %s", event.Source.UserID)
		log.Printf("EventSource GroupID: %s", event.Source.GroupID)
		log.Printf("EventSource RoomID: %s", event.Source.RoomID)
		switch event.Type {
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if err := app.handleText(message, event.ReplyToken, event.Source); err != nil {
					log.Print(err)
				}
			case *linebot.ImageMessage:
				if err := app.handleImage(message, event.ReplyToken); err != nil {
					log.Print(err)
				}
			default:
				log.Printf("Unknown message: %v", message)
			}
		case linebot.EventTypeFollow:
			if err := app.replyText(event.ReplyToken, "Got followed event"); err != nil {
				log.Print(err)
			}
		case linebot.EventTypeUnfollow:
			log.Printf("Unfollowed this bot: %v", event)
		case linebot.EventTypeJoin:
			if err := app.replyText(event.ReplyToken, "Joined "+string(event.Source.Type)); err != nil {
				log.Print(err)
			}
		case linebot.EventTypeLeave:
			log.Printf("Left: %v", event)
		case linebot.EventTypePostback:
			data := event.Postback.Data
			if data == "DATE" || data == "TIME" || data == "DATETIME" {
				data += fmt.Sprintf("(%v)", *event.Postback.Params)
			}
			if err := app.replyText(event.ReplyToken, "Got postback: "+data); err != nil {
				log.Print(err)
			}
		case linebot.EventTypeBeacon:
			if err := app.replyText(event.ReplyToken, "Got beacon: "+event.Beacon.Hwid); err != nil {
				log.Print(err)
			}
		default:
			log.Printf("Unknown event: %v", event)
		}
	}
}

func (app *Yarana) replyText(replyToken, text string) error {
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(text),
	).Do(); err != nil {
		return err
	}
	return nil
}

func (app *Yarana) handleText(message *linebot.TextMessage, replyToken string, source *linebot.EventSource) error {
	// Handle user input
	switch message.Text {
	}

	// Get Activities
	acts, err := app.dataCall.GetActivitiesByKotoDataID("123456789") // TODO: test code
	if err != nil {
		return err
	}
	actsString := fmt.Sprint(acts)

	// Make text to send
	// textToSend = message.Text // That's "Oumugaeshi"
	textToSend := actsString

	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(textToSend),
	).Do(); err != nil {
		return err
	}
	return nil
}

func (app *Yarana) handleImage(message *linebot.ImageMessage, replyToken string) error {
	return nil
}
