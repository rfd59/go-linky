 # Go-Linky

![GitHub Release](https://img.shields.io/github/v/release/rfd59/go-linky)
![GitHub Issues](https://img.shields.io/github/issues/rfd59/go-linky)
![GitHub Pull Requests](https://img.shields.io/github/issues-pr/rfd59/go-linky)
![GitHub License](https://img.shields.io/github/license/rfd59/go-linky)

![Go version](https://img.shields.io/github/go-mod/go-version/rfd59/go-linky)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/rfd59/go-linky/.github%2Fworkflows%2Fci.yml)
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
