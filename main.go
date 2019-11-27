package main

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"fmt"
)

// pushLogToLoki is send message to Loki.
func pushLogToLoki(url string, message string, service string) {
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	message = formatMessage(message)
	param := makeRequestBody(timestamp, message, service)

	// Execute Loki API.
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(param),
	)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("response: ", resp.Status, string(body))
}

// formatMessage is removes / replaces unnecessary strings from log messages.
func formatMessage(message string) string {
	// Remove line break at end of line.
	message = strings.TrimRight(message, "\n")

	// Delete the date (YYYY/MM/DD hh:mm:ss) if it is included
	re := regexp.MustCompile(`^\d{4}/\d{1,2}/\d{1,2} \d{1,2}:\d{1,2}:\d{1,2}`)
	match := re.FindStringSubmatch(message)
	if len(match) != 0 {
		message = strings.Replace(message, match[0], "", -1)
	}

	// Replace tab characters with commas.
	message = strings.Replace(message, "\t", ",", -1)
	message = strings.TrimRight(message, ",")

	return message
}

// makeRequest Body is generates a body for sending requests to Loki.
func makeRequestBody(timestamp string, message string, service string) []byte {
	tags := "\"_service\":\"" + service + "\""

	slice := strings.Split(message, ",")

	// Decompose a colon-delimited message into Key and Value and add it to the tag.
	for _, str := range slice {
		s := strings.Split(str, ": ")

		if len(s) != 0 {
			k := "_" + strings.TrimSpace(s[0])
			k = strings.ToLower(k)
			k = strings.Replace(k, " ", "_", -1)

			v := strings.TrimSpace(s[1])

			tags += ",\"" + k + "\":\"" + v + "\""
		}
	}

	param := []byte(`
	{
		"streams":[
			{
				"stream":{
					` + tags + `
				},
				"values":[
					["` + timestamp + `","` + message + `"]
				]
			}
		]
	}`)

	return param
}

func handle(ctx context.Context, logsEvent events.CloudwatchLogsEvent) {
	url := os.Getenv("LOKI_ENDPOINT")
	service := os.Getenv("SERVICE_NAME")

	data, _ := logsEvent.AWSLogs.Parse()
	for _, logEvent := range data.LogEvents {
		pushLogToLoki(url, logEvent.Message, service)
	}
}

func main() {
	lambda.Start(handle)
}
