package main

import (
	"fmt"
	"github.com/voximplant/apiclient-go/config"
	"github.com/voximplant/apiclient-go/methods"
)

func main() {
	voxConfig := config.NewConfig().WithKeyPath("vox_key_jwt.json")
	client, err := methods.NewClient(voxConfig)
	if err != nil {
		panic(err)
	}
	params := methods.GetACDQueueStatisticsParams{FromDate:2017-01-01 00:00:00}
	res, verr, err := client.Queues.GetACDQueueStatistics(params)
	fmt.Println(res, verr, err)
}