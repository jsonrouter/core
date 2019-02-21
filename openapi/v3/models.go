package openapiv3

import (
	//"reflect"
)

// Spec models an OpenApiv3 spec JSON
type Spec struct {
	OpenAPI string `json:"openapi"`
	Info *Info `json:"info"`
	Servers []*Server `json:"servers,omitempty"`
	//Host string `json:"host"`
	//BasePath string `json:"basePath"`
	//Schemes []string `json:"schemes"`
	//Consumes []string `json:"consumes"`
	//Produces []string `json:"produces"`
	Paths map[string]Path `json:"paths"`
	//Definitions map[string]*Definition `json:"definitions"`
	Components *Components `json:"components,omitempty"`
	Security []*SecurityRequirement `json:"security,omitempty"`
	Tags []*Tag `json:"tags,omitempty"`
	ExternalDocs ExternalDocumentation `json:"externalDocs,omitempty"`
}

// Server models an OpenApiv3 Server Object 
type Server struct {
	Url string `json:"url"`
	Description string `json:"description"`
	Variables map[string]*ServerVariable `json:"variables"`
}

// ServerVariable models an OpenApi-v3 Server Variable Object
type ServerVariable struct {
	Enum []string `json:"enum"`
	Default string `json:"default"`
	Description string `json:"description"`
}

// ExternalDocumentation models an OpenApi-v3 External Documentation Object
type ExternalDocumentation struct {
	Description string `json:"description"`
	Url string `json:"url"`

}

// Tag models an OpenApi-v3 Tag Object 
type Tag struct {
	Name string `json:"name"`
	Description string `json:"description"`
	ExternalDocs *ExternalDocumentation `json:"ExternalDocumentation"`
}

// SecurityRequirement models an OpenApi-v3 Security Requirement Object
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

// Components models an OpenApi-v3 Components Object
type Components struct {
	Schemas map[string]*Schema `json:"schemas,omitempty"`
	Responses map[string]*Response `json:"responses,omitempty"`
	Parameters map[string]*Parameter `json:"parameters,omitempty"`
	Examples map[string]*Example `json:"examples,omitempty"`
	SecuritySchemes map[string]*SecurityScheme `json:"securitySchemes,omitempty"`
	RequestBodies map[string]*RequestBody `json:"requestBodies,omitempty"`
	Headers map[string]*Header `json:"headers,omitempty"`
	Links map[string]*Link `json:"links,omitempty"`
	Callbacks map[string]*CallBack `json:"callbacks,omitempty"`
}

// SecurityScheme models an OpenApi-v3 Security Scheme Object
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

// 0AuthFlows models an OpenApi-v3 0AuthFlows Object
type OAuthFlows struct {
	Implicit *OAuthFlow `json:"implicit"` 
	Password *OAuthFlow `json:"password"`
	ClientCredentials *OAuthFlow  `json:"clientCredentials"`
	AuthorizationCode *OAuthFlow `json:"authorizationCode"`
}
// 0AuthFlow models an OpenApi-v3 0AuthFlow Object
type OAuthFlow struct {
	AuthorizationUrl string `json:"authorizationUrl"`
	TokenUrl string `json:"tokenUrl"`
	RefreshUrl string `json:"refreshUrl"`
	Scopes map[string]string `json:"scopes"`
}

// Path models an OpenApi-v3 Path Object
type Path map[string]*Operation

// PathItem models an OpenApi-v3 Path Item Object
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

// Operation models an OpenApi-v3 Operation Object
type Operation struct {
	Tags []string `json:"tags,omitempty"`
	Summary string `json:"summary,omitempty"`
	Description string `json:"description,omitempty"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`
	OperationId string `json:"operationId,omitempty"`
	Parameters []*Parameter `json:"parameters"`
	RequestBody *RequestBody `json:"requestBody,omitempty"`
	Responses map[int]*Response `json:"responses"`
	Callbacks map[string]*CallBack `json:"callbacks,omitempty"`
	Depreciated bool `json:"depreciated,omitempty"`
	Security []SecurityRequirement `json:"security,omitempty"`
	Servers []ServerVariable `json:"servers,omitempty"`

}

// Response models an OpenApi-v3 Response Object
func (self *Operation) Response(code int) *Response {
	if self.Responses[code] == nil {
		self.Responses[code] = &Response{
			Headers: map[string]*Header{},
		}
	}
	return self.Responses[code]
}

// Encoding models an OpenApi-v3 Encoding Object
type Encoding struct {
	ContentType string `json:"contentType"`
	Headers map[string]*Header `json:"headers"`
	Style string `json:"style"`
	Explode bool `json:"explode"`
	AllowReserved bool `json:"allowReserved"`

}

// Example models an OpenApi-v3 Example Object
type Example struct {
	Summary string `json:"summary"`
	Description string `json:"description"`
	Value map[string]interface{} `json:"value"`
 	ExternalValue string `json:"externalValue"`
}

// RequestBody models an OpenApi-v3 Request Body Object
type RequestBody struct {
	Ref string `json:"$ref"`
	Description string `json:"description,omitempty"`
	Content map[string]*MediaType `json:"content"`
	Required bool `json:"required,omitempty"`
}

// Header models an OpenApi-v3 Header Object
type Header struct {
	Name string `json:"name"`
	In string `json:"in"`
	Description string `json:"description,omitempty"`
	Required bool `json:"required,omitempty"`
	Depreciated bool `json:"depreciated,omitempty"`
	AllowEmptyvalue bool `json:"allowEmptyValue,omitempty"`

	Style string `json:"style,omitempty"`
	Explode bool `json:"example,omitempty"` 
	AllowReserved bool `json:"allowReserved,omitempty"`
	Schema *Schema `json:"schema,omitempty"`
	Example *Example `json:"example,omitempty"`
	Examples map[string]*Example `json:"examples,omitempty"`
}

// Callback models an OpenApi-v3 Callback Object
type CallBack struct {
	Expression map[string]*PathItem 
}

// Info  models an OpenApi-v3 Info Ibject
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
// Parameter models an OpenApi-v3 Parameter Object
type Parameter struct {
	// required 'fixed' fields
	Name string `json:"name"`
	// options: header, formData, query, path
	In string `json:"in"`
	Description string `json:"description,omitempty"`
	Required bool `json:"required,omitempty"`
	Depreciated bool `json:"depreciated,omitempty"`
	AllowEmptyvalue bool `json:"allowEmptyValue,omitempty"`

	Style string `json:"style,omitempty"`
	Explode bool `json:"example,omitempty"` 
	AllowReserved bool `json:"allowReserved,omitempty"`
	Schema *Schema `json:"schema,omitempty"`
	Example *Example `json:"example,omitempty"`
	Examples map[string]*Example `json:"examples,omitempty"`

}

// Items models an OpenApi-v3 Items Object
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

// Definition models an OpenApi-v3 Definition Object
type Definition struct {
	Type  string `json:"type"`
	Ref string   `json:"$ref,omitempty"`
	Required []string `json:"required,omitempty"`
	Properties map[string]Parameter `json:"properties,omitempty"`
}

// StatusSchema models an OpenApi-v3 Status Schema Object
type StatusSchema struct {
	Type string `json:"type"`
	Items map[string]string `json:"items,omitempty"`
}

// StatusCode models an OpenApi-v3 Status Code Object
type StatusCode struct {
	Description string `json:"description"`
	Schema StatusSchema `json:"schema"`
}

// Responses models an OpenApi-v3 Responses Object
type Responses map[int]*Response  

// Response models an OpenApi-v3 Response Object
type Response struct {
	Description string `json:"description"`
	Headers map[string]*Header `json:"headers,omitempty"`
	Content map[string]*MediaType `json:"content,omitempty"`
	Links map[string]*Link `json:"links,omitempty"`
}

// MediaType models an OpenApi-v3 Media Type Object
type MediaType struct {
	Schema  *Schema `json:"schema,omitempty"`
	Example interface{} `json:"example,omitempty"`
	Examples map[string]interface{} `json:"examples,omitempty"`
	Encoding map[string]*Encoding `json:"encoding,omitempty"`
}

// Link models an OpenApi-v3 Link Object
type Link struct {
	OperationRef string `json:"operationRef"`
	OperationID string `json:"operationId"`
	Parameters map[string]interface{} `json:"parameters"`
	RequestBody interface{} `json:"requestBody"`
	Description string `json:"description"`
	Server *Server `json:"server"`

}

// Schema models an OpenApi-v3 Schema Object
type Schema struct {
	Title string `json:"title,omitempty"`
	MultipleOf int `json:"multipleOf,omitempty"`
	Maximum int `json:"maximum,omitempty"`
	ExclusiveMaximum bool `json:"exclusiveMaximum,omitempty"`
	Minimum int `json:"minimum,omitempty"`
	ExclusiveMinimum bool `json:"exclusiveMinimum,omitempty"`
	MaxLength int `json:"maxLength,omitempty"`
	MinLength int `json:"minLength,omitempty"`
	Pattern string `json:"pattern,omitempty"`
	MaxItems int `json:"maxItems,omitempty"`
	MinItems int `json:"minItems,omitempty"`
	UniqueItems interface{} `json:"uniqueItems,omitempty"`
	MaxProperties int `json:"maxProperties,omitempty"`
	MinProperties int `json:"minProperties,omitempty"`
	Required bool `json:"required,omitempty"`
	Enum []string `json:"enum,omitempty"`

	Type string `json:"type,omitempty"`
	AllOf []*Schema `json:"allOf,omitempty"`
	OneOf []*Schema `json:"oneOf,omitempty"`
 	Not *Schema `json:"not,omitempty"`
	Items *Schema `json:"items,omitempty"`
	Properties map[string]*Schema `json:"properties,omitempty"`
	AdditionalProperties *Schema `json:"AdditionalProperties,omitempty"`
	Description string `json:"description,omitempty"`
	Format string `json:"format,omitempty"`
	Default string `json:"default,omitempty"`

}


