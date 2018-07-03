// Copyright 2018 BBVA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package rabbit

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/streadway/amqp"
)

//Fail do a fail
func Fail(msg string) {
	os.Stderr.WriteString(msg)
	flag.PrintDefaults()
	os.Exit(1)
}

// FailOnAbsentStringParam checks param and fail if not found
func FailOnAbsentStringParam(param *string, msg string) {
	if *param == "" {
		formattedMessage := fmt.Sprintf("%s\n\n", msg)
		Fail(formattedMessage)
	}
}

// WriteOnChannel Writes several lines to a channel
func WriteOnChannel(dial, channel string, lines []string) error {
	conn, err := amqp.Dial(dial)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		channel, // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return err
	}

	var body []byte

	for _, msg := range lines {
		body = []byte(msg)
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		if err != nil {
			return err
		}
	}

	return nil
}

// ReadFromChannel will read certain ammount of lines from a channel
func ReadFromChannel(dial, channel string, linesToRead int) ([]string, error) {
	conn, err := amqp.Dial(dial)
	if err != nil {
		return []string{}, err
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		return []string{}, err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		channel, // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return []string{}, err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return []string{}, err
	}
	linesReaded := make(chan []string)

	go func() {
		n := 0
		l := []string{}
		for d := range msgs {
			l = append(l, string(d.Body))
			n++
			if n >= linesToRead {
				break
			}
		}

		linesReaded <- l
	}()

	select {
	case res := <-linesReaded:
		return res, nil
	case <-time.After(3 * time.Second):
		return []string{}, errors.New("timeout on channel read")
	}
}
