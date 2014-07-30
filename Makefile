test:
	go test -v

coverage: clean
	go test -v -coverprofile out.cover
	go tool cover -html=out.cover
	make clean

clean:
	rm -rf out.cover bot

.PHONY: clean test coverage
