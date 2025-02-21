// Copyright 2023 Red Hat, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package profiles

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kiegroup/kogito-serverless-operator/test"
)

func Test_workflowIsDevProfile(t *testing.T) {
	workflowWithDevProfile := test.GetKogitoServerlessWorkflow("../../config/samples/"+test.KogitoServerlessWorkflowSampleDevModeYamlCR, t.Name())
	assert.True(t, IsDevProfile(workflowWithDevProfile))

	workflowWithNoProfile := test.GetKogitoServerlessWorkflow("../../config/samples/"+test.KogitoServerlessWorkflowSampleYamlCR, t.Name())
	assert.False(t, IsDevProfile(workflowWithNoProfile))

	workflowWithProdProfile := test.GetKogitoServerlessWorkflow("../../config/samples/"+test.KogitoServerlessWorkflowProdProfileSampleYamlCR, t.Name())
	assert.False(t, IsDevProfile(workflowWithProdProfile))
}
