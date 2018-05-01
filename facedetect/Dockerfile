FROM denismakogon/gocv-build-stage:edge as build-stage

LABEL maintainer="Radoslaw Kowalski. mail: radek.kwlsk@gmail.com"

ENV PROJECT_DIR=$GOPATH/src/github.com/afrometal/gocv-tests \
    PROJECT=facedetect \
    DEBUG=true

RUN go get github.com/afrometal/gocv-tests/$PROJECT

RUN go get -u -d gocv.io/x/gocv
RUN cd $PROJECT_DIR &&  go build -o $GOPATH/bin/$PROJECT ./$PROJECT/main.go

FROM denismakogon/gocv-runtime:edge

COPY --from=build-stage /go/bin/$PROJECT /$PROJECT
ENTRYPOINT ["/hello"]