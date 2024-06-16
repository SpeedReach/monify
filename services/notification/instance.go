package notification

import (
	"context"
	"encoding/json"
	"fmt"
	"monify/lib/group_bill"
)

func Start(config Config) {
	consumers, err := NewConsumers(config)
	if err != nil {
		panic(err)
	}

	for {
		m, err := consumers.GroupBillsConsumer.ReadMessage(context.Background())
		if err != nil {
			panic(err)
		}
		modification := group_bill.GroupBillModification{}
		err = json.Unmarshal(m.Value, &modification)
		if err != nil {
			panic(err)
		}

		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n",
			m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
