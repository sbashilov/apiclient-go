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
	params := methods.SetScenarioInfoParams{ScenarioId:1, ScenarioName:"scen11", ScenarioScript:"var s="hello world";"}
	res, verr, err := client.Scenarios.SetScenarioInfo(params)
	fmt.Println(res, verr, err)
}