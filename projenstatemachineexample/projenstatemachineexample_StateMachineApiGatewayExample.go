// An example construct for deploying to npm, PyPi, Maven, and Nuget with Amazon API Gateway and AWS Step Functions.
package projenstatemachineexample

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/HsiehShuJeng/projen-statemachine-go/projenstatemachineexample/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/HsiehShuJeng/projen-statemachine-go/projenstatemachineexample/v2/internal"
)

// Converted from an AWS Blog post.
//
// It is the first example mentioned in https://aws.amazon.com/tw/blogs/compute/introducing-amazon-api-gateway-service-integration-for-aws-step-functions/.
// This constcut will create an API Gateway Rest API with two methods and
// are manipulated by a state machine managed in AWS StepFucntions.
type StateMachineApiGatewayExample interface {
	constructs.Construct
	// sample input to start execution for the workflow.
	ExecutionInput() *string
	// The tree node.
	Node() constructs.Node
	// the representation of a state machine.
	StateMachine() awsstepfunctions.StateMachine
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for StateMachineApiGatewayExample
type jsiiProxy_StateMachineApiGatewayExample struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_StateMachineApiGatewayExample) ExecutionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StateMachineApiGatewayExample) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StateMachineApiGatewayExample) StateMachine() awsstepfunctions.StateMachine {
	var returns awsstepfunctions.StateMachine
	_jsii_.Get(
		j,
		"stateMachine",
		&returns,
	)
	return returns
}


func NewStateMachineApiGatewayExample(parent constructs.Construct, name *string, props *StateMachineApiGatewayExampleProps) StateMachineApiGatewayExample {
	_init_.Initialize()

	j := jsiiProxy_StateMachineApiGatewayExample{}

	_jsii_.Create(
		"projen-statemachine-example.StateMachineApiGatewayExample",
		[]interface{}{parent, name, props},
		&j,
	)

	return &j
}

func NewStateMachineApiGatewayExample_Override(s StateMachineApiGatewayExample, parent constructs.Construct, name *string, props *StateMachineApiGatewayExampleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"projen-statemachine-example.StateMachineApiGatewayExample",
		[]interface{}{parent, name, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func StateMachineApiGatewayExample_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"projen-statemachine-example.StateMachineApiGatewayExample",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StateMachineApiGatewayExample) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

