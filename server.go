package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	app, err := NewYarana(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
		os.Getenv("APP_BASE_URL"),
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
}

// NewYarana creates Yarana struct
func NewYarana(channelSecret, channelToken, appBaseURL string) (*Yarana, error) {
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
	switch message.Text {
	}
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(message.Text),
	).Do(); err != nil {
		return err
	}
	return nil
}

func (app *Yarana) handleImage(message *linebot.ImageMessage, replyToken string) error {
	return nil
}
