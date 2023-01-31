build_dir="build"

build:
	go build -o $(build_dir)/braingo ./cmd

clean:
	rm -rf $(build_dir)

vet:
	go vet ./internal/...
	go vet ./cmd/...

lint: vet
	golangci-lint run -D structcheck \
	-E bidichk \
	-E depguard \
	-E errname \
	-E errorlint \
	-E gocritic \
	-E goconst \
	-E revive \
	-E lll

