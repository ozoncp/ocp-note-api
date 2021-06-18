package main

import (
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/ozoncp/ocp-note-api/internal/config"
	"github.com/ozoncp/ocp-note-api/internal/producer"
	"github.com/rs/zerolog/log"
)

var cfg *config.Config

func subscribe(topic string, consumer sarama.Consumer) error {
	partitionList, err := consumer.Partitions(topic) //get all partitions on the given topic

	if err != nil {
		return err
	}

	initialOffset := sarama.OffsetOldest //get offset for the oldest message on the topic

	for _, partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, partition, initialOffset)

		if err != nil {
			return err
		}

		for message := range pc.Messages() {
			messageReceived(message)
		}
	}

	return nil
}

func messageReceived(message *sarama.ConsumerMessage) {
	var msg producer.Message
	err := json.Unmarshal(message.Value, &msg)

	if err != nil {
		fmt.Printf("Error unmarshalling message: %s\n", err)
	}

	log.Info().Msgf("Message: %v", msg.Body)
}

func main() {

	var err error

	cfg, err = config.Read("../../config.yml")

	if err != nil {
		log.Fatal().Err(err).Msgf("failed to open configuration file")
		return
	}

	consumer, err := sarama.NewConsumer(cfg.Kafka.Brokers, nil)

	if err != nil {
		log.Fatal().Msgf("NewConsumer error: %v", err)
	}

	log.Info().Msgf("awaiting messages from Kafka ...")

	err = subscribe(cfg.Kafka.Topic, consumer)

	if err != nil {
		log.Fatal().Msgf("Subscribe failed: %v", err)
	}
}
