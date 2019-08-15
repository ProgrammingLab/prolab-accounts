FROM golang:1.12.8-stretch as builder

ENV GO111MODULE on
WORKDIR /go/src/github.com/ProgrammingLab/prolab-accounts

RUN curl -Lo grapi https://github.com/izumin5210/grapi/releases/download/v0.4.0/grapi_linux_amd64
RUN chmod +x grapi

COPY . .
RUN ./grapi build

FROM gcr.io/distroless/base
LABEL maintainer="Ryota Egusa <egusa.ryota@gmail.com>"

COPY --from=builder /go/src/github.com/ProgrammingLab/prolab-accounts/bin/create-user /
COPY --from=builder /go/src/github.com/ProgrammingLab/prolab-accounts/bin/server /

CMD [ "./server" ]
