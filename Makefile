build:
	cd static/css/;lessc style.less > style.css

test: build
	go run *.go

run:
	nohup go run *.go>/dev/null 2>&1 &
