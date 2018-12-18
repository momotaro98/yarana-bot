package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

// ZONE is time zone area
const ZONE string = "Asia/Tokyo" // TODO: be global

// TIMEDIFF is time difference of utc
const TIMEDIFF int = 9 * 60 * 60 // TODO: be global

// UserNonActiveDuration is hour period during which BOT determines that the user is inactive
const UserNonActiveDuration time.Duration = 18

// DefaultPushDisabled is boolean which the batch app pushes message of the Koto to the user
const DefaultPushDisabled bool = false

func main() {
	dataCall, err := NewYaranaDataCall(
		os.Getenv("YARANA_API_BASE_URL"),
		os.Getenv("YARANA_API_GETUSERS_KEY"),
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
	http.HandleFunc("/batch", app.Batch)
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
			log.Printf("Follow event: %v", event)
			if err := app.replyWithHelp(event.ReplyToken, "やらなボットよ、よろしくじゃないの！"); err != nil {
				log.Print(err)
			}
		case linebot.EventTypeUnfollow:
			log.Printf("Unfollowed this bot: %v", event)
		case linebot.EventTypeJoin:
			log.Printf("Join event: %v", event)
			if err := app.replyWithHelp(event.ReplyToken, "やらなボットよ、よろしくじゃないの！"); err != nil {
				log.Print(err)
			}
		case linebot.EventTypeLeave:
			log.Printf("Left: %v", event)
		case linebot.EventTypePostback:
			log.Printf("Postback event: %v", event)
			data := event.Postback.Data
			if data == "DATE" || data == "TIME" || data == "DATETIME" {
				data += fmt.Sprintf("(%v)", *event.Postback.Params)
			}
			if err := app.replyText(event.ReplyToken, "Got postback: "+data); err != nil {
				log.Print(err)
			}
		case linebot.EventTypeBeacon:
			log.Printf("Beacon event: %v", event)
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
	requestType, variableKeyword, err := AnalyzeInputText(message.Text)
	if err != nil {
		app.replyWithHelp(replyToken, "それじゃわからないわよ")
		return nil // not regard invalid input as error
	}
	switch requestType {
	case RequestTypeHelp:
		err = app.processHelp(replyToken, source.UserID, variableKeyword)
		if err != nil {
			return err
		}
	case RequestTypeGetKotos:
		err = app.processGetKotos(replyToken, source.UserID, variableKeyword)
		if err != nil {
			return err
		}
	case RequestTypeAddKoto:
		err = app.processAddKoto(replyToken, source.UserID, variableKeyword)
		if err != nil {
			return err
		}
	case RequestTypeGetActivities:
		err = app.processGetActivities(replyToken, source.UserID, variableKeyword)
		if err != nil {
			return err
		}
	case RequestTypeAddActivity:
		err = app.processAddActivity(replyToken, source.UserID, variableKeyword)
		if err != nil {
			return err
		}
	}
	return nil
}

func (app *Yarana) handleImage(message *linebot.ImageMessage, replyToken string) error {
	return nil
}

func (app *Yarana) processHelp(replyToken string, userID string, keyword string) error {
	// Make text to send
	var textToSend string
	textToSend = ReturnHelpText()
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(strings.TrimSpace(textToSend)),
	).Do(); err != nil {
		return err
	}
	return nil
}

func (app *Yarana) processGetKotos(replyToken string, userID string, keyword string) error {
	// Get Kotos
	kotos, err := app.dataCall.GetKotosByUserID(userID)
	if err != nil {
		app.replySorry(replyToken, fmt.Sprintf("ごめんなさいね、やることを取得するの失敗しちゃったみたいなのよ。もう一度入力してみて。"))
		return err
	}
	if len(kotos) == 0 || kotos == nil {
		app.replyWithHelp(replyToken, "やることが登録されてないわよ")
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
			app.replyWithHelp(replyToken, fmt.Sprintf("%sはもう登録されてるわよ", keyword)) // TODO: Replace showHelp to showAllUsersやつこと
			return fmt.Errorf("User was going to add duplicate Koto.")
		}
	}

	kotoToAdd, _ := NewKotoData("", userID, keyword, DefaultPushDisabled)
	errChan := make(chan error, 1)

	// Add Koto Data
	go func() {
		err := app.dataCall.AddKoto(kotoToAdd)
		errChan <- err
	}()
	err = <-errChan
	if err != nil {
		app.replySorry(replyToken, fmt.Sprintf("ごめんなさいね、%sを登録するのに失敗しちゃったみたいなのよ。もう一度入力してみて。", keyword))
		return err
	}

	// Make text to send
	var textToSend string
	textToSend = fmt.Sprintf("%sを新しく登録したわよ。ちゃんと続けなさいよね。", keyword)
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
		app.replyWithHelp(replyToken, "やることが登録されてないわよ")
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

	// Make text to send
	timezone := time.FixedZone(ZONE, TIMEDIFF)
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
			sort.SliceStable(acts, func(i, j int) bool {
				return acts[i].TimeStamp.After(acts[j].TimeStamp)
			})
			for _, act := range acts {
				// convert to correct time zone
				usersTimeStamp := act.TimeStamp.In(timezone)
				datetimeForUser := app.makeDatetimeToSendUser(usersTimeStamp)
				textToSend = textToSend + datetimeForUser + "\n"
			}
		}
	}
	if textToSend == "" { // "textToSend is empty" means the user has not activities
		app.replyWithHelp(replyToken, "まだ1回もやってないじゃないの。やりなさいよ。")
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
		app.replyWithHelp(replyToken, "やることが登録されてないわよ")
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
		app.replyWithHelp(replyToken, fmt.Sprintf("%sは登録されてないわよ", keyword)) // TODO: Replace showHelp to showAllUsersやつこと
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
		app.replySorry(replyToken, fmt.Sprintf("ごめんなさいね、%sの登録に失敗しちゃったみたいなのよ。もう一度入力してみて。", keyword))
		return err
	}
	// Make text to send
	var textToSend string
	textToSend = fmt.Sprintf("%sをやったのね、えらいじゃないの！", keyword)
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
	textToSend = message + "\n\n" + ReturnHelpText()
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(textToSend),
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

// Batch is hunde function for running a batch program
func (app *Yarana) Batch(w http.ResponseWriter, r *http.Request) {
	codes, ok := r.URL.Query()["code"]
	if !ok || len(codes) != 1 {
		log.Println("Batch kick key (URL param 'code') is missing")
		return
	}
	code := codes[0]
	// Check if the request url param is correct
	if code != app.getPushBatchKickKey() {
		log.Printf("Got wrong batch kick key (URL param 'code'): %s", string(code))
		return
	}
	app.RunBatch()
}

// RunBatch runs a batch program of yarana-bot
func (app *Yarana) RunBatch() error {
	// Get Users
	users, err := app.dataCall.GetUsers()
	if err != nil {
		return err
	}
	// Push message to users
	for _, user := range users {
		go app.RunPushBatch(user)
	}
	return nil
}

// RunPushBatch runs a batch program of yarana-bot
func (app *Yarana) RunPushBatch(user *User) error {
	// Get Kotos of the user
	allKotos, err := app.dataCall.GetKotosByUserID(user.ID)
	if err != nil {
		return err
	}
	if len(allKotos) == 0 || allKotos == nil {
		return nil // do nothing if the user has no kotos
	}

	// Filter Kotos whose pushDisabled is true
	var kotos []*KotoData
	for _, koto := range allKotos {
		if !koto.PushDisabled {
			kotos = append(kotos, koto)
		}
	}
	if len(kotos) == 0 { // TODO: Add unit test
		return nil // Do nothing if the all of pushDisabled flags are true
	}

	// Find kotos to send to users as notification
	var mutex = &sync.Mutex{}
	wg := &sync.WaitGroup{}
	var pushTargetKotoTitles []string
	for _, koto := range kotos {
		wg.Add(1)
		go func(kotoId string, kotoTitle string) {
			acts, _ := app.dataCall.GetActivitiesByKotoDataID(kotoId)
			if len(acts) < 1 {
				// If the koto doesn't have no activities, our bot sends notification.
				mutex.Lock()
				pushTargetKotoTitles = append(pushTargetKotoTitles, kotoTitle)
				mutex.Unlock()
			} else {
				if !CheckIfUserDidTheKotoInADay(acts) {
					// If a user don't do the koto in a day, our bot sends notification.
					mutex.Lock()
					pushTargetKotoTitles = append(pushTargetKotoTitles, kotoTitle)
					mutex.Unlock()
				}
			}
			wg.Done()
		}(koto.ID, koto.Title)
	}
	wg.Wait()

	if len(pushTargetKotoTitles) == 0 {
		return nil
	}

	// Make text to send and Push message to the user with package of the kotos
	var textToSend string
	textToSend = strings.Join(pushTargetKotoTitles, "と") + "は今日やったのかしら！？"
	textToSend = textToSend + "\n"
	textToSend = textToSend + "済ませたら"
	for _, kotoTitle := range pushTargetKotoTitles {
		textToSend = textToSend + "\"" + kotoTitle + "をやったよ" + "\""
	}
	textToSend = textToSend + "の入力をしてね"
	app.bot.PushMessage(user.ID, linebot.NewTextMessage(strings.TrimSpace(textToSend))).Do()
	log.Printf("Bot pushed message to userId, %s. Pushed message: %s", user.ID, textToSend)

	return nil
}

func (app *Yarana) getPushBatchKickKey() string {
	return os.Getenv("YARANA_PUSH_BATCH_KICK_KEY")
}

// CheckIfUserDidTheKotoInADay checks if taken activities has an activity which is done in a day.
func CheckIfUserDidTheKotoInADay(activities []*ActivityData) (didUserDoTheKotoInADay bool) {
	timezone := time.FixedZone(ZONE, TIMEDIFF)
	for _, act := range activities {
		usersTimeStamp := act.TimeStamp.In(timezone)
		if usersTimeStamp.After(time.Now().In(timezone).Add(-1 * time.Hour * UserNonActiveDuration)) {
			didUserDoTheKotoInADay = true
		}
	}
	return didUserDoTheKotoInADay
}
