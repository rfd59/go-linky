 # Go-Linky

![GitHub Release](https://img.shields.io/github/v/release/rfd59/go-linky)
![GitHub Issues](https://img.shields.io/github/issues/rfd59/go-linky)
![GitHub Pull Requests](https://img.shields.io/github/issues-pr/rfd59/go-linky)
![GitHub License](https://img.shields.io/github/license/rfd59/go-linky)

![Go version](https://img.shields.io/github/go-mod/go-version/rfd59/go-linky)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/rfd59/go-linky/.github%2Fworkflows%2Fbuild.yml)
[![Coverage](https://codecov.io/gh/rfd59/go-linky/graph/badge.svg?token=SL969PMZ04)](https://codecov.io/gh/rfd59/go-linky)

**Go-Linky** is a lightweight Go application that collects real-time telemetry from the French **Linky smart meter** through its serial interface (TIC port). It parses the meter data and publishes it to an **MQTT broker**, making it easy to integrate electricity consumption and production metrics into your home automation setup.  

With Go-Linky, you can:  
- Monitor live power consumption directly from your Linky meter  
- Forward data to MQTT topics for use in **Home Assistant**, **Node-RED**, **Grafana**, or any other MQTT-compatible system  
- Simplify energy tracking and analysis without relying on third-party services  

---

## ✨ Features  

- 📡 **Serial interface support** – reads data directly from the Linky TIC port  
- 🔄 **Historique & Standard modes** – compatible with both data output modes of the Linky meter  
- 🪶 **Lightweight & fast** – written in Go, with minimal resource usage  
- ⚙️ **Easy configuration** – simple environment variables for serial port & MQTT settings  
- 🧩 **MQTT integration** – publishes structured metrics ready to consume in your automation tools  
- 📊 **Real-time energy monitoring** – ideal for dashboards and consumption analysis  

---

## 📗 Documentation

Visit [Wiki](https://github.com/rfd59/go-linky/wiki) pages for the full **Go-Linky** documentation.

---

## ⌨️ Development

### Requirements

- [Go](https://golang.org/doc/install) >= 1.24
- [Golangci-lint](https://golangci-lint.run/docs/welcome/install/) >= 2.4.0
- [Mosquitto](https://mosquitto.org/) >= 2.0.0
  > Can be installed into a [Docker container](https://hub.docker.com/_/eclipse-mosquitto)

### Installation

- Clone the repository and install dependencies
  ```bash
  git clone https://github.com/rfd59/go-linky.git
  cd go-linky
  go mod tidy
  ```
- Build the project
  ```bash
  go build -o ./dist/ ./...
  ```

  > Two binaries will be generated:
  > - **go-linky**: The main program.
  > - **linky-tic**: Use to simulate a Linky meter. TIC frames are sent to a serial port. See [Tests](#local-test) section.

### Tests
#### Unit Test

- Launch the unit tests and the code coverage
  ```bash
  go test -cover ./cmd/...
  ```

#### Local Test

- Launch a MQTT Broker
  > With a Docker container: `docker run -it --name mosquitto --rm -v ${PWD}/test/mosquitto:/mosquitto/config -p 1883:1883 eclipse-mosquitto:2`
- Create virtual serial ports
  ```bash
  sudo apt-get install -y socat
  socat -dd pty,rawer,echo=0,link=/tmp/ttyV0 pty,rawer,echo=0,link=/tmp/ttyV1
  ```
- Generate _Linky_ frames
  ```bash
  ./dist/linky-tic
  ```
- Launch **Go-Linky**
  ```bash
  GOLINKY_LINKY_SERIAL_PORT=/tmp/ttyV1 GOLINKY_DEBUG=log ./dist/go-linky
  ```
- With [MQTT Explorer](https://mqtt-explorer.com/), you can check the topic **linky/123456789000** and see the data.
---

## ⌨️ Contributing

Contributions are welcome!

Please open an issue first to discuss what you would like to change.

Guidelines:
- Format code with gofmt
  > `gofmt -w .`
- Lint with golangci-lint
  > `golangci-lint run`
- Write unit tests for new features

Repository folders:
```txt
go-linky/
├── build/                # Packaging and CI/CD related
│   └── Dockerfile
├── cmd/                  # Main applications
│   └── go-linky/         # Go-Linky entry point
│       ├── main.go
│       ├── core/         # Core business logic
│       ├── infra/        # Infrastructure layer
│       ├── models/       # Domain models / DTOs
│       ├── services/     # Application services / use cases
│       └── utils/        # Helpers and utility functions
├── dist/                 # Build output
├── test/                 # Additional external test apps/data
│   ├── cmd/              # Test applications
│   │   └── linky-tic/    # Linky-TIC entry point (main.go)
│   ├── mock/             # Mock definitions
│   └── mosquitto/        # Mosquitto configuration
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

---

## 📜 License  

This project is licensed under the GPL-3.0 License – see the [LICENSE](./LICENSE) file for details.  