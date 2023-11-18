# go-audiobook
## Description
A web application made with Go as a collaborative project with education students as a reading intervention tool. It allow admins to upload a story, upload a recorded audio for the story, cover image for the story, and a list of words considered difficult to pronounce. The list of words are converted into audio using [PlayHT](https://play.ht/).

## Built with
- [Go](https://go.dev/) - A fast and simple programming language designed by Google.
- [PlayHT](https://play.ht/) - A text to audio generation API.
- [Tailwindcss](https://tailwindcss.com/) - A utility-first CSS framework.
- [Gin](https://gin-gonic.com/) - A web framework for Go that features a Martini-like API.

## Requirements
go-audiobook requires a version of [Go](https://go.dev/) to be installed. If you want to edit the contents, design, and add pages, you must have a [Tailwindcss](https://tailwindcss.com/) installed. Visit [here](https://tailwindcss.com/docs/installation) for a guide on how to install Tailwindcss. The project is built using Tailwindcss standalone executable, thus it is the recommended one to use.

## How to run
1. Clone the project
```bash
git clone https://github.com/BeepLoop/go-audiobook.git
```

2. Install the dependencies
```bash
go mod tidy
```

3. Setup the config file
 - Open the `voice_config.json` inside the `config` directory
 - Update the `playht_user` and `api_key` to your api keys from PlayHT. Visit [PlayHT](https://play.ht/) for a guide on how to setup an account.

4. Run the project
```bash
go run main.go
```

5. Updating styling
  - Run Tailwindcss in watch mode
    ```bash
    tailwindcss -i ./views/styles/tailwind.css -o ./views/styles/output.css --watch
    ```
  - Start editing the HTML pages

## Building
- Using the makefile
```bash
make build
```

- Using Go
```bash
go build
```


