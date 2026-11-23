/*
Copyright 2026 The Kubernetes Authors.

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

package controller

import (
	"testing"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func TestExternalNamePorts(t *testing.T) {
	tests := []struct {
		name       string
		service    *apiv1.Service
		wantPort   int32
		wantTarget int
	}{
		{
			name: "8080",
			service: &apiv1.Service{Spec: apiv1.ServiceSpec{Ports: []apiv1.ServicePort{
				{Port: 8080},
			}}},
			wantPort:   8080,
			wantTarget: 0,
		},
		{
			name: "https",
			service: &apiv1.Service{Spec: apiv1.ServiceSpec{Ports: []apiv1.ServicePort{
				{Name: "https", Port: 443, TargetPort: intstr.FromInt(8443)},
			}}},
			wantPort:   443,
			wantTarget: 8443,
		},
		{
			name:       "2147483648",
			service:    &apiv1.Service{},
			wantPort:   0,
			wantTarget: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := externalNamePorts(tt.name, tt.service)
			if got.Port != tt.wantPort {
				t.Fatalf("port = %d, want %d", got.Port, tt.wantPort)
			}
			if got.TargetPort.IntValue() != tt.wantTarget {
				t.Fatalf("target port = %d, want %d", got.TargetPort.IntValue(), tt.wantTarget)
			}
		})
	}
}
