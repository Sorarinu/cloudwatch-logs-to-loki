![facebook_cover_photo_1](https://user-images.githubusercontent.com/16132069/73810242-3479c800-4819-11ea-87f4-0f0747a9cc05.png)

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
| LOKI_ENDPOINT | http://xxx.xxx.xxx.xxx/loki/api/v1/push |
| SERVICE_NAME  | Unique service name                      |

Register Lambda to the log group subscription you want to send to Loki.

### Clean

```bash
$ make clean
```

## License
MIT, see [LICENSE](LICENSE).
