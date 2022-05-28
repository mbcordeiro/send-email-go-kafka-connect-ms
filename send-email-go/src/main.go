package main

import (
	"crypto/tls"
	"encoding/json"

	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/mbcordeiro/send-email-go-kafka-connect-ms/send-email-go/src/email"
	"github.com/mbcordeiro/send-email-go-kafka-connect-ms/send-email-go/src/kafka"
	gomail "gopkg.in/mail.v2"
)

func main() {
	var emailChan = make(chan email.Email)
	var msgChan = make(chan *ckafka.Message)

	d := gomail.NewDialer(
		"smtp.mailgun.org",
		587,
		"matheusdebarroscordeiro@gmail.com",
		"f35ab0f611dc1e41090550f86018eb7900d06961",
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	es := email.NewMailSender()
	es.From = "matheusdebarroscordeiro@gmail.com"
	es.Dailer = d
	go es.Send(emailChan)
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9094",
		"client.id":         "emailapp",
		"group.id":          "emailapp",
	}
	topics := []string{"emails"}
	consumer := kafka.NewConsumer(configMap, topics)
	go consumer.Consume(msgChan)
	fmt.Println("Consumer msgs")
	for msg := range msgChan {
		var input email.Email
		json.Unmarshal(msg.Value, &input)
		emailCh <- input
	}
}
