package main

import (
	"fmt"
	"os"

	"github.com/CenturyLinkLabs/docker-reg-client/registry"
)

var cli *registry.Client
var basic *registry.BasicAuth
var readToken *registry.TokenAuth
var writeToken *registry.TokenAuth

func search() {
	results, err := cli.Search.Query("sequenceiq", 1, 25)
	if err != nil {
		panic(err)
	}

	for _, result := range results.Results {
		fmt.Println(result.Name)
	}
}

func init() {
	cli = registry.NewClient()
	user := os.Getenv("DOCKER_HUB_USER")
	password := os.Getenv("DOCKER_HUB_PASSWORD")

	if user == "" {
		panic("DOCKER_HUB_USER env var is required")
	}

	if password == "" {
		panic("DOCKER_HUB_PASSWORD env var is required")
	}

	basic = &registry.BasicAuth{user, password}
	var err error
	readToken, err = cli.Hub.GetReadTokenWithAuth("sequenceiq/cloudbreak", basic)
	writeToken, err = cli.Hub.GetWriteToken("sequenceiq/cloudbreak", basic)
	if err != nil {
		panic(err)
	}

	fmt.Println("readToken", readToken)
	fmt.Println("writeToken", writeToken)
}

func listTags() {
	tm, err := cli.Repository.ListTags("sequenceiq/cloudbreak", readToken)
	if err != nil {
		panic(err)
	}

	for k, v := range tm {
		fmt.Println(k, v)
	}
}

func metaData() {
	meta, err := cli.Image.GetMetadata("6a2875554bc76e4203fe6ce193c88f38b93c0258655fb4c905c70b952fbb3775", basic)
	if err != nil {
		panic(err)
	}
	fmt.Println("meta:", meta)
}
func main() {
	//search()
	listTags()
}
