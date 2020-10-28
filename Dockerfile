#adiciona o serviço do go e junto o git, bash e afins
FROM golang:alpine

ADD . /go/src/app
       
WORKDIR /go/src/app

RUN \
       apk add --no-cache bash git openssh && \
       go get -u github.com/lib/pq 
       

CMD ["go","run","cmd/main.go"]