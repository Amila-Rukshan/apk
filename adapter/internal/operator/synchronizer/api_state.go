/*
 *  Copyright (c) 2022, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package synchronizer

import (
	"github.com/wso2/apk/adapter/internal/operator/apis/dp/v1alpha1"
	gwapiv1b1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

// APIState holds the state of the deployed APIs. This state is compared with
// the state of the Kubernetes controller cache to detect updates.
// +k8s:deepcopy-gen=true
type APIState struct {
	APIDefinition *v1alpha1.API
	ProdHTTPRoute *HTTPRouteState
	SandHTTPRoute *HTTPRouteState
}

// HTTPRouteState holds the state of the deployed httpRoutes. This state is compared with
// the state of the Kubernetes controller cache to detect updates.
// +k8s:deepcopy-gen=true
type HTTPRouteState struct {
	HTTPRoute                 *gwapiv1b1.HTTPRoute
	Authentications           map[string]v1alpha1.Authentication
	ResourceAuthentications   map[string]v1alpha1.Authentication
	APIPolicies               map[string]v1alpha1.APIPolicy
	ResourceAPIPolicies       map[string]v1alpha1.APIPolicy
	InterceptorServiceMapping map[string]v1alpha1.InterceptorService
	BackendMapping            v1alpha1.BackendMapping
	Scopes                    map[string]v1alpha1.Scope
	RateLimitPolicies         map[string]v1alpha1.RateLimitPolicy
	ResourceRateLimitPolicies map[string]v1alpha1.RateLimitPolicy
}
