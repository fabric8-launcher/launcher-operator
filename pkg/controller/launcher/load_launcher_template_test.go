package launcher

import (
	"fabric8-launcher/launcher-operator/pkg/helper"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/rest"
)

func TestLoadLauncherTemplate(t *testing.T) {
	templateHelper := helper.NewOpenshiftTemplateHelper()
	templatePath := "../../../templates"
	templateName := "fabric8-launcher"
	tpl, err := templateHelper.Load(&rest.Config{}, templatePath, templateName)

	assert.Nil(t, err, "Error creating openshift template for %s", templateName)
	assert.NotNil(t, tpl, "Invalid openshift template for %s", templateName)

	assert.NotEmpty(t, tpl.Raw, "Invalid openshift template for %s", templateName)

}
