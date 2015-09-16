PORT := 8080

build:
	go build

test: build
	go run *.go

run: stop
	nohup ./newtonia>/dev/null 2>&1 &

stop:
	-lsof -t -i:${PORT} | xargs kill

.PHONY: build, run, test
