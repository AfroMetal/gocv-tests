FROM denismakogon/gocv-build-stage:edge as build-stage

LABEL maintainer="Radoslaw Kowalski. mail: radek.kwlsk@gmail.com"

ENV PROJECT_DIR=$GOPATH/src/github.com/radekwlsk/gocv-tests \
    PROJECT=hello \
    DEBUG=true

RUN go get github.com/radekwlsk/gocv-tests/hello

RUN go get -u -d gocv.io/x/gocv
RUN cd $PROJECT_DIR &&  go build -o $GOPATH/bin/hello ./hello/main.go

FROM denismakogon/gocv-runtime:edge

COPY --from=build-stage /go/bin/hello /hello
ENTRYPOINT ["/hello"]