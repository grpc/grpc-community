ARTIFACTS=tally cmd/.git_hash

GO_SRC=$(shell find cmd pkg -name '*.go')

.PHONY: all
all: ${ARTIFACTS}

go.sum: go.mod
	go mod tidy

cmd/.git_hash:
	git rev-parse HEAD | tr -d '\n'  >"$@"

tally: ${GO_SRC} go.sum cmd/.git_hash
	echo ${GO_SRC}
	go build ./cmd/tally.go

.PHONY: clean
clean:
	rm -f ${ARTIFACTS}
