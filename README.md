# yarana-bot  やらなボット

Bot application in LINE to help you to do regularly

![yarana-bot_demo](https://raw.githubusercontent.com/momotaro98/my-project-images/master/yarana-bot/yarana-bot-v1.1.gif)

## Features

### Reply function

* Register what you want to do regularly
* Record it when you do
* Show you the records

## Usage

You can use the all of followings.

### Command

```
For instance, if you want to do "Running"
# Resigter
input "AddKoto Running"
# Record your activity
input "AddActivity Running"
# Confirm what you registered
input "GetKotos"
# See your activities
input "GetActivities"
# See help
input "Help"
```

### Japanese

```
例: "筋トレ"をやる場合
● 始めにやることを登録
"筋トレを登録して" と送信
● やった後は履歴を追加
"筋トレをやったよ" と送信
● やることを確認したいとき
"やること" と送信
● 過去履歴を見たいとき
"筋トレの履歴" と送信
★ 使い方を知りたいとき
"使い方" と送信
```


## Development

###  Build app

```
$ go get github.com/momotaro98/yarana-bot
$ cd $GOPATH/src/github.com/momotaro98/yarana-bot
$ make
```

### Run unit test

```
$ make test
```