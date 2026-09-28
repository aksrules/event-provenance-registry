

package main

import (
 "log"

 "github.com/sassoftware/event-provenance-registry/pkg/message"
 "github.com/sassoftware/event-provenance-registry/pkg/watcher"
)

func main() {
 seeds := []string{"localhost:19092"}
 topics := []string{"epr.dev.events"}
 consumerGroup := "watcher-workshop"

 watcher, err := watcher.New(seeds, topics, consumerGroup)
 if err != nil {
  panic(err)
 }
 defer watcher.Client.Close()

 go watcher.StartTaskHandler(customTaskHandler)

 watcher.ConsumeRecords(customMatcher)
}

func customMatcher(msg *message.Message) bool {
 return msg.Type == "foo.bar"
}

func customTaskHandler(msg *message.Message) error {
 log.Default().Printf("I received a task with value '%v'", msg)
 return nil
}