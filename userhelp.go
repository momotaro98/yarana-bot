package main

// ReturnHelpText makes help description for user
func ReturnHelpText() (helpText string) {
	helpText =
		`使い方 ("筋トレ"をやる場合)

【始めにやることを登録】
"筋トレを登録して" と送信

【やった後は履歴を追加】
"筋トレをやったよ" と送信

【やることを確認したいとき】
"やることを教えて" と送信

【過去履歴を見たいとき】
"筋トレの履歴を教えて" または "履歴を教えて" と送信`

	return helpText
}
