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
