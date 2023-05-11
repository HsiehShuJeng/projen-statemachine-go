package projenstatemachineexample


type StateMachineApiGatewayExampleProps struct {
	// The path part for the resource.
	PartPath *string `field:"required" json:"partPath" yaml:"partPath"`
	// A stage name for the rest api.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
}

