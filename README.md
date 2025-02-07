# Just Another Micro Service (JAMS)

![](docs/blue-jam.png)

Just another micro service, written in Go.

## Build

### Prerequisites

- GNU make
- UPX
- Docker

```bash
make release
```

### Docker

```bash
docker build -f docker/Dockerfile -t go-jams .
```

## Run

```bash
docker run --rm -it -p 8080:8080 go-jams
```

Or use the provided shell script, `run.sh`.

```bash
./run.sh
```