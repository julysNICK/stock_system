# FROM golang:1.20.3-alpine3.17 AS builder
# WORKDIR /app
# COPY . .
# RUN go build -o main main.go
# RUN apk add curl
# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# FROM alpine:3.17
# WORKDIR /app
# COPY --from=builder /app/main .
# COPY --from=builder /app/migrate ./migrate
# COPY app.env .
# COPY start.sh .
# COPY wait-for.sh .
# RUN chmod +x wait-for.sh
# RUN chmod +x start.sh
# COPY db/migration ./migration
# # COPY cron-script.sh .
# # RUN chmod +x cron-script.sh

# EXPOSE 8080
# CMD ["/app/main"]
# ENTRYPOINT [ "/app/start.sh" ]

FROM golang:1.20.3-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz


#Run stage

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
RUN chmod +x wait-for.sh
RUN chmod +x start.sh
COPY db/migration ./migration


EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]