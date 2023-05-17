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
	"fmt"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/gatekeeper/gatekeeper-operator/pkg/util"
)

// RetainClusterObjectFields updates the desired object with values retained
// from the cluster object.
func RetainClusterObjectFields(desiredObj, clusterObj *unstructured.Unstructured) error {
	// Pass the same ResourceVersion as in the cluster object for update
	// operation, otherwise operation will fail.
	desiredObj.SetResourceVersion(clusterObj.GetResourceVersion())

	switch desiredObj.GetKind() {
	case util.ServiceKind:
		return retainServiceFields(desiredObj, clusterObj)
	case util.ValidatingWebhookConfigurationKind:
		fallthrough
	case util.MutatingWebhookConfigurationKind:
		return retainWebhookConfigurationFields(desiredObj, clusterObj)
	case util.SecretKind:
		return retainSecretFields(desiredObj, clusterObj)
	default:
		return nil
	}
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

func retainSecretFields(desiredObj, clusterObj *unstructured.Unstructured) error {
	data, ok, err := unstructured.NestedMap(clusterObj.Object, "data")
	if err != nil {
		return errors.Wrap(err, "Error retrieving data from secret")
	} else if ok && len(data) != 0 {
		err := unstructured.SetNestedMap(desiredObj.Object, data, "data")
		if err != nil {
			return errors.Wrap(err, "Error setting data for secret")
		}
	}
	return nil
}

func retainWebhookConfigurationFields(desiredObj, clusterObj *unstructured.Unstructured) error {
	// Retain each webhook's CABundle
	clusterWebhooks, ok, err := unstructured.NestedSlice(clusterObj.Object, "webhooks")
	if err != nil {
		return errors.Wrapf(err, "Error retrieving webhooks from cluster object %s", clusterObj.GetKind())
	} else if ok && len(clusterWebhooks) == 0 {
		err = unstructured.SetNestedSlice(desiredObj.Object, nil, "webhooks")
		if err != nil {
			return errors.Wrapf(err, "Error setting webhooks for desired object %s", desiredObj.GetKind())
		}
		return nil
	} else if !ok {
		return nil
	}

	desiredWebhooks, ok, err := unstructured.NestedSlice(desiredObj.Object, "webhooks")
	if err != nil {
		return errors.Wrapf(err, "Error retrieving webhooks from desired object %s", desiredObj.GetKind())
	} else if !ok {
		// Should never happen
		return fmt.Errorf("webhooks field not found for desired object %s", desiredObj.GetKind())
	}

	for i := range desiredWebhooks {
		for j := range clusterWebhooks {
			desiredWebhook := desiredWebhooks[i].(map[string]interface{})
			clusterWebhook := clusterWebhooks[j].(map[string]interface{})
			if desiredWebhook["name"] != clusterWebhook["name"] {
				continue
			}

			caBundle, ok, err := unstructured.NestedFieldNoCopy(clusterWebhook, "clientConfig", "caBundle")
			if err != nil {
				return errors.Wrapf(err, "Error retrieving webhooks[%d].clientConfig.caBundle from cluster object %s", j, clusterObj.GetKind())
			} else if !ok {
				// If no caBundle is configured, assume the system's CAs will be used.
				break
			}

			err = unstructured.SetNestedField(desiredWebhook, caBundle, "clientConfig", "caBundle")
			if err != nil {
				return errors.Wrapf(err, "Error setting webhooks[%d].clientConfig.caBundle for desired object %s", i, desiredObj.GetKind())
			}
			break
		}
	}

	err = unstructured.SetNestedSlice(desiredObj.Object, desiredWebhooks, "webhooks")
	if err != nil {
		return errors.Wrapf(err, "Error setting webhooks for desired object %s", desiredObj.GetKind())
	}
	return nil
}
