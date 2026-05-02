FROM golang:1.25 AS builder

COPY . $GOPATH/src/github.com/goinvest/iexcloud/v2
WORKDIR $GOPATH/src/github.com/goinvest/iexcloud/v2/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN GOFLAGS=-mod=mod go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN apt update && apt install -y clang

RUN go-fuzz-build -libfuzzer -o fuzzGetBytes.a github.com/goinvest/iexcloud/v2/mayhem && \
    clang -fsanitize=fuzzer fuzzGetBytes.a -o fuzzGetBytes.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /go/src/github.com/goinvest/iexcloud/v2/mayhem/fuzzGetBytes.libfuzzer /
