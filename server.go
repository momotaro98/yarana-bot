package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	dataCall, err := NewYaranaDataCall(
		os.Getenv("YARANA_API_BASE_URL"),
		os.Getenv("YARANA_API_ADDKOTO_KEY"),
	)
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
	dataCall   DataCall
}

// NewYarana creates Yarana struct
func NewYarana(channelSecret, channelToken, appBaseURL string, dataCall DataCall) (*Yarana, error) {
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

func (app *Yarana) handleText(message *linebot.TextMessage, replyToken string, source *linebot.EventSource) (err error) {
	// Analyze text message
	userReq := NewUserTextRequest()
	userReq.AnalyzeInputText(message.Text)
	switch reqType := userReq.Type; reqType {
	case "GetKotos":
		err = app.processGetKotos(replyToken, source.UserID, userReq.VariableKeyword)
		if err != nil {
			return err
		}
	case "AddKoto":
		err = app.processAddKoto(replyToken, source.UserID, userReq.VariableKeyword)
		if err != nil {
			return err
		}
	case "GetActivities":
		err = app.processGetActivities(replyToken, source.UserID, userReq.VariableKeyword)
		if err != nil {
			return err
		}
	case "AddActivity":
		err = app.processAddActivity(replyToken, source.UserID, userReq.VariableKeyword)
		if err != nil {
			return err
		}
	}
	return nil
}

func (app *Yarana) handleImage(message *linebot.ImageMessage, replyToken string) error {
	return nil
}

func (app *Yarana) processGetKotos(replyToken string, userID string, keyword string) error {
	// Get Kotos
	kotos, err := app.dataCall.GetKotosByUserID(userID)
	if err != nil {
		return err
	}

	// Make text to send
	var textToSend string
	if len(kotos) == 0 || kotos == nil {
		textToSend = "No Koto Data"
	} else {
		textToSend = kotos[0].Title
	}
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(textToSend),
	).Do(); err != nil {
		return err
	}
	return nil
}

func (app *Yarana) processAddKoto(replyToken string, userID string, keyword string) error {
	kotoToAdd, _ := NewKotoData("", userID, keyword)
	errChan := make(chan error, 1)

	// Add Koto Data
	go func() {
		err := app.dataCall.AddKoto(kotoToAdd)
		errChan <- err
	}()

	var textToSend string

	err := <-errChan
	if err != nil {
		textToSend = "I'm sorry I failed to add your new やること."
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewTextMessage(textToSend),
		).Do(); err != nil {
			return err
		}
		return err // TODO: I wonder the `err` scope. app.bot.ReplyMessage error or app.dataCall.AddKoto error?
	}
	textToSend = "I added your new やること"
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(textToSend),
	).Do(); err != nil {
		return err
	}

	return nil
}

func (app *Yarana) processGetActivities(replyToken string, userID string, keyword string) error {
	/*
		acts, err := app.dataCall.GetActivitiesByKotoDataID("123456789") // TODO: test code
		if err != nil {
			return err
		}
		actsString := fmt.Sprint(acts)
	*/
	return nil
}

func (app *Yarana) processAddActivity(replyToken string, userID string, keyword string) error {
	return nil
}

// UserTextRequest is struct for managing input text from user
type UserTextRequest struct {
	Type            string
	VariableKeyword string
}

// NewUserTextRequest is constructor of KotoData
func NewUserTextRequest() *UserTextRequest {
	return &UserTextRequest{}
}

// AnalyzeInputText analyzes input text from user
func (r *UserTextRequest) AnalyzeInputText(text string) error {
	// TODO: Implement analyzer
	// Recognize "GetKotos" if the text starts with "GetKoto"
	words := strings.Fields(text)
	if len(words) < 1 {
		return fmt.Errorf("Can't analyze: %s", text)
	}

	switch fWord := words[0]; fWord {
	case "GetKotos":
		r.Type = "GetKotos"
	case "AddKoto":
		if len(words) < 2 {
			return fmt.Errorf("Can't analyze: %s", text)
		}
		r.Type = "AddKoto"
		r.VariableKeyword = words[1]
	}

	return nil
}
