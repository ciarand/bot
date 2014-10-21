quicktest:
	go test -v ./...

test:
	./bin/test

cover: test
	go tool cover -html=profile.cov

clean:
	rm -rf profile.cov

.PHONY: quicktest test clean check
