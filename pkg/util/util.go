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
	"github.com/openshift/library-go/pkg/manifest"
	"github.com/pkg/errors"

	"github.com/gatekeeper/gatekeeper-operator/pkg/bindata"
)

var (
	staticAssetsDir = "config/gatekeeper/"
)

func GetManifest(asset string) (*manifest.Manifest, error) {
	manifest := &manifest.Manifest{}
	assetName := staticAssetsDir + asset
	bytes, err := bindata.Asset(assetName)
	if err != nil {
		return manifest, errors.Wrapf(err, "Unable to retrieve bindata asset %s", assetName)
	}

	err = manifest.UnmarshalJSON(bytes)
	if err != nil {
		return manifest, errors.Wrapf(err, "Unable to unmarshal YAML bytes for asset name %s", assetName)
	}
	return manifest, nil
}
