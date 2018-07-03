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
package main

import (
	"flag"
	"fmt"

	m "masquerade/pkg/rabbit"

	"github.com/streadway/amqp"
)

func main() {
	rabbitDial := flag.String("dial", "", "Dial config")
	rabbitChannel := flag.String("channel", "", "Channel from/to read/write")
	quantity := flag.Int("quantity", -1, "How many msg read, -1 for unlimited, default -1")
	flag.Parse()

	m.FailOnAbsentStringParam(rabbitDial, "Rabbit Dial expected")
	m.FailOnAbsentStringParam(rabbitChannel, "Rabbit channel expected")

	conn, err := amqp.Dial(*rabbitDial)
	if err != nil {
		m.Fail("No connection")
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		m.Fail("No channel")
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		*rabbitChannel, // name
		false,          // durable
		false,          // delete when usused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		m.Fail("No Queue")
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
		m.Fail(fmt.Sprintf("Can't create reader: %v", err))
	}
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("%s\n", string(d.Body))
			if *quantity > 0 {
				*quantity--
			}
			if *quantity == 0 {
				forever <- false
				break
			}
		}
	}()

	_ = <-forever
}
