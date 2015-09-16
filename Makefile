PORT := 8080

build:
	go build -o Newtonia

test: build
	go run *.go

run: stop build
	nohup ./Newtonia>/dev/null 2>&1 &

stop:
	-lsof -t -i:${PORT} | xargs kill

.PHONY: build, run, test
