package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var broker = "tcp://broker.hivemq.com:1883"
var topic = "golang-learning/sensor/temp"

// messageHandler handles incoming messages
var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Backend received: Topic[%s] Payload[%s]\n", msg.Topic(), msg.Payload())
}

func main() {
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID("go_iot_backend")
	opts.SetDefaultPublishHandler(messageHandler)

	// 1. Connect to Broker
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	fmt.Println("✅ Connected to MQTT Broker")

	// 2. Subscribe to Topic
	if token := client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	fmt.Printf("📡 Subscribed to %s\n", topic)

	// 3. Simulate Sensor (Publisher) in a goroutine
	go func() {
		for {
			temp := 20.0 + rand.Float64()*10.0
			payload := fmt.Sprintf(`{"temp": %.2f, "ts": %d}`, temp, time.Now().Unix())

			token := client.Publish(topic, 0, false, payload)
			token.Wait()
			fmt.Printf("device -> sent: %s\n", payload)

			time.Sleep(2 * time.Second)
		}
	}()

	// Keep running
	select {}
}
