package servicecatalogaddons_test

import (
	"testing"

	"github.com/kyma-project/kyma/components/console-backend-service/internal/domain/servicecatalogaddons"
	"github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"
	"github.com/kyma-project/kyma/components/service-binding-usage-controller/pkg/apis/servicecatalog/v1alpha1"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestUsageKindConverter_ToGQL(t *testing.T) {
	// GIVEN
	conv := servicecatalogaddons.NewUsageKindConverter()
	name := "harry"

	// WHEN
	result := conv.ToGQL(fixUsageKind(name))

	// THEN
	assert.Equal(t, result, fixUsageKindGQL(name))
}

func fixUsageKind(name string) *v1alpha1.UsageKind {
	return &v1alpha1.UsageKind{
		TypeMeta: v1.TypeMeta{
			APIVersion: "servicecatalog.kyma-project.io/v1alpha1",
			Kind:       "usagekind",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: name,
		},
		Spec: v1alpha1.UsageKindSpec{
			DisplayName: fixUsageKindDisplayName(),
			Resource: &v1alpha1.ResourceReference{
				Group:   fixResource().Group,
				Kind:    fixResource().Kind,
				Version: fixResource().Version,
			},
			LabelsPath: fixUsageKindLabelsPath(),
		},
	}
}

func fixUsageKindGQL(name string) *gqlschema.UsageKind {
	return &gqlschema.UsageKind{
		Name:        name,
		Group:       fixResource().Group,
		Kind:        fixResource().Kind,
		Version:     fixResource().Version,
		DisplayName: fixUsageKindDisplayName(),
	}
}

func fixUsageKindDisplayName() string {
	return "target"
}

func fixUsageKindLabelsPath() string {
	return "meta.data"
}

func fixResource() *v1alpha1.ResourceReference {
	return &v1alpha1.ResourceReference{
		Group:   "kubeless.io",
		Kind:    "function",
		Version: "v1beta1",
	}
}

func fixUsageKindResourceNamespace() string {
	return "space"
}
