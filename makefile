build:
	@go build -o bin/audiobook

start: build
	@./bin/audiobook

run:
	@gow -e=go,mod,html -i=target run .

tailwind:
	@tailwindcss -i views/styles/tailwind.css -o views/styles/output.css --watch

clean:
	@rm -rf bin/
	@rm audiobook
