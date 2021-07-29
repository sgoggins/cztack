package test

import (
	"testing"

	"github.com/chanzuckerberg/go-misc/tftest"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

func TestAWSIAMRoleEC2Poweruser(t *testing.T) {
	test := tftest.Test{
		Setup: func(t *testing.T) *terraform.Options {
			curAcct := tftest.AWSCurrentAccountID(t)

			return tftest.Options(
				tftest.IAMRegion,
				map[string]interface{}{
					"role_name":         random.UniqueId(),
					"source_account_id": curAcct,
				},
			)
		},
		Validate: func(t *testing.T, options *terraform.Options) {
			r := require.New(t)
			region := options.EnvVars["AWS_DEFAULT_REGION"]
			r.NotEmpty(region)
			outputs := terraform.OutputAll(t, options)
			r.NotEmpty(outputs)
			roleName := outputs["role_name"].(string)
			r.NotEmpty(roleName)
			arn := outputs["arn"].(string)
			r.NotEmpty(arn)
		},
	}

	test.Run(t)
}
