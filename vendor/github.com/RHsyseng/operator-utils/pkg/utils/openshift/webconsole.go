package openshift

import (
	"errors"
	"github.com/RHsyseng/operator-utils/pkg/resource"
	"github.com/ghodss/yaml"
	consolev1 "github.com/openshift/api/console/v1"
	amv1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
)

func GetConsoleYAMLSample(res resource.KubernetesResource) (*consolev1.ConsoleYAMLSample, error) {
	annotations := res.GetAnnotations()
	snippetStr := annotations["consoleSnippet"]
	var snippet bool = false
	if tmp, err := strconv.ParseBool(snippetStr); err == nil {
		snippet = tmp
	}

	targetAPIVersion, _ := annotations["consoleTargetAPIVersion"]
	if targetAPIVersion == "" {
		targetAPIVersion = res.GetObjectKind().GroupVersionKind().GroupVersion().String()
	}

	targetKind := annotations["consoleTargetKind"]
	if targetKind == "" {
		targetKind = res.GetObjectKind().GroupVersionKind().Kind
	}

	defaultText := res.GetName() + "-yamlsample"
	title, _ := annotations["consoleTitle"]
	if title == "" {
		title = defaultText
	}
	desc, _ := annotations["consoleDesc"]
	if desc == "" {
		desc = defaultText
	}
	name, _ := annotations["consoleName"]
	if name == "" {
		name = defaultText
	}

	delete(annotations, "consoleSnippet")
	delete(annotations, "consoleTitle")
	delete(annotations, "consoleDesc")
	delete(annotations, "consoleName")
	delete(annotations, "consoleTargetAPIVersion")
	delete(annotations, "consoleTargetKind")

	data, err := yaml.Marshal(res)
	if err != nil {
		return nil, errors.New("Failed to convert to yamlstr from KubernetesResource.")
	}

	yamlSample := &consolev1.ConsoleYAMLSample{
		ObjectMeta: amv1.ObjectMeta{
			Name:      name,
			Namespace: "openshift-console",
		},
		Spec: consolev1.ConsoleYAMLSampleSpec{
			TargetResource: amv1.TypeMeta{
				APIVersion: targetAPIVersion,
				Kind:       targetKind,
			},
			Title:       consolev1.ConsoleYAMLSampleTitle(title),
			Description: consolev1.ConsoleYAMLSampleDescription(desc),
			YAML:        consolev1.ConsoleYAMLSampleYAML(string(data)),
			Snippet:     snippet,
		},
	}
	return yamlSample, nil
}
