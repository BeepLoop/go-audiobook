build:
	@go build

start: build
	@./audiobook

run:
	@gow -e=go,mod,html -i=target run .

clean:
	@rm -rf bin/
	@rm audiobook
