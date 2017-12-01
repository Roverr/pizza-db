alpine:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo \
		-o="./pizza-db"
linux:
	GOOS=linux GOARCH=amd64 go build -o="./pizza-db-linux64"
