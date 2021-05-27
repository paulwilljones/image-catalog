package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/api/iterator"

	asset "cloud.google.com/go/asset/apiv1p5beta1"
	assetpb "google.golang.org/genproto/googleapis/cloud/asset/v1p5beta1"
)

func main() {
	ctx := context.Background()
	client, err := asset.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	assetType := "containerregistry.googleapis.com/Image"
	req := &assetpb.ListAssetsRequest{
		Parent:      fmt.Sprintf("projects/%s", projectID),
		AssetTypes:  []string{assetType},
		ContentType: assetpb.ContentType_RESOURCE,
	}

	// Call ListAssets API to get an asset iterator.
	it := client.ListAssets(ctx, req)

	// Traverse and print the first 10 listed assets in response.
	for i := 0; i < 100; i++ {
		response, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s %s\n", response.Resource.Data.Fields["name"], response.Resource.Data.Fields["tags"])
	}
}
