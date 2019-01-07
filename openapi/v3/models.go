package openapiv3

import {
	"reflect"
}

type Spec struct {
	OpenAPI string `json:"openapi"`
	Info *Info `json:"info"`
	Servers []*Server
	//Host string `json:"host"`
	//BasePath string `json:"basePath"`
	//Schemes []string `json:"schemes"`
	//Consumes []string `json:"consumes"`
	//Produces []string `json:"produces"`
	Paths map[string]*Path `json:"paths"`
	//Definitions map[string]*Definition `json:"definitions"`
	Components *Components `json:"components,omitempty"`
	Security []*SecurityRequirementObject `json:"security"`
	Tags []*TagObject `json:"tags"`
	ExternalDocs ExternalDocumentationObject `json:"externalDocs"`
}

type Server struct {
	Url string `json:"url"`
	Description string `json:"description"`
	Variables map[string]*ServerVariableObject `json:"variables"`
}

type ServerVariableObject struct {
	Enum []string `json:"enum"`
	Default string `json:"default"`
	Description string `json:"description"`
}

type ExternalDocumentationObject struct {
	Description string `json:"description"`
	Url string `json:"url"`

}

type TagObject struct {
	Name string `json:"name"`
	Description string `json:"description"`
	ExternalDocs *ExternalDocumentationObject `json:"externalDocumentationObject"`
}

type SecurityRequirementObject struct {
	Schemas map[string]*SchemaObject `json:"schemas"`
	Responses map[string]*Response `json:"responses"`
	Parameters map[string]*ParameterObject `json:"parameters"`
	Examples map[string]*ExampleObject `json:"examples"`
	RequestBodies map[string]*RequestBodyObject `json:"requestBodies"`
	Headers map[string]*HeaderObject `json:"headers"`
	SecurityRequirements map[string]string `json:"securityRequirements"`
	Links map[string]*LinkObject `json:"links"`
	Callbacks map[string]*CallbackObject `json:"callbacks"`

}

type Components struct {
	SecuritySchemes map[string]*SecuritySchemeObject `json:"securitySchemes"`
}

type SecuritySchemeObject struct {
	Type string `json:"type"`
	Description string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
	In string `json:"in,omitempty"`
	Scheme string `json:"scheme,omitempty"`
	BearerFormat string `json:"bearerFormat,omitempty"`
	// oauth
	Flows OAuthFlowsObject `json:"flows,omitempty"` // "implicit", "password", "application" or "accessCode"
	OpenIdConnectUrl string `json:"OpenIdConnectUrl,omitempty"` 
}

type OAuthFlowsObject struct {
	Implicit *OAuthFlowObject `json:"implicit"` 
	Password *OAuthFlowObject `json:"password"`
	ClientCredentials *OAuthFlowObject  `json:"clientCredentials"`
	AuthorizationCode *OAuthFlowObject `json:"authorizationCode"`
}

type OAuthFlowObject struct {
	AuthorizationUrl string `json:"authorizationUrl"`
	TokenUrl string `json:"tokenUrl"`
	RefreshUrl string `json:"refreshUrl"`
	Scopes map[string]string `json:"scopes"`
}

type Path map[string]*PathItemObject


type PathItemObject struct {
	Ref *PathItemObject `json:"$ref"`
	Summary string `json:"summary"`
	Descriptiom string `json:"description"`
	Get *OperationObject `json:"get"`
	Put *OperationObject `json:"put"`
	Post *OperationObject `json:"post"`
	Delete *OperationObject `json:"delete"`
	Options *OperationObject `json:"options"`
	Head *OperationObject `json:"head"`
	Path *OperationObject `json:"path"`
	Trace *OperationObject `json:"trace"`
	Servers []*Server  `json:"servers"`
	Parameters []*ParameterObject `json:"parameters"`
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

type Parameter struct {
	// required 'fixed' fields
	Name string `json:"name,omitempty"`
	// options: header, formData, query, path
	In string `json:"in,omitempty"`
	Description string `json:"description,omitempty"`
	Required bool `json:"required,omitempty"`
	Depreciated bool `json:"depreciated,omitempty"`
	AllowEmptyvalue bool `json:"allowEmptyValue,omitempty"`

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

type StatusCode struct {
	Description string `json:"description"`
	Schema StatusSchema `json:"schema"`
}

// Comments refer to the line above the comment :)
type Responses struct {
	Responses map[int]*Response{}  `json:"responses"`
}

type Response struct {
	Description String `json:"description"`
	Headers map[string]*Header `json:"headers"`
	Content map[string]*MediaTypeObject `json:"content"`
	Links map[string]*LinkObject `json:"links"`
}

type Header struct {
	Type string `json:"type,omitempty"`
	Format string `json:"format,omitempty"`
	Description string `json:"description,omitempty"`
	Default interface{} `json:"default,omitempty"`
}
type MediaTypeObject struct {
	Schema  *SchemaObject `json:"schema"`
	Example interface{} `json:"example"`
	Examples map[string]interface{} `json:"examples"`
	Encoding map[string]*EncodingObject `json:"encoding"`
}

type LinkObject struct {
	OperationRef string `json:"operationRef"`
	OperationID string `json:"operationId"`
	Parameters map[string]interface{} `json:"parameters"`
	RequestBody interface{} `json:"requestBody"`
	Description string `json:"description"`
	Server *Server `json:"server"`

}

type SchemaObject struct {
	Title string `json:"title"`
	MultipleOf interface{} `json:"multipleOf"`
	Maximum interface{} `json:"maximum"`
	ExclusiveMaximum bool `json:"exclusiveMaximum"`
	Minimum interface{} `json:"minimum"`
	ExclusiveMinimum bool `json:"exclusiveMinimum"`
	MaxLength interface{} `json:"maxLength"`
	MinLength interface{} `json:"minLength"`
	Pattern string `json:"pattern"`
	MaxItems int64 `json:"maxItems"`
	MinItems int64 `json:"minItems"`
	UniqueItems interface{} `json:"uniqueItems"`
	MaxProperties int `json:"maxProperties"`
	MinProperties int `json:"minProperties"`
	Required bool `json:"required"`
	Enum interface{} `json:"enum"`

	Type string `json:"type"`
	AllOf *SchemaObject `json:"allOf"`
	OneOf *SchemaObject `json:"oneOf"`
 	Not *SchemaObject `json:"not"`
	Items *SchemaObject `json:"items"`
	Properties *SchemaObject `json:"properties"`
	AdditionalProperties *SchemaObject `json:"AdditionalProperties"`
	Description string `json:"description"`
	Format string `json:"format"`
	Default interface{} `json:"default"`

}
