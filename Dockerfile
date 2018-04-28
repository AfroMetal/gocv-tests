FROM denismakogon/gocv-build-stage:edge as build-stage

LABEL maintainer="Denis Makogon. mail: lildee1991@gmail.com"

ENV PROJECT_DIR=/go/src/github.com/afrometal/gocv-tests \
    DEBUG=true

WORKDIR $PROJECT_DIR

RUN go get github.com/afrometal/gocv-tests

RUN go get -u -d gocv.io/x/gocv
RUN go build -o $GOPATH/bin/hello ./hello/main.go

FROM denismakogon/gocv-runtime:edge

COPY --from=build-stage /go/bin/hello /hello
ENTRYPOINT ["/hello"]