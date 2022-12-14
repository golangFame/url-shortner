FROM golang:1.19.4-alpine
COPY . /opt/url-shortner
WORKDIR /opt/url-shortner


RUN go get
RUN go build -o /app/url-shortner ./cmd/...
#COPY server.crt /opt/url-shortner
RUN rm -rf /opt/url-shortner

ENTRYPOINT ["/app/url-shortner"]