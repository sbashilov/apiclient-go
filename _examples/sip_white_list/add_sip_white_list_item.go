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
	params := methods.AddSipWhiteListItemParams{SipWhitelistNetwork:"192.168.1.5/16"}
	res, verr, err := client.SipWhiteList.AddSipWhiteListItem(params)
	fmt.Println(res, verr, err)
}