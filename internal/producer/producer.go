package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
)

type Producer interface {
	Send(message Message) error
}

var brockerAddress = []string{"127.0.0.1:9092"}

type producer struct {
	dataProducer sarama.SyncProducer
	topic        string
	messageChan  chan *sarama.ProducerMessage
}

type MessageType int

const (
	Create MessageType = iota
	Update
	Remove
)

type Message struct {
	type_ MessageType
	Body  map[string]interface{}
}

func New(ctx context.Context, topic string) (Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer_, err := sarama.NewSyncProducer(brockerAddress, config)

	if err != nil {
		log.Error().Err(err).Msg("failed to create a producer")
		return nil, err
	}

	newProducer := &producer{
		dataProducer: producer_,
		topic:        topic,
		messageChan:  make(chan *sarama.ProducerMessage),
	}

	go newProducer.handleMessage(ctx)

	return newProducer, nil
}

func (dProducer *producer) handleMessage(ctx context.Context) {
	select {
	case msg := <-dProducer.messageChan:
		dProducer.dataProducer.SendMessage(msg)
	case <-ctx.Done():
		close(dProducer.messageChan)
		return
	}
}

func (dProducer *producer) Send(message Message) error {

	dataBytes, err := json.Marshal(message)

	if err != nil {
		log.Error().Err(err).Msg("failed to marshal message to json")
		return err
	}

	dProducer.messageChan <- &sarama.ProducerMessage{
		Topic:     dProducer.topic,
		Key:       sarama.StringEncoder(dProducer.topic),
		Value:     sarama.StringEncoder(dataBytes),
		Partition: -1,
		Timestamp: time.Time{},
	}

	return nil
}

func CreateMessage(type_ MessageType, id uint64, timestamp time.Time) Message {
	return Message{
		type_: type_,
		Body: map[string]interface{}{
			"Id":        id,
			"Operation": fmt.Sprintf("%s note", convertMessageTypeToString(type_)),
			"Timestamp": timestamp,
		},
	}
}

func convertMessageTypeToString(type_ MessageType) string {
	switch type_ {
	case Create:
		return "Created"
	case Update:
		return "Updated"
	case Remove:
		return "Removed"
	}

	return "unknown message type"
}
