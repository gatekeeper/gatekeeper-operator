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

package merge

import (
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// RetainClusterObjectFields updates the desired object with values retained
// from the cluster object.
func RetainClusterObjectFields(desiredObj, clusterObj *unstructured.Unstructured) error {
	// Pass the same ResourceVersion as in the cluster object for update
	// operation, otherwise operation will fail.
	desiredObj.SetResourceVersion(clusterObj.GetResourceVersion())

	if desiredObj.GetKind() == "Service" {
		return retainServiceFields(desiredObj, clusterObj)
	}
	return nil
}
func retainServiceFields(desiredObj, clusterObj *unstructured.Unstructured) error {
	// ClusterIP is allocated to Service by cluster, so if it exists, retain it
	// while updating.
	clusterIP, ok, err := unstructured.NestedString(clusterObj.Object, "spec", "clusterIP")
	if err != nil {
		return errors.Wrap(err, "Error retrieving clusterIP from cluster service")
	} else if ok && clusterIP != "" {
		err := unstructured.SetNestedField(desiredObj.Object, clusterIP, "spec", "clusterIP")
		if err != nil {
			return errors.Wrap(err, "Error setting clusterIP for service")
		}
	} // !ok could indicate that a clusterIP was not assigned

	return nil
}
