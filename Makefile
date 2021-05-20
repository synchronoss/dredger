prefix=/usr/local

dredger: main.go components/*/*.go components/config/base_config.yaml
	go build -o $@ $<

install: dredger
	install dredger $(prefix)/bin/

test:
	go test -covermode=count "./components/..."

clean:
	rm -f dredger

.PHONY=test clean
