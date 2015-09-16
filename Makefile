PORT := 8080

build:
	cd static/css/;lessc style.less > style.css

test: build
	go run *.go

run: stop
	nohup go run *.go>/dev/null 2>&1 &

stop:
	-lsof -t -i:${PORT} | xargs kill

.PHONY: build, run, test
