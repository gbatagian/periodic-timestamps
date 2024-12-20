run:
	go run .

test:
	go test -count=1 ./...

coverage:
	go test -count=1 -cover ./...

coverage-profile:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out  