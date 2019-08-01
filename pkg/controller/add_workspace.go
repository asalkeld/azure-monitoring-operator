package controller

import (
	"github.com/openshift/azure-monitoring-operator/pkg/controller/workspace"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, workspace.Add)
}
