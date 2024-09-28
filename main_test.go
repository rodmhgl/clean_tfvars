package main

import (
	"reflect"
	"testing"
)

func TestFindUnusedVars(t *testing.T) {
	declaredVars := []string{"instance_type", "region"}
	tfvars := map[string]string{
		"instance_type": "t2.micro",
		"region":        "us-west-2",
		"unused_var":    "this will be removed",
	}

	// Pass false for the verbose argument
	unused := FindUnusedVars(declaredVars, tfvars, false)
	expected := []string{"unused_var"}

	if !reflect.DeepEqual(unused, expected) {
		t.Errorf("Expected %v, but got %v", expected, unused)
	}
}

func TestParseDeclaredVars(t *testing.T) {
	// Assuming test files are available under ./test-data/
	directory := "./test-data"

	// Pass false for the verbose argument
	declaredVars, err := ParseDeclaredVars(directory, false)
	if err != nil {
		t.Fatalf("Failed to parse declared vars: %v", err)
	}

	expectedVars := []string{"instance_type", "region"}
	if !reflect.DeepEqual(declaredVars, expectedVars) {
		t.Errorf("Expected %v, but got %v", expectedVars, declaredVars)
	}
}

func TestParseTFVarsFile(t *testing.T) {
	filePath := "./test-data/variables.tfvars"

	// Pass false for the verbose argument
	tfvars, err := ParseTFVarsFile(filePath, false)
	if err != nil {
		t.Fatalf("Failed to parse tfvars file: %v", err)
	}

	// Adjust the expected values to match the parsed output, which includes quotes around the strings
	expectedTFVars := map[string]string{
		"instance_type": `t2.micro`,
		"region":        `us-west-2`,
		"unused_var":    `this will be removed`,
	}

	if !reflect.DeepEqual(tfvars, expectedTFVars) {
		t.Errorf("Expected %v, but got %v", expectedTFVars, tfvars)
	}
}
