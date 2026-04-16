package main

import (
	"encoding/json"
	"time"
)

type Oracle struct {
	Provider string
	Data     map[string]interface{}
}

func NewOracle(provider string) *Oracle {
	return &Oracle{
		Provider: provider,
		Data:     make(map[string]interface{}),
	}
}

func (o *Oracle) FetchExternal(url string) {
	time.Sleep(1 * time.Second)
	o.Data["source"] = url
	o.Data["timestamp"] = time.Now().Unix()
	o.Data["value"] = 100
}

func (o *Oracle) PackData() []byte {
	b, _ := json.Marshal(o.Data)
	return b
}

func (o *Oracle) Verify() bool {
	return o.Data["value"] != nil
}
