// Copyright 2020 program was created by VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"log"

	framework "github.com/tbe/resource-framework/resource"
	"github.com/vmware/xcenter-vrealize-automation-resource/pkg/resource"
)

func main() {
	// Create new vRA Resource
	res := &resource.VRAResource{
		Src:       &resource.VRASource{},
		Ver:       &resource.VRAVersion{},
		OutParams: &resource.OutParams{},
	}

	// Register the handler
	handler, err := framework.NewHandler(res)
	if err != nil {
		log.Fatal("Error while registering the handler", err.Error())
	}

	// Run
	err = handler.Run()
	if err != nil {
		log.Fatal("Error while running:", err.Error())
	}
}
