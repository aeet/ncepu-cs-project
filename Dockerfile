FROM golang:1.20 as biulder

WORKDIR /build
COPY . /build/
RUN go build -o ncepu-cs-project


FROM golang:1.20 as runner
WORKDIR /app
COPY --from=biulder /build/ncepu-cs-project /app/
COPY static /app/static
COPY config.yaml /app/config.yaml

EXPOSE 8080
CMD ["./ncepu-cs-project"]