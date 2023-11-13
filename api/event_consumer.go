package api

import (
	"bufio"
	"io"
	"log"
	"strings"
)

func EventConsumer(source io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(source)
	var streamResult []string

	for scanner.Scan() {
		event := scanner.Text()

		log.Println("recieved event")
		eventData := DataFromTextSSE(event)

        // temporary storage for arrayed data sa event 
		var arrayedEventLog []string

		// naay change nga walay sulod ang eventData
		// if naay sulod ang eventData, e process nato
		if len(eventData) > 0 {
			dataArray := strings.Split(eventData[0], ",")

			// data is like: ["foo":"bar" "baz":1]
			for _, entry := range dataArray {
				splittedEntry := strings.Split(entry, "\":")

				// entry in splitted is like: ["foo "bar"]
				for _, entr := range splittedEntry {
                    // remove all qoutes in each entry
					entr = strings.ReplaceAll(entr, "\"", "")

                    // push all entry to temp arrayedEventLog
					arrayedEventLog = append(arrayedEventLog, entr)
				}
			}

            // reasign a value sa streamResult to arrayedEventLog
            // para ang sulod sa streamResult ky ang lates event ra
			streamResult = arrayedEventLog
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return streamResult, nil
}
