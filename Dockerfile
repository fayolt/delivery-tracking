FROM golang:1.15.5-alpine
 
RUN mkdir -p /app
 
WORKDIR /app
 
ADD . /app

RUN go build .
 
ENTRYPOINT [ "./delivery-tracking" ]