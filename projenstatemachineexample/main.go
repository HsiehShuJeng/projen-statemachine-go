// An example construct for deploying to npm, PyPi, Maven, and Nuget with Amazon API Gateway and AWS Step Functions.
package projenstatemachineexample

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"projen-statemachine-example.StateMachineApiGatewayExample",
		reflect.TypeOf((*StateMachineApiGatewayExample)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "executionInput", GoGetter: "ExecutionInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stateMachine", GoGetter: "StateMachine"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StateMachineApiGatewayExample{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"projen-statemachine-example.StateMachineApiGatewayExampleProps",
		reflect.TypeOf((*StateMachineApiGatewayExampleProps)(nil)).Elem(),
	)
}
