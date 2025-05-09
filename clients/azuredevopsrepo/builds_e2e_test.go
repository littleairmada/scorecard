// Copyright 2024 OpenSSF Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package azuredevopsrepo

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ossf/scorecard/v5/clients"
)

var _ = Describe("E2E TEST: azuredevopsrepo.buildsHandler", func() {
	Context("Builds - Azure DevOps", func() {
		It("Should return successful builds", func() {
			repo, err := MakeAzureDevOpsRepo("https://dev.azure.com/jamiemagee/jamiemagee/_git/jamiemagee")
			Expect(err).Should(BeNil())

			repoClient, err := CreateAzureDevOpsClient(context.Background(), repo)
			Expect(err).Should(BeNil())

			err = repoClient.InitRepo(repo, clients.HeadSHA, 0)
			Expect(err).Should(BeNil())

			builds, err := repoClient.ListSuccessfulWorkflowRuns("azure-pipelines.yml")
			Expect(err).Should(BeNil())
			Expect(len(builds)).Should(BeNumerically(">", 0))
		})
	})
})
