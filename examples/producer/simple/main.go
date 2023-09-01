/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/zhiyin2021/rocketmq-client-go"
	"github.com/zhiyin2021/rocketmq-client-go/primitive"
	"github.com/zhiyin2021/rocketmq-client-go/producer"
)

const (
	endpoint  = "http://MQ_INST_5845280770824374_BYi43TiH.ap-southeast-1.mq.aliyuncs.com:80"
	accesskey = "LTAI5t9WMn6fLwuiF4PRNUgQ"
	secretkey = "QA1LhXVTNzXKX5ORBFDQKO9w16X3IL"
	namespace = "MQ_INST_5845280770824374_BYi43TiH"
	topic     = "GO_TEST"
	groupName = "GID_GO_TEST"
)

// Package main implements a simple producer to send message.
func main() {
	p, _ := rocketmq.NewProducer(
		producer.WithCredentials(primitive.Credentials{
			AccessKey: accesskey,
			SecretKey: secretkey,
		}),
		producer.WithGroupName(groupName),
		producer.WithNamespace(namespace),

		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{endpoint})),
		producer.WithRetry(2),
	)
	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}

	msg := &primitive.Message{
		Topic: topic,
		Body:  []byte("Hello RocketMQ Go Client! "),
	}
	res, err := p.SendSync(context.Background(), msg)

	if err != nil {
		fmt.Printf("send message error: %s\n", err)
	} else {
		fmt.Printf("send message success: result=%s\n", res.String())
	}
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s", err.Error())
	}
}
