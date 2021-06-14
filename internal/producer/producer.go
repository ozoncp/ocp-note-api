package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
)

type MessageType = int

type Producer interface {
	Send(msg Message) error
}

const (
	capacity = 256
)

const (
	Created MessageType = iota
	Updated
	Removed
)

func MessageTypeToString(type_ MessageType) string {
	switch type_ {
	case Created:
		return "created"
	case Updated:
		return "updated"
	case Removed:
		return "removed"
	}

	return "unknown event"
}

type Message struct {
	type_ MessageType
	Body  map[string]interface{}
}

var brokerAddress = []string{"127.0.0.1:9092"}

func New(ctx context.Context, topic string) (Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokerAddress, config)

	if err != nil {
		return nil, err
	}

	messageChan := make(chan *sarama.ProducerMessage, capacity)
	newProducer := &dataProducer{producer: producer, topic: topic, messageChan: messageChan}
	go newProducer.handleMessages(ctx)

	return newProducer, nil
}

type dataProducer struct {
	producer    sarama.SyncProducer
	topic       string
	messageChan chan *sarama.ProducerMessage
}

func (dProducer *dataProducer) Send(msg Message) error {
	msgBytes, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	dProducer.messageChan <- &sarama.ProducerMessage{
		Topic:     dProducer.topic,
		Partition: -1,
		Key:       sarama.StringEncoder(dProducer.topic),
		Value:     sarama.StringEncoder(msgBytes),
	}

	return nil
}

func (dProducer *dataProducer) handleMessages(ctx context.Context) {
	for {
		select {
		case msg := <-dProducer.messageChan:
			_, _, err := dProducer.producer.SendMessage(msg)
			if err != nil {
				log.Error().Msgf("failed to send message to kafka: %v", err)
			}
		case <-ctx.Done():
			close(dProducer.messageChan)
			return
		}
	}
}

func CreateMessage(type_ MessageType, noteId uint64, timestamp time.Time) Message {
	return Message{
		type_: type_,
		Body: map[string]interface{}{
			"note_id":   noteId,
			"operation": fmt.Sprintf("%s note", MessageTypeToString(type_)),
			"timestamp": timestamp,
		},
	}
}
