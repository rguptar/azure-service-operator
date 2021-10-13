// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpecARM, DatabaseAccountsSqlDatabasesContainersStoredProceduresSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpecARM runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpecARM(subject DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM instances for property testing - lazily
//instantiated by DatabaseAccountsSqlDatabasesContainersStoredProceduresSpecARMGenerator()
var databaseAccountsSqlDatabasesContainersStoredProceduresSpecARMGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesContainersStoredProceduresSpecARMGenerator returns a generator of DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM instances for property testing.
// We first initialize databaseAccountsSqlDatabasesContainersStoredProceduresSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesContainersStoredProceduresSpecARMGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesContainersStoredProceduresSpecARMGenerator != nil {
		return databaseAccountsSqlDatabasesContainersStoredProceduresSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpecARM(generators)
	databaseAccountsSqlDatabasesContainersStoredProceduresSpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpecARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpecARM(generators)
	databaseAccountsSqlDatabasesContainersStoredProceduresSpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM{}), generators)

	return databaseAccountsSqlDatabasesContainersStoredProceduresSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpecARM(gens map[string]gopter.Gen) {
	gens["APIVersion"] = gen.OneConstOf(DatabaseAccountsSqlDatabasesContainersStoredProceduresSpecAPIVersion20210515)
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.OneConstOf(DatabaseAccountsSqlDatabasesContainersStoredProceduresSpecTypeMicrosoftDocumentDBDatabaseAccountsSqlDatabasesContainersStoredProcedures)
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = SqlStoredProcedureCreateUpdatePropertiesARMGenerator()
}

func Test_SqlStoredProcedureCreateUpdatePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlStoredProcedureCreateUpdatePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlStoredProcedureCreateUpdatePropertiesARM, SqlStoredProcedureCreateUpdatePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlStoredProcedureCreateUpdatePropertiesARM runs a test to see if a specific instance of SqlStoredProcedureCreateUpdatePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlStoredProcedureCreateUpdatePropertiesARM(subject SqlStoredProcedureCreateUpdatePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlStoredProcedureCreateUpdatePropertiesARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SqlStoredProcedureCreateUpdatePropertiesARM instances for property testing - lazily instantiated by
//SqlStoredProcedureCreateUpdatePropertiesARMGenerator()
var sqlStoredProcedureCreateUpdatePropertiesARMGenerator gopter.Gen

// SqlStoredProcedureCreateUpdatePropertiesARMGenerator returns a generator of SqlStoredProcedureCreateUpdatePropertiesARM instances for property testing.
func SqlStoredProcedureCreateUpdatePropertiesARMGenerator() gopter.Gen {
	if sqlStoredProcedureCreateUpdatePropertiesARMGenerator != nil {
		return sqlStoredProcedureCreateUpdatePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlStoredProcedureCreateUpdatePropertiesARM(generators)
	sqlStoredProcedureCreateUpdatePropertiesARMGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureCreateUpdatePropertiesARM{}), generators)

	return sqlStoredProcedureCreateUpdatePropertiesARMGenerator
}

// AddRelatedPropertyGeneratorsForSqlStoredProcedureCreateUpdatePropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlStoredProcedureCreateUpdatePropertiesARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsARMGenerator())
	gens["Resource"] = SqlStoredProcedureResourceARMGenerator()
}

func Test_SqlStoredProcedureResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlStoredProcedureResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlStoredProcedureResourceARM, SqlStoredProcedureResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlStoredProcedureResourceARM runs a test to see if a specific instance of SqlStoredProcedureResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlStoredProcedureResourceARM(subject SqlStoredProcedureResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlStoredProcedureResourceARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SqlStoredProcedureResourceARM instances for property testing - lazily instantiated by
//SqlStoredProcedureResourceARMGenerator()
var sqlStoredProcedureResourceARMGenerator gopter.Gen

// SqlStoredProcedureResourceARMGenerator returns a generator of SqlStoredProcedureResourceARM instances for property testing.
func SqlStoredProcedureResourceARMGenerator() gopter.Gen {
	if sqlStoredProcedureResourceARMGenerator != nil {
		return sqlStoredProcedureResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlStoredProcedureResourceARM(generators)
	sqlStoredProcedureResourceARMGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureResourceARM{}), generators)

	return sqlStoredProcedureResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlStoredProcedureResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlStoredProcedureResourceARM(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.AlphaString()
}
