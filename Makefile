test: vendor
	go test ./...

vendor:
	if [ -d ./internal/rules/testdata/src/vendor ]; then rm -r ./internal/rules/testdata/src/vendor; fi
	go mod vendor
	mv vendor ./internal/rules/testdata/src/
