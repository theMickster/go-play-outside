package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

type DbClient struct {
	client *azcosmos.Client
	config *DbConfig
}

type DbConfig struct {
	DdName        string
	DbUrl         string
	Container     string
	ClientOptions *azidentity.DefaultAzureCredentialOptions
}

func NewDatabaseConfig() *DbConfig {
	opt := &azidentity.DefaultAzureCredentialOptions{}
	c := &DbConfig{
		DdName:        os.Getenv("TailwindCosmosDatabase"),
		DbUrl:         os.Getenv("TailwindCosmosEndpoint"),
		Container:     os.Getenv("TailwindCosmosContainer"),
		ClientOptions: opt,
	}
	return c
}

func NewDatabaseClient(config *DbConfig) (*DbClient, error) {
	cred, err := azidentity.NewDefaultAzureCredential(config.ClientOptions)
	if err != nil {
		return nil, err
	}

	d := &DbClient{
		config: config,
	}

	client, err := azcosmos.NewClient(d.config.DbUrl, cred, nil)
	if err != nil {
		return nil, err
	}

	d.client = client
	return d, nil
}

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type ProductItem struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	StockUnits  string `json:"stockUnits"`
	ReleaseDate string `json:"releaseDate"`
	Description string `json:"description"`
	StarRating  int64  `json:"starRating"`
	ImageUrl    string `json:"imageUrl"`
}

func main() {

	myCtx := context.TODO()
	var (
		cosmosEndpoint  = os.Getenv("TailwindCosmosEndpoint")
		cosmosKey       = os.Getenv("TailwindCosmosKey")
		cosmosDatabase  = os.Getenv("TailwindCosmosDatabase")
		cosmosContainer = os.Getenv("TailwindCosmosContainer")
	)

	subscriptionId, keyExists := os.LookupEnv("AzureDevTestSubscriptionId")
	if !keyExists {
		handle(errors.New("unable to locate AzureDevTestSubscriptionId from environment variables"))
	}
	fmt.Printf("Tailwind Products via Cosmos Db in the Azure Subscription ==> %s", subscriptionId)

	// only necessary when debugging the fetching of connection string information.
	//fmt.Printf("cosmosEndpoint - %s :: cosmosKey - %s :: cosmosDatabase - %s :: cosmosContainer %s \n", cosmosEndpoint, cosmosKey, cosmosDatabase, cosmosContainer)

	// when using azure token generation over the cosmos db key
	//credential, err := azidentity.NewDefaultAzureCredential(nil)

	credential, err := azcosmos.NewKeyCredential(cosmosKey)
	handle(err)

	client, err := azcosmos.NewClientWithKey(cosmosEndpoint, credential, nil)
	handle(err)

	database, err := client.NewDatabase(cosmosDatabase)
	handle(err)

	resultDb, err := database.Read(myCtx, nil)
	handle(err)
	fmt.Printf("Database Resource Id  %s \n", resultDb.DatabaseProperties.ResourceID)

	container, err := client.NewContainer(cosmosDatabase, cosmosContainer)
	handle(err)

	partitionKey := azcosmos.NewPartitionKeyString("Porter-Cable")

	id := "00236498-f266-40bf-870d-96876c5008e4"
	itemResponse, err := container.ReadItem(myCtx, partitionKey, id, nil)
	if err != nil {
		var responseErr *azcore.ResponseError
		errors.As(err, &responseErr)
		handle(responseErr)
	}

	if itemResponse.RawResponse.StatusCode == 200 {
		item := ProductItem{}
		err = json.Unmarshal(itemResponse.Value, &item)
		handle(err)

		fmt.Printf("Product Item Read Successfully...  Id - %s :: Name - %s :: Consuming %v RU", item.Id, item.Name, itemResponse.RequestCharge)
	}
}
