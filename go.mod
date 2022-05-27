module github.com/gatekeeper/gatekeeper-operator

go 1.16

replace golang.org/x/crypto => golang.org/x/crypto v0.0.0-20211202192323-5770296d904e // CVE-2021-43565
replace github.com/prometheus/client_golang => github.com/prometheus/client_golang v1.11.1 // CVE-2022-21698

require (
	github.com/go-bindata/go-bindata v3.1.2+incompatible
	github.com/go-logr/logr v0.4.0
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.13.0
	github.com/pkg/errors v0.9.1
	k8s.io/api v0.21.2
	k8s.io/apiextensions-apiserver v0.21.2
	k8s.io/apimachinery v0.21.2
	k8s.io/client-go v0.21.2
	sigs.k8s.io/controller-runtime v0.9.2
)
