package mqtt

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttClient interface{}
type mqttClient struct {
	mqttClient mqtt.Client
}

func NewMqttClient() MqttClient {
	var broker = "broker.emqx.io"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername("emqx")
	opts.SetPassword("public")
	// opts.SetDefaultPublishHandler(messagePubHandler)
	// opts.OnConnect = connectHandler
	// opts.OnConnectionLost = connectLostHandler
	return &mqttClient{
		mqttClient: mqtt.NewClient(opts),
	}
}
