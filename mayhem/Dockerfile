FROM golang:1.18 as builder

COPY . /iexcloud
WORKDIR /iexcloud/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN apt update && apt install -y clang


RUN go-fuzz-build -libfuzzer -o fuzzGetBytes.a && \
    clang -fsanitize=fuzzer fuzzGetBytes.a -o fuzzGetBytes.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /iexcloud/mayhem/fuzzGetBytes.libfuzzer /