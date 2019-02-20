package openapiv2

type Spec struct {
	Swagger string `json:"swagger"`
	Info *Info `json:"info"`
	Host string `json:"host"`
	BasePath string `json:"basePath,omitempty"`
	Schemes []string `json:"schemes"`
	Consumes []string `json:"consumes"`
	Produces []string `json:"produces"`
	Paths map[string]Path `json:"paths"`
	Definitions map[string]*Definition `json:"definitions"`
	SecurityDefinitions map[string]*SecurityDefinition `json:"securityDefinitions,omitempty"`
}

// WARNING
// Operation does not require an API key
// callers may invoke the method without specifying an associated API-consuming project.
// To enable API key all the SecurityRequirement Objects (https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#security-requirement-object)
// inside security definition must reference at least one SecurityDefinition of type : 'apiKey'.

type SecurityDefinition struct {
	Type string `json:"type"`
	In string `json:"in"`
	Name string `json:"name"`
	Flow string `json:"flow,omitempty"`
	TokenUrl string `json:"tokenUrl,omitempty"`
}


type Info struct {
	Version string `json:"version"`
	Title string `json:"title"`
	Description string `json:"description"`
	TermsOfService string `json:"termsOfService"`
	Contact struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		URL   string `json:"url"`
	} `json:"contact"`
	License struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"license"`
}

type Path map[string]*PathMethod

type PathMethod struct {
	Description string `json:"description,omitempty"`
	OperationID string `json:"operationId,omitempty"`
	Parameters []*Parameter `json:"parameters,omitempty"`
	Responses map[int]*Response `json:"responses"`
	Produces []string `json:"produces,omitempty"`
//	Security []*SecuritySchemeObject `json:"security,omitempty"`
}

func (self *PathMethod) Response(code int) *Response {
	if self.Responses[code] == nil {
		self.Responses[code] = &Response{
			Headers: map[string]*Header{},
		}
	}
	return self.Responses[code]
}

type Parameter struct {
	// required 'fixed' fields
	Name string `json:"name,omitempty"`
	// options: header, formData, query, path
	In string `json:"in,omitempty"`
	Description string `json:"description,omitempty"`
	Required bool `json:"required,omitempty"`
	// if body
	Schema *Schema `json:"schema,omitempty"`
	// else all of the below
	Type string `json:"type,omitempty"`
	// https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#dataTypeFormat
	Format string `json:"format,omitempty"`
	AllowEmptyValue bool `json:"allowEmptyValue,omitempty"`
	Items map[string]string `json:"items,omitempty"`
	CollectionFormat string `json:"collectionFormat,omitempty"`
//
	Default interface{} `json:"default,omitempty"`
	// String validations
	MaxLength *int64 `json:"maxLength,omitempty"`
	MinLength *int64 `json:"minLength,omitempty"`
	Pattern   string `json:"pattern,omitempty"`
	// Number validations
	Minimum *float64 `json:"minimum,omitempty"`
	Maximum *float64 `json:"maximum,omitempty"`
	ExclusiveMinimum bool `json:"exclusiveMinimum,omitempty"`
	ExclusiveMaximum bool `json:"exclusiveMaximum,omitempty"`
	MultipleOf *float64 `json:"multipleOf,omitempty"`
	Enum []interface{} `json:"enum,omitempty"`
	// slice validations
	MinProperties int `json:"minProperties,omitempty"`
	MaxProperties int `json:"maxProperties,omitempty"`
	MinItems *int64 `json:"minItems,omitempty"`
	MaxItems *int64 `json:"maxItems,omitempty"`
	UniqueItems bool `json:"uniqueItems,omitempty"`
	// "string", "number", "integer", "boolean", "array" or "file". If type is "file" see docs
}

// Schema is taken from https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#schemaObject
// bits stolen from https://github.com/go-swagger/go-swagger/blob/master/generator/structs.go
type Schema struct {
	Ref string `json:"$ref,omitempty"`
	Type string `json:"type,omitempty"`
	// misc
	Format string `json:"format,omitempty"`
	Required bool `json:"required,omitempty"`
	Title string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	// String validations
	MinLength *int64 `json:"minLength,omitempty"`
	MaxLength *int64 `json:"maxLength,omitempty"`
	Pattern   string `json:"pattern,omitempty"`
	// Number validations
	MultipleOf *float64 `json:"multipleOf,omitempty"`
	Minimum *float64 `json:"minimum,omitempty"`
	Maximum *float64 `json:"maximum,omitempty"`
	ExclusiveMinimum bool `json:"exclusiveMinimum,omitempty"`
	ExclusiveMaximum bool `json:"exclusiveMaximum,omitempty"`
	Enum []interface{} `json:"enum,omitempty"`
	// slice validations
	Default interface{} `json:"default,omitempty"`
	MinProperties int `json:"minProperties,omitempty"`
	MaxProperties int `json:"maxProperties,omitempty"`
	MinItems *int64 `json:"minItems,omitempty"`
	MaxItems *int64 `json:"maxItems,omitempty"`
	UniqueItems bool `json:"uniqueItems,omitempty"`
	// "string", "number", "integer", "boolean", or "array" etc
	Items *Items `json:"items,omitempty"`
}

type Items struct {
	// misc
	// "string", "number", "integer", "boolean", or "array" etc
	Type string `json:"type,omitempty"`
	Format string `json:"format,omitempty"`
	// String validations
	MaxLength *int64 `json:"maxLength,omitempty"`
	MinLength *int64 `json:"minLength,omitempty"`
	Pattern   string `json:"pattern,omitempty"`
	// Number validations
	MultipleOf *float64 `json:"multipleOf,omitempty"`
	Minimum *float64 `json:"minimum,omitempty"`
	Maximum *float64 `json:"maximum,omitempty"`
	ExclusiveMinimum bool `json:"exclusiveMinimum,omitempty"`
	ExclusiveMaximum bool `json:"exclusiveMaximum,omitempty"`
	Enum []interface{} `json:"enum,omitempty"`
	// slice validations
	Default interface{} `json:"default,omitempty"`
	MinProperties int `json:"minProperties,omitempty"`
	MaxProperties int `json:"maxProperties,omitempty"`
	MinItems *int64 `json:"minItems,omitempty"`
	MaxItems *int64 `json:"maxItems,omitempty"`
	UniqueItems bool `json:"uniqueItems,omitempty"`
	Items *Items `json:"items,omitempty"`
}

type Definition struct {
	Type  string `json:"type"`
	Ref string   `json:"$ref,omitempty"`
	Required []string `json:"required,omitempty"`
	Properties map[string]Parameter `json:"properties,omitempty"`
}

type StatusSchema struct {
	Type string `json:"type"`
	Items map[string]string `json:"items,omitempty"`
}

type Response struct {
	Description string `json:"description"`
	Schema *StatusSchema `json:"schema,omitempty"`
	Headers map[string]*Header `json:"headers,omitempty"`
}

type Header struct {
	Type string `json:"type,omitempty"`
	Format string `json:"format,omitempty"`
	Description string `json:"description,omitempty"`
	Default interface{} `json:"default,omitempty"`
}
