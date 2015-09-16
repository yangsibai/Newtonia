PORT := 8080

build:
	cd static/css/;lessc style.less > style.css

test: build
	go run *.go

run: stop
	nohup go run *.go>/dev/null 2>&1 &

stop:
	lsof -i tcp:${PORT} | awk 'NR!=1 {print $$2}' | xargs kill

.PHONY: build, run, test
