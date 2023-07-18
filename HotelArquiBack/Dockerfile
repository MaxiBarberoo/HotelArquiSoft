FROM golang:alpine3.18

WORKDIR /usr/src/app

COPY HotelArquiBack ./
RUN go mod tidy

CMD sleep 10 && go run main.go
