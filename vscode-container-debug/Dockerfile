FROM golang:1.14.12-stretch
RUN apt-get update && \
    apt-get install -y exiftool
COPY main.go ./
COPY in_data/gobook.pdf ./in_data/gobook.pdf
ENTRYPOINT ["go", "run", "main.go"]
