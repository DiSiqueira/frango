FROM fabiorphp/golang-glide:1.8-alpine

RUN apk add --no-cache make

RUN apk add --no-cache git

ENV APP_DIR $GOPATH/src/github.com/disiqueira/frango/src/api

COPY ./Makefile ${APP_DIR}/Makefile
COPY ./glide.yaml ${APP_DIR}/glide.yaml

WORKDIR ${APP_DIR}

RUN make deps

COPY . ${APP_DIR}

CMD CompileDaemon -build="make install" -command="api"