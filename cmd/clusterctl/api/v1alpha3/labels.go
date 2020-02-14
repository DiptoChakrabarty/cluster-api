/*
Copyright 2019 The Kubernetes Authors.

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

package v1alpha3

const (
	// ClusterctlLabelName defines the label that is applied to all the components managed by clusterctl.
	ClusterctlLabelName = "clusterctl.cluster.x-k8s.io"

	// ClusterctlCoreLabelName defines the label that is applied to all the core objects managed by clusterctl.
	ClusterctlCoreLabelName = "clusterctl.cluster.x-k8s.io/core"

	// ClusterctlSharedResourceLabelName defines the label that is applied to all the objects that are shared between
	// instances of the same provider. e.g. CRDs, ValidatingWebhookConfiguration, MutatingWebhookConfiguration etc.
	ClusterctlSharedResourceLabelName = "clusterctl.cluster.x-k8s.io/shared"
)
