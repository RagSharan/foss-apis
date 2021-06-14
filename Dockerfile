FROM golang:latest

LABEL maintainer="Quique <email@domain.com>"

#RUN mkdir /app
#COPY . /app

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8000

RUN go build 

CMD ["./foss-apis"]