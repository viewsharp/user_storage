FROM golang:1.16

WORKDIR /app
COPY . /app

RUN go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.14.1 \
 && go build -o server cmd/server/*

ENV DATABASE_PATH=/app/db.sqlite \
    PORT=8000

EXPOSE 8000

CMD migrate -source file:///app/migrations -database sqlite3://${DATABASE_PATH} up \
 && ./server
