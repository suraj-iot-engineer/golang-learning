# 12. IoT Backend Projects

Go is excellent for IoT backends due to its high concurrency and low latency. This module covers MQTT integration and sensor data ingestion.

## 📚 Topics Covered

1.  **MQTT Protocol**: The standard for IoT messaging.
2.  **Paho MQTT Client**: Using `github.com/eclipse/paho.mqtt.golang`.
3.  **Real-time Data Ingestion**: Handling high-frequency sensor streams.
4.  **Pub/Sub Pattern**: Decoupling devices from the backend.

## 🛠️ Prerequisites

- An MQTT Broker (e.g., Mosquitto) running locally or in the cloud.
- Public test broker: `tcp://broker.hivemq.com:1883` (Used in example)

## 🚀 Example
The example simulates a temperature sensor (Publisher) and a backend service (Subscriber).

```bash
go run main.go
```
