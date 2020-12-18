/*
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
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"
)

var (
	DefaultGatekeeperNamespace          = "gatekeeper-system"
	DefaultOpenShiftGatekeeperNamespace = "openshift-gatekeeper-system"
)

// GetOperatorNamespace returns the namespace the operator is running in from
// the associated service account secret.
func GetOperatorNamespace() (string, error) {
	nsBytes, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		if os.IsNotExist(err) {
			return "", errors.New("namespace not found for current environment")
		}
		return "", err
	}
	ns := strings.TrimSpace(string(nsBytes))
	return ns, nil
}

// GetPlatformNamespace returns the namespace for the designated platform.
func GetPlatformNamespace(platformName string) string {
	switch PlatformType(platformName) {
	case OpenShift:
		return DefaultOpenShiftGatekeeperNamespace
	default:
		return DefaultGatekeeperNamespace
	}
}
