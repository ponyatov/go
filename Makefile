SRC = main.go
gods.exe: $(SRC)	
	go build -o bin/$@ $^
