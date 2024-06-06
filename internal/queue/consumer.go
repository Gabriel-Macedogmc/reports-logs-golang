package queue

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/dto"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/reports"
	"github.com/IBM/sarama"
	"github.com/alecthomas/kingpin"
)

var (
	messageCountStart = kingpin.Flag("messageCountStart", "Message counter start from:").Int()
)

func Consumer(interrupt chan os.Signal) {
	kingpin.Parse()
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	brokers := *brokerList

	fmt.Println(brokers)

	master, err := sarama.NewConsumer(brokers, config)

	if err != nil {
		log.Panic(err)
	}

	defer func() {
		if err := master.Close(); err != nil {
			log.Panic(err)
		}
	}()

	consumer, err := master.ConsumePartition(*topic, 0, sarama.OffsetOldest)

	if err != nil {
		log.Panic(err)
	}

	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				log.Println("AQUI")
				log.Println(err.Err.Error())

			case msg := <-consumer.Messages():
				*messageCountStart++
				log.Println("Received messages", string(msg.Key), string(msg.Value))
				var data []dto.ReportDTO

				err := json.Unmarshal(msg.Value, &data)
				if err != nil {
					log.Println(err)
				}

				reports.GenerateDailyReport(data)
			case <-interrupt:
				log.Println("Interrupt is detected")
				doneCh <- struct{}{}
				return
			}
		}
	}()
	<-doneCh
	log.Println("Processed", *messageCountStart, "messages")
}
