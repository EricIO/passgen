all: passgen

build: clean passgen

passgen:
	go build passgen.go

clean:
	rm -rf passgen
