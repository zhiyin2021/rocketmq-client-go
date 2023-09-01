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
	"log"
	"os"
	"time"

	"github.com/zhiyin2021/rocketmq-client-go"
	"github.com/zhiyin2021/rocketmq-client-go/consumer"
	"github.com/zhiyin2021/rocketmq-client-go/primitive"
	"github.com/zhiyin2021/rocketmq-client-go/rlog"
)

const (
	endpoint  = "http://MQ_INST_5845280770824374_BYi43TiH.ap-southeast-1.mq.aliyuncs.com:80"
	accesskey = "LTAI5t9WMn6fLwuiF4PRNUgQ"
	secretkey = "QA1LhXVTNzXKX5ORBFDQKO9w16X3IL"
	namespace = "MQ_INST_5845280770824374_BYi43TiH"
	topic     = "GO_TEST"
	groupName = "GID_GO_TEST"
)

func main() {
	sig := make(chan os.Signal)
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithCredentials(primitive.Credentials{
			AccessKey: accesskey,
			SecretKey: secretkey,
		}),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName(groupName),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{endpoint})),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithNamespace(namespace),
		consumer.WithTrace(&primitive.TraceConfig{
			Access: primitive.Cloud,
			// GroupName:    groupName,
			NamesrvAddrs: []string{endpoint},
			Credentials: primitive.Credentials{
				AccessKey: accesskey,
				SecretKey: secretkey,
			},
		}),
	)

	rlog.SetLogLevel("warn")
	err := c.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			fmt.Printf("#############subscribe callback: %v \n", msgs[i])
			time.Sleep(time.Millisecond * 50)
		}

		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	// Note: start after subscribe
	err = c.Start()
	log.Println("consumer start success")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	<-sig
	err = c.Shutdown()
	if err != nil {
		fmt.Printf("shutdown Consumer error: %s", err.Error())
	}
}
