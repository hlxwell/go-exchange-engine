FROM golang:onbuild

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go build -o orderbook .

ENTRYPOINT ["/app/orderbook"]
