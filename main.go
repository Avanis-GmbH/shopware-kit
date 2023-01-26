package main

import (
	"context"
	"fmt"

	"github.com/Avanis-GmbH/SUC/com"
	"github.com/Avanis-GmbH/SUC/model"
)

func main() {
	ctx := context.Background()
	creds := com.NewIntegrationCredentials("client_id", "client_secret", []string{"write"})
	client, err := com.NewClient(ctx, "", creds, nil)
	if err != nil {
		panic(err)
	}

	categoryCollection := model.CategoryCollection{}
	err = client.Search(com.NewApiContext(ctx), com.Criteria{}, categoryCollection)
	if err != nil {
		panic(err)
	}

	for _, category := range categoryCollection.GetData() {
		fmt.Printf("%+v \n", category)
	}

}
