package openapiv3

func NewV3(host, serviceName string) *Spec {
	return &Spec{
		OpenAPI: "3.0.0",
		Info: &Info{
			Title: serviceName,
		},
		Host: host,
		Paths: map[string]*Path{},
		Schemes: []string{"http"},
		Consumes: []string{"application/json"},
		Produces: []string{"application/json"},
		Definitions: map[string]*Definition{},
		Components: &Components{
			SecuritySchemes: map[string]*SecuritySchemeObject{},
		},
	}
}
