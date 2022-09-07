build:
	mkdir -p target/bin
	go build -o target/bin/tenprint main.go

install: build
	cp target/bin/tenprint /usr/bin/tenprint

uninstall:
	rm /usr/bin/tenprint
