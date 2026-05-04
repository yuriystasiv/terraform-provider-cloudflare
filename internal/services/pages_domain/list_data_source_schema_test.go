// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_domain_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/pages_domain"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func TestPagesDomainsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*pages_domain.PagesDomainsDataSourceModel)(nil)
	schema := pages_domain.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}

func TestPagesDomainsDataSourceAccountIDRequired(t *testing.T) {
	t.Parallel()
	s := pages_domain.ListDataSourceSchema(context.TODO())
	attr, ok := s.Attributes["account_id"]
	if !ok {
		t.Fatalf("account_id attribute missing from schema")
	}
	stringAttr, ok := attr.(datasourceschema.StringAttribute)
	if !ok {
		t.Fatalf("account_id attribute is %T, want data source schema.StringAttribute", attr)
	}
	if !stringAttr.Required {
		t.Fatalf("account_id should be required")
	}
}
