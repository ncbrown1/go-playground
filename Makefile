PID      = /tmp/go-playground.pid
GO_FILES = $(wildcard *.go) $(shell find ./app -name '*.go')
APP      = ./server

serve: restart
	@fswatch -o ./app | xargs -n1 -I{}  make restart || make kill

kill:
	@kill `cat $(PID)` || true

before:
	@:

$(APP): $(GO_FILES)
	@go build -o server server.go

restart: kill before $(APP)
	@./server -p 8000 & echo $$! > $(PID)

.PHONY: serve restart kill before