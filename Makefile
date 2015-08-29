build:
	cd static/css/;lessc style.less > style.css

test: build
	go run *.go
