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
	params := methods.AddPstnBlackListItemParams{PstnBlacklistPhone:"123456789"}
	res, verr, err := client.PstnBlacklist.AddPstnBlackListItem(params)
	fmt.Println(res, verr, err)
}