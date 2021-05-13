prefix=/usr/local

dredger: main.go components/*/*.go
	go build -o $@ $<

install: dredger
	install dredger $(prefix)/bin/

test:
	go test -covermode=count "./components/..."

clean:
	rm -f dredger

.PHONY=test clean
