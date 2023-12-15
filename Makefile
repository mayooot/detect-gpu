BINARY = detect-gpu
GOARCH = amd64

GITHUB_USER = mayooot
CURRENT_DIR =$(shell pwd)
BUILD_DIR=${CURRENT_DIR}/cmd/${BINARY}
BIN_DIR=${CURRENT_DIR}/bin

linux:
	@cd ${BUILD_DIR}; \
	GOOS=linux GOARCH=${GOARCH} go build -o ${BIN_DIR}/${BINARY}-linux-${GOARCH} . ; \
	cd - >/dev/null


clean:
	- rm -f ${BIN_DIR}/*

fmt:
	gofmt -l -w .

imports:
	goimports-reviser --rm-unused -local github.com/${GITHUB_USER}/${BINARY} -format ./...

.PHONY: linux clean fmt imports