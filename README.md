# Just Another Micro Service (JAMS)

![](docs/blue-jam.png)

Just another micro service, written in Go.

## Build

### Prerequisites

- GNU make
- UPX
- Docker
- [mkcert](https://github.com/FiloSottile/mkcert)

Create a CA and certificates.

```bash
mkcert -install
mkcert jams.com "*.jams.com" jams.test localhost 127.0.0.1 ::1
```

Create a release build with all targets.

```bash
make release
```

Produce a Docker container image.

```bash
make docker
```

### Docker

To build a Docker container image manually.

```bash
docker build -f docker/Dockerfile -t go-jams:test .
```

## Run

```bash
docker run --rm -it -p 8080:8080 go-jams:test
```

Or use the provided shell script, `run.sh`.

```bash
./run.sh
```