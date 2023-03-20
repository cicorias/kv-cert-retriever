package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azkeys"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

func GetCertificate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		// TODO: handle error
	}

	client, err := azkeys.NewClient("https://<TODO: your vault name>.vault.azure.net", cred, nil)
	if err != nil {
		// TODO: handle error
	}

	thekey, err := client.GetKey(context.TODO(), "<TODO: your key name>", "", nil)

	// certificateBundle, _ := client.GetCertificate(context.Background(), "<your-certificate-name>", "<your-certificate-version>")

	// Do something with the certificateBundle (e.g. access the certificate data)
}

func createSecret() {
	keyVaultName := os.Getenv("KEY_VAULT_NAME")
	secretName := "quickstart-secret"
	secretValue := "createdWithGO"
	keyVaultUrl := fmt.Sprintf("https://%s.vault.azure.net/", keyVaultName)

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}

	client, err := azsecrets.NewClient(keyVaultUrl, cred, nil)
	if err != nil {
		log.Fatalf("failed to create a client: %v", err)
	}

	resp, err := client.SetSecret(context.TODO(), secretName, secretValue, nil)
	if err != nil {
		log.Fatalf("failed to create a secret: %v", err)
	}

	fmt.Printf("Name: %s, Value: %s\n", *resp.ID, *resp.Value)
}

func main() {
	createSecret()
}
