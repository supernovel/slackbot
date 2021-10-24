// +heroku goVersion go1.14
// +heroku install ./cmd/server.go

module github.com/supernovel/slackbot

go 1.14

require (
	github.com/joho/godotenv v1.4.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/slack-go/slack v0.9.5
	github.com/stretchr/testify v1.3.0 // indirect
)
