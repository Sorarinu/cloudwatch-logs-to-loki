# CloudWatchLogsToLoki

Lambda function to send tab-delimited log output to CloudWatch Logs to Grafana Loki.

## Usage

Run the Build.

```bash
$ make
```

It works if you register the completed zip file with AWS Lambda.


The following environment variables must be set.

| Key           | Value                                    |
| :---          | :---                                     |
| LOKI_ENDPOINT | http://<loki server ip>/loki/api/v1/push |
| SERVICE_NAME  | Unique service name                      |

Register Lambda to the log group subscription you want to send to Loki.

### Clean

```bash
$ make clean
```

## License
MIT, see [LICENSE](LICENSE).