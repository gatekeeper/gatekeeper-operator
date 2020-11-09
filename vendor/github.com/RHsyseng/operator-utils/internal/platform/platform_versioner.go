package platform

import (
	"context"
	"encoding/json"
	"errors"

	openapi_v2 "github.com/googleapis/gnostic/openapiv2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var (
	log                   = logf.Log.WithName("utils")
	clusterVersionAPIPath = "apis/config.openshift.io/v1/clusterversions/version"
)

type PlatformVersioner interface {
	GetPlatformInfo(discoverer Discoverer, cfg *rest.Config) (PlatformInfo, error)
}

type Discoverer interface {
	ServerVersion() (*version.Info, error)
	ServerGroups() (*v1.APIGroupList, error)
	OpenAPISchema() (*openapi_v2.Document, error)
	RESTClient() rest.Interface
}

type K8SBasedPlatformVersioner struct{}

/*
MapKnownVersion maps from K8S version of PlatformInfo to equivalent OpenShift version

Result: OpenShiftVersion{ Version: 4.1.2 }
*/
func MapKnownVersion(info PlatformInfo) OpenShiftVersion {
	k8sToOcpMap := map[string]string{
		"1.10+": "3.10",
		"1.10":  "3.10",
		"1.11+": "3.11",
		"1.11":  "3.11",
		"1.13+": "4.1",
		"1.13":  "4.1",
		"1.14+": "4.2",
		"1.14":  "4.2",
		"1.16+": "4.3",
		"1.16":  "4.3",
		"1.17+": "4.4",
		"1.17":  "4.4",
		"1.18+": "4.5",
		"1.18":  "4.5",
		"1.19+": "4.6",
		"1.19":  "4.6",
	}
	return OpenShiftVersion{Version: k8sToOcpMap[info.K8SVersion]}
}

// deal with cfg coming from legacy method signature and allow injection for client testing
func (K8SBasedPlatformVersioner) DefaultArgs(client Discoverer, cfg *rest.Config) (Discoverer, *rest.Config, error) {
	if cfg == nil {
		var err error
		cfg, err = config.GetConfig()
		if err != nil {
			return nil, nil, err
		}
	}
	if client == nil {
		var err error
		client, err = discovery.NewDiscoveryClientForConfig(cfg)
		if err != nil {
			return nil, nil, err
		}
	}
	return client, cfg, nil
}

func (pv K8SBasedPlatformVersioner) GetPlatformInfo(client Discoverer, cfg *rest.Config) (PlatformInfo, error) {
	log.Info("detecting platform version...")
	info := PlatformInfo{Name: Kubernetes}

	var err error
	client, cfg, err = pv.DefaultArgs(client, cfg)
	if err != nil {
		log.Info("issue occurred while defaulting client/cfg args")
		return info, err
	}

	k8sVersion, err := client.ServerVersion()
	if err != nil {
		log.Info("issue occurred while fetching ServerVersion")
		return info, err
	}
	info.K8SVersion = k8sVersion.Major + "." + k8sVersion.Minor
	info.OS = k8sVersion.Platform

	apiList, err := client.ServerGroups()
	if err != nil {
		log.Info("issue occurred while fetching ServerGroups")
		return info, err
	}

	for _, v := range apiList.Groups {
		if v.Name == "route.openshift.io" {

			log.Info("route.openshift.io found in apis, platform is OpenShift")
			info.Name = OpenShift
			break
		}
	}
	log.Info(info.String())
	return info, nil
}

/*
OCP4.1+ requires elevated cluster configuration user security permissions for version fetch
REST call URL requiring permissions: /apis/config.openshift.io/v1/clusterversions
*/
func (pv K8SBasedPlatformVersioner) LookupOpenShiftVersion(client Discoverer, cfg *rest.Config) (OpenShiftVersion, error) {

	osv := OpenShiftVersion{}
	client, _, err := pv.DefaultArgs(nil, nil)
	if err != nil {
		log.Info("issue occurred while defaulting args for version lookup")
		return osv, err
	}
	doc, err := client.OpenAPISchema()
	if err != nil {
		log.Info("issue occurred while fetching OpenAPISchema")
		return osv, err
	}

	switch doc.Info.Version[:4] {
	case "v3.1":
		osv.Version = doc.Info.Version

	// OCP4 returns K8S major/minor from old API endpoint [bugzilla-1658957]
	case "v1.1":
		rest := client.RESTClient().Get().AbsPath(clusterVersionAPIPath)

		result := rest.Do(context.TODO())
		if result.Error() != nil {
			log.Info("issue making API version rest call: " + result.Error().Error())
			return osv, result.Error()
		}

		// error handling before/after Raw() seems redundant, but error detail can be lost in convert
		body, err := result.Raw()
		if err != nil {
			log.Info("issue pulling raw result from API call")
			return osv, err
		}

		var cvi PlatformClusterInfo
		err = json.Unmarshal(body, &cvi)
		if err != nil {
			log.Info("issue occurred while unmarshalling PlatformClusterInfo")
			return osv, err
		}
		osv.Version = cvi.Status.Desired.Version
	}
	return osv, nil
}

func (pv K8SBasedPlatformVersioner) CompareOpenShiftVersion(client Discoverer, cfg *rest.Config, version string) (int, error) {
	info, err := pv.GetPlatformInfo(client, cfg)
	if err != nil {
		return -1, err
	}
	if !info.IsOpenShift() {
		return -1, errors.New("There is no OpenShift platform detected.")
	}
	curVersion := MapKnownVersion(info)
	return curVersion.Compare(OpenShiftVersion{Version: version})
}
