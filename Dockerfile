FROM golang:1.15 as build
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build main /
CMD ["/main"]