// Copyright 2020 The Lokomotive Authors
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

// +build aws aws-edge packet
// +build e2e

package gangway

import (
	"testing"
	"time"

	testutil "github.com/kinvolk/lokomotive/test/components/util"
)

func TestGangwayDeployment(t *testing.T) {
	namespace := "gangway"

	client := testutil.CreateKubeClient(t)

	t.Run("deployment", func(t *testing.T) {
		t.Parallel()
		deployment := "gangway"

		testutil.WaitForDeployment(t, client, namespace, deployment, time.Second*5, time.Minute*5)
	})
}
