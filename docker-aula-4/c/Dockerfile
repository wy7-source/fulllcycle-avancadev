FROM golang:1.15

WORKDIR /go/src/avancadev-c

COPY . .
# Para baixar as dependencias do GO
RUN go get ./... 

RUN GOOS=linux go build

EXPOSE 9092

ENTRYPOINT ["./avancadev-c"]