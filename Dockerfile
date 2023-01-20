FROM golang:latest
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o dibagi 
EXPOSE 8080
ENTRYPOINT [ "/app/dibagi" ]

