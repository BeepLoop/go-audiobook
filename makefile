build:
	@go build

start: build
	@./go-audiobook

run:
	@gow -e=go,mod,html -i=target run .

clean:
	@rm -rf bin/
	@rm assets/downloads/*
	@rm assets/audios/*
	@rm assets/thumbnails/*
	@rm data/*
