// Copyright 2020 Datawire.  All rights reserved
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

///////////////////////////////////////////////////////////////////////////
// Important: Run "make update-yaml" to regenerate code after modifying
// this file.
///////////////////////////////////////////////////////////////////////////

package v2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RateLimitServiceSpec defines the desired state of RateLimitService
type RateLimitServiceSpec struct {
	// Common to all Ambassador objects.
	AmbassadorID AmbassadorID `json:"ambassador_id,omitempty"`

	// +kubebuilder:validation:Required
	Service string               `json:"service,omitempty"`
	Timeout *MillisecondDuration `json:"timeout_ms,omitempty"`
	Domain  string               `json:"domain,omitempty"`
	// +k8s:conversion-gen=false
	TLS *BoolOrString `json:"tls,omitempty"`
	// +kubebuilder:validation:Enum={"v2","v3"}
	ProtocolVersion string `json:"protocol_version,omitempty"`

	// +k8s:conversion-gen:rename=StatsName
	V3StatsName string `json:"v3StatsName,omitempty"`

	// +k8s:conversion-gen:rename=FailureModeDeny
	//
	// FailureModeDeny when set to true, envoy will deny traffic if it
	// is unable to communicate with the rate limit service.
	V3FailureModeDeny bool `json:"failure_mode_deny,omitempty"`
}

// RateLimitService is the Schema for the ratelimitservices API
//
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
type RateLimitService struct {
	metav1.TypeMeta   `json:""`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec RateLimitServiceSpec `json:"spec,omitempty"`
}

// RateLimitServiceList contains a list of RateLimitServices.
//
// +kubebuilder:object:root=true
type RateLimitServiceList struct {
	metav1.TypeMeta `json:""`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RateLimitService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RateLimitService{}, &RateLimitServiceList{})
}
