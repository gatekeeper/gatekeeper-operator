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
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/kubernetes/scheme"

	"github.com/gatekeeper/gatekeeper-operator/pkg/bindata"
)

var staticAssetsDir = "config/gatekeeper-rendered/"

func GetManifestObject(asset string) (*unstructured.Unstructured, error) {
	assetName := staticAssetsDir + asset
	bytes, err := bindata.Asset(assetName)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to retrieve bindata asset %s", assetName)
	}

	obj, err := unmarshalJSON(bytes)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to unmarshal YAML bytes for asset name %s", assetName)
	}
	return obj, nil
}

func unmarshalJSON(in []byte) (*unstructured.Unstructured, error) {
	if in == nil {
		return nil, errors.New("input bytes is nil")
	}

	objInterface, _, err := scheme.Codecs.UniversalDecoder().Decode(in, nil, &unstructured.Unstructured{})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode bytes")
	}
	obj, ok := objInterface.(*unstructured.Unstructured)
	if !ok {
		return nil, fmt.Errorf("type assertion of object interface to *unstructured.Unstructured failed, got %T", obj)
	}

	return obj, nil
}

// ToMap Convenience method to convert any struct into a map
func ToMap(obj interface{}) map[string]interface{} {
	var result map[string]interface{}
	resultRec, _ := json.Marshal(obj)
	json.Unmarshal(resultRec, &result)
	return result
}

// ToArg Converts a key, value pair into a valid container argument. e.g. '--argName', 'argValue' returns '--argName=argValue'
func ToArg(name, value string) string {
	return name + "=" + value
}

// FromArg Converts a container argument into a key, value pair. e.g. '--argName=argValue' returns '--argName', 'argValue'
func FromArg(arg string) (key, value string) {
	parts := strings.Split(arg, "=")
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[1]
}
