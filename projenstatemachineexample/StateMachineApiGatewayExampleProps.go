package projenstatemachineexample


type StateMachineApiGatewayExampleProps struct {
	// The path part for the resource.
	// Default: 'pets'.
	//
	PartPath *string `field:"required" json:"partPath" yaml:"partPath"`
	// A stage name for the rest api.
	// Default: 'default'.
	//
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
}

