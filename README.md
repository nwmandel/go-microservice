# go-microservice
A basic RESTful microservice using gokit

Time service to either:

- Get time in format `YYYY-MM-DD HH:mm:ss`.
- Validate time in format `YYYY-MM-DD HH:mm:ss`.

## Start Service
To run locally
```bash
go run main.go
```

### Send Payloads Locally
Request | Response |
--- | --- |
GET `curl http://localhost:8080/get` | `{"date":"2020-04-13 10:50:59"}` |
GET `curl http://localhost:8080/status` | `{"status":"ok"}` |
POST `curl -POST -d '{"date":"2020-04-01 15:04:05"}' http://localhost:8080/validate` | `{"valid":true}` |
POST `curl -POST -d '{"date":"32/12/2020"}' http://localhost:8080/validate` | `{"valid":false,"err":"parsing time \"32/12/2020\" as \"2006-01-02 15:04:05\": cannot parse \"2/2020\" as \"2006\""}` |

### Test service
```bash
go test ./service/
```
