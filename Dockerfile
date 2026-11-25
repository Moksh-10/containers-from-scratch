FROM golang:1.25-bookworm

WORKDIR /app

COPY go.mod ./
COPY . .

RUN go build -o container .

ENTRYPOINT ["./container"]
