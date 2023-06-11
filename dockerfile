FROM golang:latest

WORKDIR /
COPY ./employee-app .

CMD ./employee-app