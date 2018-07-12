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
	"bufio"
	"fmt"
	"io"
	"os"

	m "github.com/BBVA/masquerade/pkg/rabbit"
	"github.com/spf13/cobra"

	"github.com/streadway/amqp"
)

var (
	rabbitDial    string
	rabbitChannel string
)

var rootCmd = &cobra.Command{
	Use:   "maskrabbitout",
	Short: "masquerade rabbit mq export command",
	Run:   maskrabbitoutMain,
}

func maskrabbitoutMain(cmd *cobra.Command, args []string) {
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

	var body []byte

	snr := bufio.NewScanner(os.Stdin)
	for snr.Scan() {
		line := snr.Text()
		if len(line) == 0 {
			break
		}

		body = []byte(line)
		pubErr := ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		if pubErr != nil {
			e := fmt.Sprintf("%v\n", pubErr)
			os.Stderr.WriteString(e)
		}
	}
	if err := snr.Err(); err != nil {
		if err != io.EOF {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func main() {
	rootCmd.Flags().StringVar(&rabbitDial, "dial", "", "Dial config, the rabbit to we write data")
	rootCmd.Flags().StringVar(&rabbitChannel, "channel", "", "Channel to write data")
	rootCmd.MarkFlagRequired("dial")
	rootCmd.MarkFlagRequired("channel")

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
