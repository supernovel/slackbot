// +heroku goVersion go1.14
// +heroku install ./cmd/server.go

module github.com/supernovel/slackbot

go 1.14

require (
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/slack-go/slack v0.6.4
	go.mongodb.org/mongo-driver v1.3.3
)
