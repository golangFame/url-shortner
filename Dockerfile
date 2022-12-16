FROM golang:1.19.4-alpine
COPY . /opt/url-shortner
WORKDIR /opt/url-shortner


ENV PORT=80

RUN go get
RUN go build -o /app/url-shortner ./cmd/...
#COPY server.crt /opt/url-shortner
RUN rm -rf /opt/url-shortner

EXPOSE 80

ENTRYPOINT ["/app/url-shortner"]