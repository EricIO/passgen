all: passgen

build: clean passgen

passgen:
	go build passgen.go

install:
	mv ./passgen /usr/local/bin
	gzip -k ./passgen.1
	install -g 0 -o 0 -m 0644 ./passgen.1.gz /usr/local/share/man/man1

clean:
	rm -rf passgen
