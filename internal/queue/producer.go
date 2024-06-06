package queue

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/dto"
	"github.com/IBM/sarama"
	"github.com/alecthomas/kingpin"
)

var (
	brokerList = kingpin.Flag("brokerList", "List of brokers to connect").Default("localhost:9094").Strings()
	topic      = kingpin.Flag("topic", "Topic name").Default("important").String()
	maxRetry   = kingpin.Flag("maxRetry", "Retry limit").Default("5").Int()
)

func Producer(data []dto.ReportDTO) {
	kingpin.Parse()
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = *maxRetry
	config.Producer.Return.Successes = true

	fmt.Println(topic)

	producer, err := sarama.NewSyncProducer(*brokerList, config)

	if err != nil {
		log.Panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Panic(err)
		}
	}()

	encoded, err := json.Marshal(data)

	if err != nil {
		log.Panic(err)
	}

	msg := &sarama.ProducerMessage{
		Topic: *topic,
		Value: sarama.StringEncoder(encoded),
	}

	partition, offset, err := producer.SendMessage(msg)

	if err != nil {
		log.Panic(err)
	}

	log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", *topic, partition, offset)
}
