FROM golang:latest 
RUN go get github.com/nlopes/slack
RUN go get github.com/skratchdot/open-golang/open
RUN go get github.com/boltdb/bolt
ENV SLACK_TOKEN='TOKEN'
RUN mkdir /app 
RUN mkdir /db
ADD . /app/ 
WORKDIR /app 
RUN go build -o main . 
CMD ["/app/main"]