FROM golang:1.15.3 AS build-env

WORKDIR /app/

ADD . .

RUN go build -o executable -mod vendor ./...


FROM gcr.io/distroless/base-debian10

WORKDIR /app/

COPY --from=build-env /app/executable .
COPY --from=build-env /app/employees.db .

EXPOSE 8080

CMD ["/app/executable"]