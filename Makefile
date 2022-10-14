build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
	zip main.zip main

# curl --location --request POST '$YOUR_API_ENDPOINT' 
post:
	curl --location --request POST '' \
	--header 'Content-Type: application/json' \
	--data-raw '{"a":"b"}'