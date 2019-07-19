SRC = main.go
EDS = bin/eds$(EXE)

run: $(EDS)
	./$^
$(EDS): $(SRC)	
	go build -o $@ $^

packages: bin/gocode$(EXE)
bin/gocode$(EXE):
	go get github.com/nsf/gocode

