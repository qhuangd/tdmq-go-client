// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package main

import (
	"context"
	"github.com/TencentCloud/tdmq-go-client/pulsar"
	"log"
	"strconv"
)

func main() {
	authParams := make(map[string]string)
	authParams["secretId"] = "AKxxxxxxxxxxCx"
	authParams["secretKey"] = "SDxxxxxxxxxxCb"
	authParams["region"] = "ap-guangzhou"
	authParams["ownerUin"] = "xxxxxxxxxx"
	authParams["uin"] = "xxxxxxxxxx"
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:       "pulsar://localhost:6650",
		AuthCloud: pulsar.NewAuthenticationCloudCam(authParams),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		DisableBatching: true,
		Topic:           "persistent://appid/namespace/topic-1",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	ctx := context.Background()

	for j := 0; j < 10; j++ {
		if msgId, err := producer.Send(ctx, &pulsar.ProducerMessage{
			Payload: []byte("Hello " + strconv.Itoa(j)),
		}); err != nil {
			log.Fatal(err)
		} else {
			log.Println("Published message: ", msgId)
		}
	}

}
