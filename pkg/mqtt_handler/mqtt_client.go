package mqtt_handler

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTClient struct {
	client mqtt.Client
}

func NewMQTTClient(broker string, clientID string) *MQTTClient {
	options := mqtt.NewClientOptions()
	options.AddBroker(broker)
	options.SetClientID(clientID)
	options.SetOrderMatters(false)
	options.SetCleanSession(false)
	options.SetAutoReconnect(true)

	client := mqtt.NewClient(options)
	return &MQTTClient{client: client}
}

func (c *MQTTClient) Connect() error {
	token := c.client.Connect()
	token.Wait()
	return token.Error()
}

func (c *MQTTClient) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) error {
	token := c.client.Subscribe(topic, qos, callback)
	token.Wait()
	return token.Error()
}

func (c *MQTTClient) Publish(topic string, qos byte, payload string) error {
	token := c.client.Publish(topic, qos, false, payload)
	token.Wait()
	return token.Error()
}
