# syntax=docker/dockerfile:1

FROM golang
WORKDIR /app
COPY . .

LABEL version="1.0" 
LABEL creators="@arturzhamaliyev @Pashtetium  @darzox"
RUN go build -o cmd/forum cmd/main.go
EXPOSE 8081

CMD [ "cmd/forum" ]

