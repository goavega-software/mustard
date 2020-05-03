package events

import (
	"context"
	"encoding/json"
	"log"
	mustardcore "mustard/core"
	"strings"

	kafka "github.com/segmentio/kafka-go"
)

type EventListener struct {
	URL     string
	Topic   string
	GroupID string
}

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3,
		MaxBytes: 10e5,
	})
}

func (ev EventListener) listen() {
	reader := getKafkaReader(ev.URL, ev.Topic, ev.GroupID)
	defer reader.Close()
	log.Println("starting listening on topic ", ev.Topic)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln("kafka.readMessage failed with error", err)
		}
		eventMessage := mustardcore.EventData{}
		json.Unmarshal(m.Value, &eventMessage)
		mustardcore.GetEventsManager().Notify(eventMessage)
	}
}

func (ev EventListener) Start() {
	if ev.URL == "" || ev.Topic == "" {
		log.Println("No Kafka broker url or topic provided...skipping setup")
		return
	}
	go func() {
		ev.listen()
	}()
}
