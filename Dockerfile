# Build stage
FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN rm -f ./database/db.sql
RUN rm -f ./docker-compose.yml
RUN rm -f ./Dockerfile
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /server

# Final stage
FROM alpine:latest
WORKDIR /
ENV TWILIO_TOKEN=example
ENV TWILIO_SID=example
ENV TWILIO_PH_NO=example
ENV MYSQL_DATABASE=app
ENV MYSQL_USER=xyz
ENV MYSQL_PASSWORD=wordpass
ENV MYSQL_ROOT_PASSWORD=password
COPY --from=builder /server /server
EXPOSE 8080
CMD [ "/server" ]
