FROM golang:1.9-alpine3.7

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh curl

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64 && chmod +x /usr/local/bin/dep

WORKDIR /go/src/zerocoin
COPY . .

RUN dep ensure -vendor-only

# RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

EXPOSE 1323

CMD ["go-wrapper", "run"] # ["zerocoin"]