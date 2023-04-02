// An example construct for deploying to npm, PyPi, Maven, and Nuget with Amazon API Gateway and AWS Step Functions.
package projenstatemachineexample


type StateMachineApiGatewayExampleProps struct {
	// The path part for the resource.
	PartPath *string `field:"required" json:"partPath" yaml:"partPath"`
	// A stage name for the rest api.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
}

