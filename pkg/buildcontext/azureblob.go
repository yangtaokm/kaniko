/*
Copyright 2018 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package buildcontext

import (
	"context"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/GoogleContainerTools/kaniko/pkg/constants"
	"github.com/GoogleContainerTools/kaniko/pkg/util"
)

// GCS struct for Google Cloud Storage processing
type AzureBlob struct {
	context string
}

func (b *AzureBlob) UnpackTarFromBuildContext() (string, error) {

	
	u, _ := url.Parse(b.context)
	parts := azblob.NewBlobURLParts(*u)
	accountName := strings.Split(parts.Host, ".")[0]

	accountKey := os.Getenv("AZURE_STORAGE_ACCESS_KEY")

	if (len(accountKey) == 0 ){
		log.Fatal("AZURE_STORAGE_ACCESS_KEY environment variable is not set")
	}

	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		return accountName, err
	}

	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})
	blobURL :=azblob.NewBlobURL(*u, p)

	directory := constants.BuildContextDir
	//directory :="/home/tao/go/src/github.com/GoogleContainerTools/kaniko/out"
	tarPath := filepath.Join(directory, constants.ContextTar)

	if err := os.MkdirAll(directory, 0750); err != nil {
		return directory, err
	}

	f, err := os.Create(tarPath)
	if err != nil {
		return directory, err
	}

	
	ctx := context.Background()
	if err := azblob.DownloadBlobToFile(ctx, blobURL, 0, 0, f, azblob.DownloadFromBlobOptions{}); err != nil {
		//if stgErr, ok := err.(azblob.StorageError); ok {
			//println("blob err", stgErr.ServiceCode())
			//return directory, err
		//}else {
			return directory,err
		//}
		

	}

	return directory, util.UnpackCompressedTar(tarPath, directory)
}
