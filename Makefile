BUILD_VERSION   := $(shell cat version)
BUILD_TIME      := $(shell date "+%F %T")
COMMIT_SHA1     := $(shell git rev-parse HEAD)

all: clean
	bash .cross_compile.sh

release: all
	ghr -u mritd -t ${GITHUB_RELEASE_TOKEN} -replace -recreate --debug ${BUILD_VERSION} dist

docker: all
	docker build -t mritd/certmonitor:${BUILD_VERSION} .

clean:
	rm -rf dist

install:
	go install

.PHONY : all release docker clean install