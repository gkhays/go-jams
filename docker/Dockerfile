FROM registry.access.redhat.com/ubi9/ubi-minimal:9.3

RUN microdnf install -y golang

WORKDIR /app

COPY bin/server-0.1.0-dev-linux-amd64 /app/server

EXPOSE 8080
CMD ["/app/server"]