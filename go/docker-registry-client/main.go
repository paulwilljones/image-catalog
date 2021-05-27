package main

import (
	"fmt"

	"github.com/heroku/docker-registry-client/registry"
)

func main() {
	url := "https://registry-1.docker.io/"
	username := "" // anonymous
	password := "" // anonymous
	hub, _ := registry.New(url, username, password)

	//repositories, _ := hub.Repositories()
	//fmt.Println(repositories)

	tags, _ := hub.Tags("library/nginx")
	//fmt.Println(tags)
	for _, tag := range tags {
		digest, _ := hub.ManifestDigest("library/nginx", tag)
		//fmt.Println(digest)
		fmt.Printf("%s %s", tag, digest)
	}
}
