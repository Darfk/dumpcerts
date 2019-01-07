SRC=dumpcerts.go

LDFLAGS="-w"

.PHONY: clean

clean:
	rm -r release/

releases:
	GOOS=linux   go build -ldflags $(LDFLAGS) -o release/linux64/dumpcerts $(SRC)
	GOOS=darwin  go build -ldflags $(LDFLAGS) -o release/osx/dumpcerts $(SRC)
	GOOS=windows go build -ldflags $(LDFLAGS) -o release/windows/dumpcerts.exe $(SRC)
