package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	dataCall, err := NewYaranaDataCall(
		os.Getenv("YARANA_API_BASE_URL"),
		os.Getenv("YARANA_API_ADDKOTO_KEY"),
		os.Getenv("YARANA_API_ADDACTIVITY_KEY"),
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
	err = userReq.AnalyzeInputText(message.Text)
	if err != nil {
		app.replyWithHelp(replyToken, "I'm sorry that's invalid input for me.")
		return nil // not regard invalid input as error
	}
	switch userReq.Type {
	case RequstTypeGetKotos:
		err = app.processGetKotos(replyToken, source.UserID, userReq.VariableKeyword)
		if err != nil {
			return err
		}
	case RequstTypeAddKoto:
		err = app.processAddKoto(replyToken, source.UserID, userReq.VariableKeyword)
		if err != nil {
			return err
		}
	case RequstTypeGetActivities:
		err = app.processGetActivities(replyToken, source.UserID, userReq.VariableKeyword)
		if err != nil {
			return err
		}
	case RequstTypeAddActivity:
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
	if len(kotos) == 0 || kotos == nil {
		app.replyWithHelp(replyToken, "No Koto Data. Please add your やること.") // TODO: Show help
		return fmt.Errorf("No Koto data in the user")
	}

	// Make text to send
	var textToSend string
	for _, koto := range kotos {
		textToSend = textToSend + koto.Title + "\n"
	}
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(strings.TrimSpace(textToSend)),
	).Do(); err != nil {
		return err
	}
	return nil
}

func (app *Yarana) processAddKoto(replyToken string, userID string, keyword string) error {
	// Get Kotos at first
	kotos, err := app.dataCall.GetKotosByUserID(userID)
	if err != nil {
		return err
	}
	// Check if the koto is duplicate
	for _, koto := range kotos {
		if koto.Title == keyword {
			app.replyWithHelp(replyToken, fmt.Sprintf("You've already had the やること, %s", keyword)) // TODO: show user's all やること
			return fmt.Errorf("User was going to add duplicate Koto.")
		}
	}

	kotoToAdd, _ := NewKotoData("", userID, keyword)
	errChan := make(chan error, 1)

	// Add Koto Data
	go func() {
		err := app.dataCall.AddKoto(kotoToAdd)
		errChan <- err
	}()
	err = <-errChan
	if err != nil {
		app.replySorry(replyToken, fmt.Sprintf("I'm sorry I failed to add your new やること, %s.", keyword))
		return err
	}

	// Make text to send
	var textToSend string
	textToSend = fmt.Sprintf("I added your new やること, %s", keyword)
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(strings.TrimSpace(textToSend)),
	).Do(); err != nil {
		return err
	}

	return nil
}

func (app *Yarana) processGetActivities(replyToken string, userID string, keyword string) error {
	// Get Kotos at first
	kotos, err := app.dataCall.GetKotosByUserID(userID)
	if err != nil {
		return err
	}
	if len(kotos) == 0 || kotos == nil {
		app.replyWithHelp(replyToken, "No Koto Data. Please add your やること.") // TODO: Show help
		return fmt.Errorf("No Koto data in the user")
	}

	// Get specified Koto ID
	var specifiedKotoID string
	for _, koto := range kotos {
		if koto.Title == keyword {
			specifiedKotoID = koto.ID
		}
	}
	var kotoIDList []string
	if specifiedKotoID != "" {
		kotoIDList = append(kotoIDList, specifiedKotoID)
	} else {
		for _, koto := range kotos {
			kotoIDList = append(kotoIDList, koto.ID)
		}
	}
	// Get Activities in parallel
	activitiesChannel := make(chan []*ActivityData, len(kotoIDList))
	wg := &sync.WaitGroup{}
	for _, kotoID := range kotoIDList {
		wg.Add(1)
		go func(kotoId string) {
			acts, _ := app.dataCall.GetActivitiesByKotoDataID(kotoId)
			activitiesChannel <- acts
			wg.Done()
		}(kotoID)
	}
	wg.Wait()
	close(activitiesChannel)

	jst := time.FixedZone("Asia/Tokyo", 9*60*60) // TODO: Move this to proper place
	// Make text to send
	var textToSend string
	for acts := range activitiesChannel {
		if len(acts) > 0 {
			var kotoTitle string
			for _, koto := range kotos {
				if koto.ID == acts[0].KotoID {
					kotoTitle = koto.Title
				}
			}
			textToSend = textToSend + kotoTitle + "\n"
			for _, act := range acts {
				// convert to correct time zone
				usersTimeStamp := act.TimeStamp.In(jst)
				datetimeForUser := app.makeDatetimeToSendUser(usersTimeStamp)
				textToSend = textToSend + datetimeForUser + "\n"
			}
		}
	}
	if textToSend == "" { // "textToSend is empty" means the user has not activities
		app.replyWithHelp(replyToken, "You have no activities.") // TODO: show help
		return fmt.Errorf("No activity data in the user")
	}
	// Reply to user
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(strings.TrimSpace(textToSend)),
	).Do(); err != nil {
		return err
	}

	return nil
}

func (app *Yarana) processAddActivity(replyToken string, userID string, keyword string) error {
	// Get Kotos at first
	kotos, err := app.dataCall.GetKotosByUserID(userID)
	if err != nil {
		return err
	}
	if len(kotos) == 0 || kotos == nil {
		app.replyWithHelp(replyToken, "No Koto Data. Please add your やること.") // TODO: Show help
		return fmt.Errorf("No Koto data in the user")
	}

	// Get specified Koto ID
	var specifiedKotoID string
	for _, koto := range kotos {
		if koto.Title == keyword {
			specifiedKotoID = koto.ID
		}
	}
	// Stop process if koto.Title doesn't exist in the user's data
	if specifiedKotoID == "" {
		app.replyWithHelp(replyToken, fmt.Sprintf("You don't have the やること, %s", keyword)) // TODO: show user's all やること
		return fmt.Errorf("Not found \"%s\" in the user", keyword)
	}
	// Make a new Activity object
	activityToAdd, _ := NewActivityData("", specifiedKotoID, time.Now())
	// Add Activity Data
	errChan := make(chan error, 1)
	go func() {
		err := app.dataCall.AddActivity(activityToAdd)
		errChan <- err
	}()
	err = <-errChan
	if err != nil {
		app.replySorry(replyToken, fmt.Sprintf("I'm sorry I failed to add your %s activity.", keyword))
		return err
	}
	// Make text to send
	var textToSend string
	textToSend = fmt.Sprintf("I added your %s activity.", keyword)
	// Reply to user
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(strings.TrimSpace(textToSend)),
	).Do(); err != nil {
		return err
	}
	return nil
}

func (app *Yarana) replyWithHelp(replyToken string, message string) error {
	var textToSend string
	textToSend = message
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(textToSend), // TODO: Show HELP View to user
	).Do(); err != nil {
		return err
	}
	return nil
}

func (app *Yarana) replySorry(replyToken string, sorryMessage string) error {
	var textToSend string
	textToSend = sorryMessage
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(textToSend),
	).Do(); err != nil {
		return err
	}
	return nil
}

func (app *Yarana) makeDatetimeToSendUser(timestamp time.Time) string {
	datetimeStr := timestamp.String()
	if len(datetimeStr) > 16 {
		datetimeStr = datetimeStr[:16] // "2009-11-10 23:00" from 2009-11-10 23:00:00 +0000 UTC m=+0.000000001
	}
	return datetimeStr
}
