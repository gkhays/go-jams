FROM registry.access.redhat.com/ubi9/ubi-minimal:9.3

RUN microdnf install -y golang

WORKDIR /app

COPY ../server /app/

EXPOSE 8080
CMD ["server"]