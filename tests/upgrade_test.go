//
// DISCLAIMER
//
// Copyright 2018 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//
// Author Jan Christoph Uhde <jan@uhdejc.com>
//
package tests

import (
	"context"
	"fmt"
	"strings"
	"testing"

	api "github.com/arangodb/kube-arangodb/pkg/apis/deployment/v1alpha"
	kubeArangoClient "github.com/arangodb/kube-arangodb/pkg/client"
	"github.com/arangodb/kube-arangodb/pkg/util"
	"github.com/dchest/uniuri"
)

// func TestUpgradeClusterRocksDB33pto34p(t *testing.T) {
// 	upgradeSubTest(t, api.DeploymentModeCluster, api.StorageEngineRocksDB, "arangodb/arangodb-preview:3.3", "arangodb/arangodb-preview:3.4")
// }

// test upgrade single server mmfiles 3.2 -> 3.3
// func TestUpgradeSingleMMFiles32to33(t *testing.T) {
// 	upgradeSubTest(t, api.DeploymentModeSingle, api.StorageEngineMMFiles, "arangodb/arangodb:3.2.16", "arangodb/arangodb:3.3.13")
// }

// // test upgrade single server rocksdb 3.3 -> 3.4
// func TestUpgradeSingleRocksDB33to34(t *testing.T) {
// 	upgradeSubTest(t, api.DeploymentModeSingle, api.StorageEngineRocksDB, "3.3.13", "3.4.0")
// }

/*// test upgrade active-failover server rocksdb 3.3 -> 3.4
func TestUpgradeActiveFailoverRocksDB33to34(t *testing.T) {
	upgradeSubTest(t, api.DeploymentModeActiveFailover, api.StorageEngineRocksDB, "3.3.13", "3.4.0")
}*/

// // test upgrade active-failover server mmfiles 3.3 -> 3.4
// func TestUpgradeActiveFailoverMMFiles33to34(t *testing.T) {
// 	upgradeSubTest(t, api.DeploymentModeActiveFailover, api.StorageEngineMMFiles, "3.3.13", "3.4.0")
// }

// test upgrade cluster rocksdb 3.2 -> 3.3
// func TestUpgradeClusterRocksDB32to33(t *testing.T) {
// 	upgradeSubTest(t, api.DeploymentModeCluster, api.StorageEngineRocksDB, "3.2.16", "3.3.13")
// }

// // test upgrade cluster mmfiles 3.3 -> 3.4
// func TestUpgradeClusterMMFiles33to34(t *testing.T) {
// 	upgradeSubTest(t, api.DeploymentModeCluster, api.StorageEngineRocksDB, "3.3.13", "3.4.0")
// }

// // test downgrade single server mmfiles 3.3.17 -> 3.3.16
// func TestDowngradeSingleMMFiles3317to3316(t *testing.T) {
// 	upgradeSubTest(t, api.DeploymentModeSingle, api.StorageEngineMMFiles, "arangodb/arangodb:3.3.16", "arangodb/arangodb:3.3.17")
// }

// // test downgrade ActiveFailover server rocksdb 3.3.17 -> 3.3.16
// func TestDowngradeActiveFailoverRocksDB3317to3316(t *testing.T) {
// 	upgradeSubTest(t, api.DeploymentModeActiveFailover, api.StorageEngineRocksDB, "arangodb/arangodb:3.3.16", "arangodb/arangodb:3.3.17")
// }

// // test downgrade cluster rocksdb 3.3.17 -> 3.3.16
// func TestDowngradeClusterRocksDB3317to3316(t *testing.T) {
// 	upgradeSubTest(t, api.DeploymentModeCluster, api.StorageEngineRocksDB, "arangodb/arangodb:3.3.16", "arangodb/arangodb:3.3.17")
// }

func upgradeSubTest(t *testing.T, mode api.DeploymentMode, engine api.StorageEngine, fromImage, toImage string) error {
	// check environment
	longOrSkip(t)

	ns := getNamespace(t)
	kubecli := mustNewKubeClient(t)
	c := kubeArangoClient.MustNewInCluster()

	depl := newDeployment(strings.Replace(fmt.Sprintf("tu-%s-%s-%s", mode[:2], engine[:2], uniuri.NewLen(4)), ".", "", -1))
	depl.Spec.Mode = api.NewMode(mode)
	depl.Spec.StorageEngine = api.NewStorageEngine(engine)
	depl.Spec.TLS = api.TLSSpec{} // should auto-generate cert
	depl.Spec.Image = util.NewString(fromImage)
	depl.Spec.SetDefaults(depl.GetName()) // this must be last

	// Create deployment
	deployment, err := c.DatabaseV1alpha().ArangoDeployments(ns).Create(depl)
	if err != nil {
		t.Fatalf("Create deployment failed: %v", err)
	}
	defer deferedCleanupDeployment(c, depl.GetName(), ns)

	// Wait for deployment to be ready
	deployment, err = waitUntilDeployment(c, depl.GetName(), ns, deploymentIsReady())
	if err != nil {
		t.Fatalf("Deployment not running in time: %v", err)
	}

	// Create a database client
	ctx := context.Background()
	DBClient := mustNewArangodDatabaseClient(ctx, kubecli, deployment, t, nil)

	if err := waitUntilArangoDeploymentHealthy(deployment, DBClient, kubecli, ""); err != nil {
		t.Fatalf("Deployment not healthy in time: %v", err)
	}

	// Try to change image version
	deployment, err = updateDeployment(c, depl.GetName(), ns,
		func(spec *api.DeploymentSpec) {
			spec.Image = util.NewString(toImage)
		})
	if err != nil {
		t.Fatalf("Failed to upgrade the Image from version : " + fromImage + " to version: " + toImage)
	} else {
		t.Log("Updated deployment")
	}

	deployment, err = waitUntilDeployment(c, depl.GetName(), ns, deploymentIsReady())
	if err != nil {
		t.Fatalf("Deployment not running in time: %v", err)
	} else {
		t.Log("Deployment running")
	}

	if err := waitUntilArangoDeploymentHealthy(deployment, DBClient, kubecli, toImage); err != nil {
		t.Fatalf("Deployment not healthy in time: %v", err)
	} else {
		t.Log("Deployment healthy")
	}

	// Cleanup
	removeDeployment(c, depl.GetName(), ns)

	return nil
}
