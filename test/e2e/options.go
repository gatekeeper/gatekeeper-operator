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

package e2e

import (
	"flag"
	"time"

	"github.com/gatekeeper/gatekeeper-operator/pkg/util"
)

var gatekeeperNamespace string
var pollInterval time.Duration
var timeout time.Duration
var deleteTimeout time.Duration

func init() {
	flag.StringVar(&gatekeeperNamespace, "namespace", util.DefaultGatekeeperNamespace, "The namespace to run tests")
	flag.DurationVar(&pollInterval, "poll-interval", 1*time.Second, "The length of time between polls")
	flag.DurationVar(&timeout, "timeout", 1*time.Minute, "The length of time to poll before giving up")
	flag.DurationVar(&deleteTimeout, "delete-timeout", 5*time.Minute, "The length of time to wait for deletion of all resources")
}
