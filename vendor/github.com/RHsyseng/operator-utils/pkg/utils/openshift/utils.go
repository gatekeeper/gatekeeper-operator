package openshift

import (
	"github.com/RHsyseng/operator-utils/internal/platform"
	"k8s.io/client-go/rest"
)

/*
GetPlatformInfo examines the Kubernetes-based environment and determines the running platform, version, & OS.
Accepts <nil> or instantiated 'cfg' rest config parameter.

Result: PlatformInfo{ Name: OpenShift, K8SVersion: 1.13+, OS: linux/amd64 }
*/
func GetPlatformInfo(cfg *rest.Config) (platform.PlatformInfo, error) {
	return platform.K8SBasedPlatformVersioner{}.GetPlatformInfo(nil, cfg)
}

/*
IsOpenShift is a helper method to simplify boolean OCP checks against GetPlatformInfo results
Accepts <nil> or instantiated 'cfg' rest config parameter.
*/
func IsOpenShift(cfg *rest.Config) (bool, error) {
	info, err := GetPlatformInfo(cfg)
	if err != nil {
		return false, err
	}
	return info.IsOpenShift(), nil
}

/*
GetPlatformName is a helper method to return the platform name from GetPlatformInfo results
Accepts <nil> or instantiated 'cfg' rest config parameter.
*/
func GetPlatformName(cfg *rest.Config) (string, error) {
	info, err := GetPlatformInfo(cfg)
	if err != nil {
		return "", err
	}
	return string(info.Name), nil
}

/*
LookupOpenShiftVersion fetches OpenShift version info from API endpoints
*** NOTE: OCP 4.1+ requires elevated user permissions, see PlatformVersioner for details
Accepts <nil> or instantiated 'cfg' rest config parameter.

Result: OpenShiftVersion{ Version: 4.1.2 }
*/
func LookupOpenShiftVersion(cfg *rest.Config) (platform.OpenShiftVersion, error) {
	return platform.K8SBasedPlatformVersioner{}.LookupOpenShiftVersion(nil, cfg)
}

/*
Supported platform: OpenShift
cfg : OpenShift platform config, use runtime config if nil is passed in.
version: Supported version format : Major.Minor
	       e.g.: 4.3
*/
func CompareOpenShiftVersion(cfg *rest.Config, version string) (int, error) {
	return platform.K8SBasedPlatformVersioner{}.CompareOpenShiftVersion(nil, cfg, version)
}

/*
MapKnownVersion maps from K8S version of PlatformInfo to equivalent OpenShift version

Result: OpenShiftVersion{ Version: 4.1.2 }
*/
func MapKnownVersion(info platform.PlatformInfo) platform.OpenShiftVersion {
	return platform.MapKnownVersion(info)
}
