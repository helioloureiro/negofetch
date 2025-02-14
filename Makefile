SOURCES = main.go logos.go
SOURCES += utils/utils.go
SOURCES += linux/linux.go
SOURCES += posix/posix.go
SOURCES += macos/macos.go

BINARY = negofetch

BUILD_OPTIONS = -modcacherw
#BUILD_OPTIONS += -race
BUILD_OPTIONS += -ldflags="-w -X 'main.Version=$$(git tag -l --sort taggerdate | tail -1)'"
BUILD_OPTIONS += -buildmode=pie
BUILD_OPTIONS += -tags netgo,osusergo
BUILD_OPTIONS += -trimpath

all: $(SOURCES) dependencies $(BINARY)

dependencies:
	go mod tidy

$(BINARY): $(SOURCES)
	env GOAMD64=v2 \
		CGO_ENABLED=1 \
	go build $(BUILD_OPTIONS) .

tagpush: all
	./bin/stepup_tag.sh
	git push origin HEAD
	git push origin HEAD --tags

run:
	go run $(BUILD_OPTIONS) .

clean:
	rm -f $(BINARY)
