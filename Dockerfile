FROM golang

RUN mkdir /appair
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]