ARCH=arm64
OS=darwin
BINARY_NAME=validator
MAIN_PATH=cmd/yaml_validator
.DEFAULT_GOAL := run

build:
	GOARCH=${ARCH} GOOS=${OS} go build -o ${MAIN_PATH}/${BINARY_NAME} ${MAIN_PATH}/main.go
	chmod +x ${MAIN_PATH}/${BINARY_NAME}

test: build
	./${MAIN_PATH}/${BINARY_NAME} \
    --schema=$(shell pwd)/resources/schema.json \
    --dir=$(shell pwd)/resources \
    --yaml=smoke.yml,regress.yml

clean:
	rm -f ${MAIN_PATH}/${BINARY_NAME}
