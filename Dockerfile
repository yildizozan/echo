FROM golang:1.18.8-alpine3.16 as builder

ENV GOBIN=/go/bin

WORKDIR /go/src/github.com/yildizozan/echo

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go install

FROM gcr.io/distroless/static-debian11@sha256:ed05c7a5d67d6beebeba19c6b9082a5513d5f9c3e22a883b9dc73ec39ba41c04

COPY --from=builder /go/bin/echo /

CMD ["/echo"]