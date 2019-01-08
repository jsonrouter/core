package openapiv3

import (
	//"reflect"
)

type Spec struct {
	OpenAPI string `json:"openapi"`
	Info *Info `json:"info"`
	Servers []*Server
	//Host string `json:"host"`
	//BasePath string `json:"basePath"`
	//Schemes []string `json:"schemes"`
	//Consumes []string `json:"consumes"`
	//Produces []string `json:"produces"`
	Paths map[string]Path `json:"paths"`
	//Definitions map[string]*Definition `json:"definitions"`
	Components *Components `json:"components,omitempty"`
	Security []*SecurityRequirement `json:"security"`
	Tags []*Tag `json:"tags"`
	ExternalDocs ExternalDocumentation `json:"externalDocs"`
}

type Server struct {
	Url string `json:"url"`
	Description string `json:"description"`
	Variables map[string]*ServerVariable `json:"variables"`
}

type ServerVariable struct {
	Enum []string `json:"enum"`
	Default string `json:"default"`
	Description string `json:"description"`
}

type ExternalDocumentation struct {
	Description string `json:"description"`
	Url string `json:"url"`

}

type Tag struct {
	Name string `json:"name"`
	Description string `json:"description"`
	ExternalDocs *ExternalDocumentation `json:"ExternalDocumentation"`
}

type SecurityRequirement struct {
	Schemas map[string]*Schema `json:"schemas"`
	Responses map[string]*Response `json:"responses"`
	Parameters map[string]*Parameter `json:"parameters"`
	Examples map[string]*Example `json:"examples"`
	RequestBodies map[string]*RequestBody `json:"requestBodies"`
	Headers map[string]*Header `json:"headers"`
	SecurityRequirements map[string]string `json:"securityRequirements"`
	Links map[string]*Link `json:"links"`
	Callbacks map[string]*CallBack `json:"callbacks"`

}

type Components struct {
	Schemas map[string]*Schema `json:"schemas"`
	Responses map[string]*Response `json:"responses"`
	Parameters map[string]*Parameter `json:"parameters"`
	Examples map[string]*Example `json:"examples"`
	SecuritySchemes map[string]*SecurityScheme `json:"securitySchemes"`
	RequestBodies map[string]*RequestBody `json:"RequestBodies"`
	Headers map[string]*Header `json:"headers"`
	Links map[string]*Link `json:"links"`
	Callbacks map[string]*CallBack `json:"callbacks"`
}

type SecurityScheme struct {
	Type string `json:"type"`
	Description string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
	In string `json:"in,omitempty"`
	Scheme string `json:"scheme,omitempty"`
	BearerFormat string `json:"bearerFormat,omitempty"`
	// oauth
	Flows OAuthFlows `json:"flows,omitempty"` // "implicit", "password", "application" or "accessCode"
	OpenIdConnectUrl string `json:"OpenIdConnectUrl,omitempty"` 
}

type OAuthFlows struct {
	Implicit *OAuthFlow `json:"implicit"` 
	Password *OAuthFlow `json:"password"`
	ClientCredentials *OAuthFlow  `json:"clientCredentials"`
	AuthorizationCode *OAuthFlow `json:"authorizationCode"`
}

type OAuthFlow struct {
	AuthorizationUrl string `json:"authorizationUrl"`
	TokenUrl string `json:"tokenUrl"`
	RefreshUrl string `json:"refreshUrl"`
	Scopes map[string]string `json:"scopes"`
}

type Path map[string]*PathItem


type PathItem struct {
	Ref *PathItem `json:"$ref"`
	Summary string `json:"summary"`
	Descriptiom string `json:"description"`
	Get *Operation `json:"get"`
	Put *Operation `json:"put"`
	Post *Operation `json:"post"`
	Delete *Operation `json:"delete"`
	Options *Operation `json:"options"`
	Head *Operation `json:"head"`
	Path *Operation `json:"path"`
	Trace *Operation `json:"trace"`
	Servers []*Server  `json:"servers"`
	Parameters []*Parameter `json:"parameters"`
}

type Operation struct {
	Tags []string
	Summary string
	Description string
	ExternalDocs *ExternalDocumentation
	OperationId string
	Parameters []*Parameter
	RequestBody *RequestBody 
	Responses *Responses
	Callbacks map[string]*CallBack
	Depreciated bool
	Security []SecurityRequirement 
	Servers []ServerVariable

}

type Encoding struct {
	ContentType string `json:"contentType"`
	Headers map[string]*Header `json:"headers"`
	Style string `json:"style"`
	Explode bool `json:"explode"`
	AllowReserved bool `json:"allowReserved"`

}

type Example struct {
	Summary string `json:"summary"`
	Description string `json:"description"`
	Value interface{} `json:"value"`
 	ExternalValue string `json:"externalValue"`
}

type RequestBody struct {
	Description string `json:"description"`
	Content map[string]*MediaType `json:"contents"`
	Required bool `json:"required"`
}

type Header struct {
	Name string `json:"name"`
	In string `json:"in"`
	Description string `json:"description,omitempty"`
	Required bool `json:"required,omitempty"`
	Depreciated bool `json:"depreciated,omitempty"`
	AllowEmptyvalue bool `json:"allowEmptyValue,omitempty"`
}

type CallBack struct {
	Expression map[string]*PathItem 
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
	Responses map[int]*Response  `json:"responses"`
}

type Response struct {
	Description string `json:"description"`
	Headers map[string]*Header `json:"headers"`
	Content map[string]*MediaType `json:"content"`
	Links map[string]*Link `json:"links"`
}

type MediaType struct {
	Schema  *Schema `json:"schema"`
	Example interface{} `json:"example"`
	Examples map[string]interface{} `json:"examples"`
	Encoding map[string]*Encoding `json:"encoding"`
}

type Link struct {
	OperationRef string `json:"operationRef"`
	OperationID string `json:"operationId"`
	Parameters map[string]interface{} `json:"parameters"`
	RequestBody interface{} `json:"requestBody"`
	Description string `json:"description"`
	Server *Server `json:"server"`

}

type Schema struct {
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
	AllOf *Schema `json:"allOf"`
	OneOf *Schema `json:"oneOf"`
 	Not *Schema `json:"not"`
	Items *Schema `json:"items"`
	Properties *Schema `json:"properties"`
	AdditionalProperties *Schema `json:"AdditionalProperties"`
	Description string `json:"description"`
	Format string `json:"format"`
	Default interface{} `json:"default"`

}
