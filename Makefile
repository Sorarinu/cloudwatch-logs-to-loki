.PHONY: default
default: cloud-watch-logs-to-loki

cloud-watch-logs-to-loki: main.go
	env GOOS=linux GOARCH=amd64 go build -o cloud-watch-logs-to-loki main.go
	zip -r cloud-watch-logs-to-loki.zip ./cloud-watch-logs-to-loki

.PHONY: clean
clean:
	rm cloud-watch-logs-to-loki*
