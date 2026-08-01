build-lib-current:
	go build -o libcookbook.so -buildmode=c-shared ./cbindings
	patchelf --remove-rpath libcookbook.so

build-lib-386:
	GOOS=linux GOARCH=386 CGO_ENABLED=1 CC=$$CC_386 go build -buildmode=c-shared -o libcookbook.so ./cbindings
	patchelf --remove-rpath libcookbook.so

build-lib-arm64:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=1 CC=$$CC_ARM64 go build -buildmode=c-shared -o libcookbook.so ./cbindings
	patchelf --remove-rpath libcookbook.so

build-lib-arm7:
	GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 CC=$$CC_ARMV7 go build -buildmode=c-shared -o libcookbook.so ./cbindings
	patchelf --remove-rpath libcookbook.so

release-lib-current: build-lib-current
	mkdir -p out/current-os
	mv libcookbook.so out/current-os/
	mv libbw.h out/current-os/
	cp cbindings/bw_*.h out/current-os/

release-lib-386: build-lib-386
	mkdir -p out/386
	mv libcookbook.so out/386/
	mv libbw.h out/386/
	cp cbindings/bw_*.h out/386/

release-lib-arm7: build-lib-arm7
	mkdir -p out/arm7
	mv libcookbook.so out/arm7/
	mv libbw.h out/arm7/
	cp cbindings/bw_*.h out/arm7/

release-lib-arm64: build-lib-arm64
	mkdir -p out/arm64
	mv libcookbook.so out/arm64/
	mv libbw.h out/arm64/
	cp cbindings/bw_*.h out/arm64/

release-all: release-lib-current release-lib-arm7 release-lib-386 release-lib-arm64

# aliases

build-lib-armv7hl: build-lib-arm7
build-lib-aarch64: build-lib-arm64
build-lib-i486: build-lib-386