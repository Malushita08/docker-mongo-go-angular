#
#FROM golang:latest AS builder
#RUN apt-get update
#ENV GO111MODULE=on \
#    CGO_ENABLED=0 \
#    GOOS=linux \
#    GOARCH=amd64
#WORKDIR /go/src
#COPY go.mod .
#RUN go mod download
#COPY . .
#
#EXPOSE 8001
#
##ENV DB_USER="admin"
##ENV DB_PASSWORD="123"
##ENV DB_HOST="localhost"
##ENV DB_PORT="27017"
##ENV DB_NAME="db_books"
##ENV SERVER_PORT="8080"
##ENV GO_ENV="production"
#
#RUN go build main.go
#FROM scratch
#COPY --from=builder /go/src .
#ENTRYPOINT  ["./main"]

# Obtenga la última imagen de golang docker
FROM golang:latest
 # Cree un directorio dentro del contenedor para almacenar nuestra aplicación web
RUN mkdir -p /go/src/go_web
 # Luego conviértalo en un directorio de trabajo.
WORKDIR /go/src/go_web/
 # Copie el directorio go_web en el contenedor
COPY . /go/src/go_web/
 # Compilar, compilar en un archivo ejecutable
RUN go build /go/src/go_web/main.go
# Compilar, compilar en un archivo ejecutable
#RUN go run /go/src/go_web/main.go

 # Establecer la variable de entorno PORT
ENV PORT 8001
 # Exponga el puerto 80 al host para que las redes externas puedan acceder a su aplicación
EXPOSE 8001

ENV DB_USER="admin"
ENV DB_PASSWORD="123"
ENV DB_HOST="mongodb2"
ENV DB_PORT="27017"
ENV DB_NAME="db_books"
ENV SERVER_PORT="8001"
ENV GO_ENV="production"

 # Dile a Docker que inicie el comando para ejecutar el contenedor
ENTRYPOINT  ["./main"]
#CMD /go/src/go_web/main