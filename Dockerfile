FROM golang:1.10.3 as builder
WORKDIR /go/src/github.com/bgzzz/go-triangle/
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
COPY . /go/src/github.com/bgzzz/go-triangle/
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM scratch
COPY --from=builder /go/src/github.com/bgzzz/go-triangle/go-triangle /.
ENTRYPOINT ["/go-triangle"]
