package main

// Reply messages
const (
	// Yarinasaiyo is a kind of joke message
	Yarinasaiyo = "やりなさいよ！"
)

// AnalyzeInputTextForJoke returns a joke reply message based on the user input text
func AnalyzeInputTextForJoke(text string) string {
	if replyMessage := analyzeInputTextForJokeInJapanese(text); replyMessage != "" {
		return replyMessage
	}
	return ""
}


func analyzeInputTextForJokeInJapanese(text string) (replyMessage string) {
	// Analyze for やりなさいよ！ reply
	if n := len(text) - 3*7; n >= 0 {
		if text[n:] == "をやってないよ" {
			return Yarinasaiyo
		}
	}
	if n := len(text) - 3*6; n >= 0 {
		if text[n:] == "をやってない" || text[n:] == "やってないよ" {
			return Yarinasaiyo
		}
	}
	if n := len(text) - 3*5; n >= 0 {
		if text[n:] == "やってない"  {
			return Yarinasaiyo
		}
	}

	return ""
}