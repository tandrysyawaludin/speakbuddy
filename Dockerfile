FROM golang:1.21

RUN apt-get update && apt-get install -y ffmpeg && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY pkg ./pkg
COPY cli ./cli
COPY go.mod ./
RUN go mod tidy -e

RUN cd cli && go build -o speakbuddybebinary

CMD ["/app/cli/speakbuddybebinary"]