package main

import (
	"fmt"
	"github.com/adamhoof/InternetOfToys/pkg/config"
	databaseInterface "github.com/adamhoof/InternetOfToys/pkg/database"
	database "github.com/adamhoof/InternetOfToys/pkg/database/implementations"
	"github.com/adamhoof/InternetOfToys/pkg/mqtt_handler"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

func main() {
	// Load the configuration
	conf, err := config.LoadConfig("/home/adamhoof/Projects/InternetOfToys/Config.json")
	if err != nil {
		log.Fatalf("Failed to load conf: %v", err)
	}

	// Connect to the databaseInterface
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Dbname)
	db, err := database.NewPostgres(connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to databaseInterface: %v", err)
	}

	// Connect to the MQTT broker
	mqttClient := mqtt_handler.NewMQTTClient(conf.MQTT.ServerAndPort, "middleware")

	if err := mqttClient.Connect(); err != nil {
		log.Fatalf("Failed to connect to MQTT broker: %v", err)
	}

	// Subscribe to the /boot topic
	if err = mqttClient.Subscribe(mqtt_handler.TopicBoot, 2, bootHandler(db)); err != nil {
		log.Fatalf("failed to subscribe: %s", err)
	}

	// Subscribe to the /command_reply topic
	if err = mqttClient.Subscribe(mqtt_handler.TopicCommandReply, 0, commandReplyHandler(db)); err != nil {
		log.Fatalf("failed to subscribe: %s", err)
	}

	// Start the web server or bot here
}

func bootHandler(db databaseInterface.Database) mqtt.MessageHandler {
	return func(client mqtt.Client, msg mqtt.Message) {
		// Handle boot messages here

	}
}

func commandReplyHandler(db databaseInterface.Database) mqtt.MessageHandler {
	return func(client mqtt.Client, msg mqtt.Message) {
		// Handle command reply messages here
	}
}

func sendCommand(deviceMacAddress string, command string) {
	// Publish a command to the /command/[mac_address] topic
}
