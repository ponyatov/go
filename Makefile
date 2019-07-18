SRC = main.go

run: bin/gods$(EXE)
	./$^
bin/gods$(EXE): $(SRC)	
	go build -o $@ $^
