// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/pages_project"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func TestPagesProjectModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*pages_project.PagesProjectModel)(nil)
	schema := pages_project.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}

func TestPagesProjectAccountIDRequired(t *testing.T) {
	t.Parallel()
	s := pages_project.ResourceSchema(context.TODO())
	attr, ok := s.Attributes["account_id"]
	if !ok {
		t.Fatalf("account_id attribute missing from schema")
	}
	stringAttr, ok := attr.(resourceschema.StringAttribute)
	if !ok {
		t.Fatalf("account_id attribute is %T, want resource schema.StringAttribute", attr)
	}
	if !stringAttr.Required {
		t.Fatalf("account_id should be required")
	}
}
