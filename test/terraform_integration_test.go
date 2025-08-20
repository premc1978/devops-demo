package test

import (
	"testing"
	"os"
	"path/filepath"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformqa(t *testing.T) {
	t.Parallel()

	// Use environment variable or default path for Terraform code
	tfDir := os.Getenv("TF_TEST_DIR")
	if tfDir == "" {
		tfDir = filepath.Join("..", "tf", "environments", "qa")
	}

	terraformOptions := &terraform.Options{
		TerraformDir: tfDir,
		// Optionally set variables or backend config here
		// Vars: map[string]interface{}{
		//     "example_var": "value",
		// },
		NoColor: true,
	}

	// Ensure resources are destroyed at the end of the test
	defer terraform.Destroy(t, terraformOptions)

	// Init and apply Terraform code
	terraform.InitAndApply(t, terraformOptions)

	// Example: Validate output variable
	output := terraform.Output(t, terraformOptions, "example_output")
	assert.NotEmpty(t, output, "example_output should not be empty")

	// Add more assertions as needed for your infrastructure
}
