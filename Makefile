OUTPUT_DIR = ./build
OUTPUT = ${OUTPUT_DIR}/serves
MAIN = *.go
GO = go

.PHONY: all build

all: mkdir build

mkdir:
	mkdir -p ${OUTPUT_DIR} 2>/dev/null

build:
	${GO} build -o ${OUTPUT} ${MAIN}
	@echo "Build success."
	@echo "copy './${OUTPUT}' into your path."

clean:
	rm -rf ${OUTPUT} 2>/dev/null
