FROM golang:1.20 AS builder

WORKDIR /tmp/build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go install .

WORKDIR /tmp/alpha

CMD ["tail", "-f", "/dev/null"]