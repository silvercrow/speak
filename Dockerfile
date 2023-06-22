FROM golang:1.20.5

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

ARG say=hello
ARG voiceId=Arthur


RUN CGO_ENABLED=0 GOOS=linux go build -o /speak
EXPOSE 8080

# Run
CMD ["/speak"]