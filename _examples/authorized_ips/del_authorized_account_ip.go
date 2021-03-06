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
	params := methods.DelAuthorizedAccountIPParams{AuthorizedIp:"92.255.220.0/24"}
	res, verr, err := client.AuthorizedIps.DelAuthorizedAccountIP(params)
	fmt.Println(res, verr, err)
}