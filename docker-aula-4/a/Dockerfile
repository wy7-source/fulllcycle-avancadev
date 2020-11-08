FROM golang:1.15

WORKDIR /go/src/avancadev-a

COPY . .
# Para baixar as dependencias do GO
RUN go get ./... 

RUN GOOS=linux go build

EXPOSE 9090

ENTRYPOINT ["./avancadev-a"]