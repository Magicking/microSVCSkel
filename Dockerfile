FROM golang
MAINTAINER Sylvain Laurent

ENV GOBIN $GOPATH/bin
ENV PROJECT_DIR github.com/Magicking/microSVCSkel
ENV PROJECT_NAME micro-s-v-c-server

ADD vendor /usr/local/go/src
ADD cmd /go/src/${PROJECT_DIR}/cmd
ADD models /go/src/${PROJECT_DIR}/models
ADD restapi /go/src/${PROJECT_DIR}/restapi

WORKDIR /go/src/${PROJECT_DIR}

RUN go build -v -o /go/bin/main /go/src/${PROJECT_DIR}/cmd/${PROJECT_NAME}/main.go
ENTRYPOINT /go/bin/main --host=0.0.0.0    --port=8090
