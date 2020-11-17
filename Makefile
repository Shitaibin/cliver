VERSION=0.1
GIT_COMMIT=$(shell git rev-list -1 HEAD)
BUILD_DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

.PHONY: build
build:
	echo $(GIT_VERSION)
	go build -ldflags "-X github.com/shitaibin/cliver/version.Version=$(VERSION)  -X github.com/shitaibin/cliver/version.GitCommit=$(GIT_COMMIT) -X github.com/shitaibin/cliver/version.BuildDate=$(BUILD_DATE)" .

.PHONY: clean
clean:
	rm cliver