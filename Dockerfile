#Puxa a imagem do Golang do dockerhub
FROM golang:alpine
# Adiciona os arquivos do projeto na imagem
ADD . /go/src/app
       
WORKDIR /go/src/app
#Adiciona as dependencia na imagem
RUN \
       apk add --no-cache bash git openssh postgresql-client && \
       go get -u github.com/lib/pq && \
       go get -u github.com/Nhanderu/brdoc
       
#Executa o comando ao subir o container
CMD ["./wait-for-postgres.sh","db","go","run","cmd/main.go"]