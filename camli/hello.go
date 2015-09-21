// camli will connect to two clients: a
package main

import (
	"fmt"
	"log"
	"os"

	"camlistore.org/pkg/auth"
	camli "camlistore.org/pkg/client"
	"camlistore.org/pkg/search"
)

func main() {
	fmt.Printf("oi\n")

	camliAuth := auth.NewBasicAuth("nictuku", "SOMETHING")

	client := camli.NewFromParams("http://<IP>:3179", camliAuth)
	client.InsecureTLS = true
	recent, err := client.GetRecentPermanodes(&search.RecentRequest{})
	if err != nil {
		log.Printf("error getting recent permanodes:", err)
		os.Exit(1)
	}
	for _, desc := range recent.Meta {
		fmt.Printf("title: %v\ndescription: %v\n--\n", desc.Title(), desc.Description())
	}
	fmt.Println("printed %d items", len(recent.Meta))
}
