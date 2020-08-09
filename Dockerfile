FROM golang:1.14.7-stretch as builder

ENV GO111MODULE on
WORKDIR /go/src/github.com/ProgrammingLab/prolab-accounts

RUN curl -Lo grapi_linux_amd64.tar.gz https://github.com/izumin5210/grapi/releases/download/v0.5.0/grapi_linux_amd64.tar.gz \
    && tar -xzf grapi_linux_amd64.tar.gz \
    && mv ./grapi_linux_amd64/grapi . \
    && chmod +x grapi \
    && rm -f grapi_linux_amd64.tar.gz

COPY . .
RUN ./grapi build

FROM gcr.io/distroless/base
LABEL maintainer="Ryota Egusa <egusa.ryota@gmail.com>"

COPY --from=builder /go/src/github.com/ProgrammingLab/prolab-accounts/bin/create-user /
COPY --from=builder /go/src/github.com/ProgrammingLab/prolab-accounts/bin/server /

CMD [ "./server" ]
