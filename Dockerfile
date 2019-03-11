#FROM golang:latest 
#RUN mkdir /app 
#ADD . /app/ 
#WORKDIR /app 
#RUN go build -o main . 
#CMD ["/app/main"]

FROM golang:onbuild

RUN go get github.com/pilu/fresh

EXPOSE 8080