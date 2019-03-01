FROM golang:latest 

COPY . /go/src/simple-rest

WORKDIR /go/src/simple-rest
COPY ./main.go .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go", "run", "main.go"]
# ENTRYPOINT ["app", "-f=7", "-s=9"]

# RUN mkdir /app 
# ADD . /app/ 
# WORKDIR /app 
# RUN go build -o main . 
# CMD ["/app/main"]
