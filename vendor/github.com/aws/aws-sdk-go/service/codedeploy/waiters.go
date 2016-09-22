// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package codedeploy

import (
	"github.com/aws/aws-sdk-go/private/waiter"
)

func (c *CodeDeploy) WaitUntilDeploymentSuccessful(input *GetDeploymentInput) error {
	waiterCfg := waiter.Config{
		Operation:   "GetDeployment",
		Delay:       15,
		MaxAttempts: 120,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "path",
				Argument: "deploymentInfo.status",
				Expected: "Succeeded",
			},
			{
				State:    "failure",
				Matcher:  "path",
				Argument: "deploymentInfo.status",
				Expected: "Failed",
			},
			{
				State:    "failure",
				Matcher:  "path",
				Argument: "deploymentInfo.status",
				Expected: "Stopped",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}