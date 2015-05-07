FROM       alpine
MAINTAINER Johannes 'fish' Ziemke <fish@docker.com> (@discordianfish)

ENV  GOPATH /go
ENV APPPATH $GOPATH/src/github.com/docker-infra/docker-backup
COPY . $APPPATH
RUN apk add --update -t build-deps go git && cd $APPPATH \
    && go get -d && go build -o /bin/docker-backup \
    && mkdir /docker-backup \
    && ln -s /bin/docker-backup /docker-backup/docker-backup \
    && apk del --purge build-deps && rm -rf $GOPATH

WORKDIR    /docker-backup
ENTRYPOINT [ "/bin/docker-backup" ]
