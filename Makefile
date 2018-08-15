GOTEST_FLAGS=-cpu=1,2,4 -benchmem -benchtime=5s

TEXT_PKGS=Logrus Zap Gologging Zerolog Seelog Log15 Gokit
JSON_PKGS=Logrus Zap Zerolog Log15 Gokit

TEXT_PKG_TARGETS=$(addprefix test-text-,$(TEXT_PKGS))
JSON_PKG_TARGETS=$(addprefix test-json-,$(JSON_PKGS))

.PHONY: all deps test test-text test-json $(TEXT_PKG_TARGETS) $(JSON_PKG_TARGETS)

all: deps test

deps:
	go get  github.com/sirupsen/logrus
	go get  gopkg.in/inconshreveable/log15.v2
	go get  github.com/op/go-logging
	go get  github.com/cihub/seelog
	go get  github.com/go-kit/kit/log
	go get  github.com/rs/zerolog
	go get  go.uber.org/zap

test: test-text test-json

test-text: $(TEXT_PKG_TARGETS)

$(TEXT_PKG_TARGETS): test-text-%:
	go test $(GOTEST_FLAGS) -bench "$*.*Text"

test-json: $(JSON_PKG_TARGETS)

$(JSON_PKG_TARGETS): test-json-%:
	go test $(GOTEST_FLAGS) -bench "$*.*JSON"
