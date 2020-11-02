# Josh's HTTP API Service
A simple GoLang (`1.15.3`) http server that reads from a sqlite database file and returns json data.
For reference/example purposes.

LIVE ENDPOINT: https://api.joshbacon.name/employees

## TL;DR
```bash
export PORT=8080
go run main.go
```
* Test command: `curl localhost:8080/employees

## DOCKER STUFF

### 8. Docker Build
```bash
docker build -t api-joshbacon-name:latest .
```

### 10. Docker Run
```bash
docker run -e PORT=8080 -p 8080:8080 api-joshbacon-name:latest
```

### 9. Docker Publish
```bash
docker tag api-joshbacon-name:latest gcr.io/rich-involution-105622/api-joshbacon-name:latest
docker push gcr.io/rich-involution-105622/api-joshbacon-name:latest
```

## DEPLOYMENT

### 11. Google Cloud Run 
```bash
gcloud run deploy api-joshbacon-name \
--image gcr.io/rich-involution-105622/api-joshbacon-name:latest \
--platform=managed \
--region=us-west1
```

## OTHER USEFUL GOLANG COMMANDS

### 1. Linting - Prints out style mistakes
```
go get -u golang.org/x/lint/golint
~/go/bin/golint -set_exit_status ./...
```

### 2. Vet - Examines Go source code and reports suspicious constructs (`go doc cmd/vet`)
```
go vet ./... 2> go-vet-report.out
```

### 3. Vendor  - Adds/Updates all dependencies to the `vendor/` directory
```
go mod vendor
```

### 4. Test - Run all tests (for all included go modules)
```bash
go test -mod vendor ./... -json > go-test-report.json
go test -mod vendor ./... -coverprofile=go-test-coverage.out
```
- Test logs are printed to `STDOUT` and `go-test-report.json`
- Test coverage report is printed to `go-test-coverage.json`

### 7. Build Executable
```bash
go build -o executable -mod vendor ./...
```