ARTIFACTS=tally cmd/.git_hash

.PHONY: all
all: ${ARTIFACTS}

go.sum: go.mod
	go mod tidy

cmd/.git_hash:
	git rev-parse HEAD | tr -d '\n'  >"$@"

tally: cmd/tally.go go.sum cmd/.git_hash
	go build ./cmd/tally.go

.PHONY: clean
clean:
	rm -f ${ARTIFACTS}
