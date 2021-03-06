// Copyright 2017 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package policygen

var policiesTestSuite = PolicyTestSuite{
	l3Checks: []PolicyTestKind{
		{
			name:  "No Policy",
			kind:  ingress,
			tests: ConnResultAllOK,
			template: map[string]string{
				"FromEndpoints": `[{}]`,
			},
		},
		{
			name:  "Ingress Label",
			kind:  ingress,
			tests: ConnResultAllOK,
			template: map[string]string{
				"FromEndpoints": `[{"matchLabels": { "id": "{{.SrcPod}}"}}]`,
			},
		},
		{
			name:  "Ingress Label Invalid",
			kind:  ingress,
			tests: ConnResultAllTimeout,
			template: map[string]string{
				"FromEndpoints": `[{"matchLabels": { "id": "{{.SrcPod}}Invalid"}}]`,
			},
		},
	},
	l4Checks: []PolicyTestKind{
		{
			name:     "No Policy",
			kind:     ingress,
			tests:    ConnResultAllOK,
			template: map[string]string{},
		},
		{
			name:  "Ingress Port 80 No protocol",
			kind:  ingress,
			tests: ConnResultOnlyHTTP,
			template: map[string]string{
				"Ports": `[{"port": "80"}]`,
			},
		},
		{
			name:  "Ingress Port 80 TCP",
			kind:  ingress,
			tests: ConnResultOnlyHTTP,
			template: map[string]string{
				"Ports": `[{"port": "80", "protocol": "TCP"}]`,
			},
		},
		{
			name:  "Ingress Port 80 UDP",
			kind:  ingress,
			tests: ConnResultAllTimeout,
			template: map[string]string{
				"Ports": `[{"port": "80", "protocol": "UDP"}]`,
			},
		},
	},
	l7Checks: []PolicyTestKind{
		{
			name:     "No Policy",
			kind:     ingress,
			tests:    ConnResultAllOK,
			template: map[string]string{},
		},
		{
			name:  "Ingress policy /private/",
			kind:  ingress,
			tests: ConnResultOnlyHTTPPrivate,
			template: map[string]string{
				"Rules": `{"http": [{"method": "GET", "path": "/private"}]}`,
				"Ports": `[{"port": "80", "protocol": "TCP"}]`,
			},
		},
		{
			name:  "Egress policy to /private/",
			kind:  egress,
			tests: ConnResultOnlyHTTPPrivate,
			template: map[string]string{
				"Rules": `{"http": [{"method": "GET", "path": "/private"}]}`,
				"Ports": `[{"port": "80", "protocol": "TCP"}]`,
			},
		},
	},
}

// GeneratedTestSpec returns a `TestSpec` array with all the policies
// possibilities based on all combinations of `policiesTestSuite`
func GeneratedTestSpec() []TestSpec {
	var testSpecs = []TestSpec{}
	for _, l3 := range policiesTestSuite.l3Checks {
		for _, l4 := range policiesTestSuite.l4Checks {
			for _, l7 := range policiesTestSuite.l7Checks {
				for _, dst := range DestinationsTypes {
					testSpecs = append(testSpecs, TestSpec{
						l3:          l3,
						l4:          l4,
						l7:          l7,
						Destination: dst,
					})
				}
			}
		}
	}
	return testSpecs
}
