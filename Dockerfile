FROM golang:1.18
WORKDIR /app
COPY ./Backend/ ./
RUN go mod download && go mod verify
RUN go build -o /todo-app
EXPOSE 8888
CMD ["/todo-app"]