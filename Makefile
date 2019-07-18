SRC = main.go
go:
	gofmt -w $(SRC)
	go run $(SRC)
