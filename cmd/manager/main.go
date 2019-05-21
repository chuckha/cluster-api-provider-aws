/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"net/http"

	"sigs.k8s.io/cluster-api-provider-aws/pkg/cloud/aws/webserver/cluster"
	"sigs.k8s.io/cluster-api-provider-aws/pkg/cloud/aws/webserver/machine"
)

func main() {
	clusterServer, err := cluster.NewServer()
	if err != nil {
		panic(err)
	}

	machineServer, err := machine.NewServer()
	if err != nil {
		panic(err)
	}

	fmt.Println("Serving machine service on :8000")
	go http.ListenAndServe(":8000", machineServer)
	fmt.Println("serving cluster service on :8001")
	http.ListenAndServe(":8001", clusterServer)
}
