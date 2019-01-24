FROM golang:1.11.1-stretch as builder

WORKDIR /go/src/github.com/ProgrammingLab/prolab-accounts

RUN curl -Lo grapi https://github.com/izumin5210/grapi/releases/download/v0.3.2/grapi_linux_amd64
RUN chmod +x grapi
RUN go get -v -u github.com/golang/dep/cmd/dep

COPY . .
RUN dep ensure -v -vendor-only
RUN ./grapi build

FROM gcr.io/distroless/base
LABEL maintainer="gedorinku <gedorinku@yahoo.co.jp>"

COPY --from=builder /go/src/github.com/ProgrammingLab/prolab-accounts/bin/server /

ENTRYPOINT [ "./server" ]
