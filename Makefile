
.PHONY: test
test:
	go test

.PHONY: benchmark
benchmark:
	go test -bench=.
