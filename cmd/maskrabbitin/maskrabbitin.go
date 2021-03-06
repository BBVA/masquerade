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
	"fmt"
	"os"

	"github.com/spf13/cobra"

	m "github.com/BBVA/masquerade/pkg/rabbit"

	"github.com/streadway/amqp"
)

var (
	rabbitDial    string
	rabbitChannel string
	quantity      int
)

var rootCmd = &cobra.Command{
	Use:   "maskrabbitin",
	Short: "masquerade rabbit mq import command",
	Run:   maskrabbitinMain,
}

func maskrabbitinMain(cmd *cobra.Command, args []string) {
	conn, err := amqp.Dial(rabbitDial)
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
		rabbitChannel, // name
		false,         // durable
		false,         // delete when usused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
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
			if quantity > 0 {
				quantity--
			}
			if quantity == 0 {
				forever <- false
				break
			}
		}
	}()

	_ = <-forever
}

func main() {
	rootCmd.Flags().StringVar(&rabbitDial, "dial", "", "Dial config, the rabbit from we read data")
	rootCmd.Flags().StringVar(&rabbitChannel, "channel", "", "Channel to read data")
	rootCmd.Flags().IntVar(&quantity, "quantity", -1, "how many read after command kill itself")

	rootCmd.MarkFlagRequired("dial")
	rootCmd.MarkFlagRequired("channel")
	rootCmd.Flags().MarkHidden("quantity")
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}
