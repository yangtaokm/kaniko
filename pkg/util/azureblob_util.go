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

package util

import (
	"strings"

	"github.com/GoogleContainerTools/kaniko/pkg/constants"
)

//return storage accoutName, containerName and blobName from context string
//format of context [accountName]/[containerName]/[PathToContextfile]

func GetContainerAndBlob(context string) (accoutName string, containerName string, blobName string) {
	split := strings.SplitN(context, "/", 3)
	accountName := strings.Split(split[0], ".")[0]
	containerName = split[1]
	if len(split) == 3 && split[2] != "" {
		blobName = split[2]
	} else {
		blobName = constants.ContextTar
	}

	return accountName, containerName, blobName
}
