 # Go-Linky  

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
