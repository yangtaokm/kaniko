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
	"testing"

	"github.com/GoogleContainerTools/kaniko/pkg/constants"
	"github.com/GoogleContainerTools/kaniko/testutil"
)

func Test_GetContainerAndBlob(t *testing.T) {
	tests := []struct {
		name                  string
		context               string
		expectedAccountName   string
		expectedContainerName string
		expectedBlobName      string
	}{
		{
			name:                  "blob with fold structure",
			context:               "accountname/containername/fold/blobname",
			expectedAccountName:   "accountname",
			expectedContainerName: "containername",
			expectedBlobName:      "fold/blobname",
		},
		{
			name:                  "blob without fold",
			context:               "accountname/containername/blobname",
			expectedAccountName:   "accountname",
			expectedContainerName: "containername",
			expectedBlobName:      "blobname",
		},
		{
			name:                  "without blobname (default blobname)",
			context:               "accountname/containername",
			expectedAccountName:   "accountname",
			expectedContainerName: "containername",
			expectedBlobName:      constants.ContextTar,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotAccountName, gotContainerName, gotBlobName := GetContainerAndBlob(test.context)
			testutil.CheckDeepEqual(t, test.expectedAccountName, gotAccountName)
			testutil.CheckDeepEqual(t, test.expectedContainerName, gotContainerName)
			testutil.CheckDeepEqual(t, test.expectedBlobName, gotBlobName)
		})
	}
}
