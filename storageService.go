package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func ConnectToStorageAccount(connStr string) *azblob.ServiceClient {
	fmt.Println("Connection to storage acc start.")

	serviceClient, err := azblob.NewServiceClientFromConnectionString(connStr, nil)

	if err != nil {
		log.Fatal("Connection failed.", err)
		panic(err)
	}

	fmt.Println("Connection to storage acc end.")
	return serviceClient
}

// TODO: implement
func CreateContainerIfNotExists(containerName string, serviceClient *azblob.ServiceClient) *azblob.ContainerClient {
	fmt.Println("Create Container if not exists start.")

	// First, create a container client, and use the Create method to create a new container in your account
	containerClient, err := serviceClient.NewContainerClient(containerName)
	if err != nil {
		log.Fatal("Create Container if not exists failed: ", err)
		panic(err)
	}

	// All APIs have an options' bag struct as a parameter.
	// The options' bag struct allows you to specify optional parameters such as metadata, public access types, etc.
	// If you want to use the default options, pass in nil.
	_, err = containerClient.Create(context.TODO(), nil)
	if err != nil {
		log.Fatal("Create Container if not exists failed: ", err)
		panic(err)
	}

	fmt.Println("Create Container if not exists end.")
	return containerClient
}

// TODO: implement
func UploadLogs(containerClient *azblob.ContainerClient, logs string) {

	fmt.Println("Upload logs start.")

	// Create a new blockBlobClient from the containerClient
	blockBlobClient, err := containerClient.NewBlockBlobClient("Logs1.json")
	if err != nil {
		log.Fatal("Upload logs failed: ", err)
		panic(err)
	}

	_, err = blockBlobClient.Upload(
		context.TODO(),
		streaming.NopCloser(strings.NewReader(logs)),
		nil)

	if err != nil {
		log.Fatal("Upload logs failed: ", err)
		panic(err)
	}

	fmt.Println("Upload logs End.")
}

//func other() {
// // Download the blob's contents and ensure that the download worked properly
// blobDownloadResponse, err := blockBlobClient.Download(context.TODO(), nil)
// if err != nil {
// 	log.Fatal(err)
// }

// // Use the bytes.Buffer object to read the downloaded data.
// // RetryReaderOptions has a lot of in-depth tuning abilities, but for the sake of simplicity, we'll omit those here.
// reader := blobDownloadResponse.Body(nil)
// downloadData, err := ioutil.ReadAll(reader)
// if err != nil {
// 	log.Fatal(err)
// }
// if string(downloadData) != uploadData {
// 	log.Fatal("Uploaded data should be same as downloaded data")
// }

// err = reader.Close()
// if err != nil {
// 	return
// }

// // ===== 3. List blobs =====
// // List methods returns a pager object which can be used to iterate over the results of a paging operation.
// // To iterate over a page use the NextPage(context.Context) to fetch the next page of results.
// // PageResponse() can be used to iterate over the results of the specific page.
// // Always check the Err() method after paging to see if an error was returned by the pager. A pager will return either an error or the page of results.
// pager := containerClient.ListBlobsFlat(nil)
// for pager.NextPage(context.TODO()) {
// 	resp := pager.PageResponse()
// 	for _, v := range resp.Segment.BlobItems {
// 		fmt.Println(*v.Name)
// 	}
// }

// if err = pager.Err(); err != nil {
// 	log.Fatal(err)
// }

// // Delete the blob.
// _, err = blockBlobClient.Delete(context.TODO(), nil)
// if err != nil {
// 	log.Fatal(err)
// }

// // Delete the container.
// _, err = containerClient.Delete(context.TODO(), nil)
// if err != nil {
// 	log.Fatal(err)
// }
//}

// func getContainerURL(ctx context.Context, accountName, accountGroupName, containerName string) azblob.ContainerURL {
// 	key := getAccountPrimaryKey(ctx, accountName, accountGroupName)
// 	c, _ := azblob.NewSharedKeyCredential(accountName, key)
// 	p := azblob.NewPipeline(c, azblob.PipelineOptions{
// 		Telemetry: azblob.TelemetryOptions{Value: config.UserAgent()},
// 	})
// 	u, _ := url.Parse(fmt.Sprintf(blobFormatString, accountName))
// 	service := azblob.NewServiceURL(*u, p)
// 	container := service.NewContainerURL(containerName)
// 	return container
// }

// func GetContainer(ctx context.Context, accountName, accountGroupName, containerName string) (azblob.ContainerURL, error) {
// 	c := getContainerURL(ctx, accountName, accountGroupName, containerName)

// 	_, err := c.GetProperties(ctx, azblob.LeaseAccessConditions{})
// 	return c, err
// }
