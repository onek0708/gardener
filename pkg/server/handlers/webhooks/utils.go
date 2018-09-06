// Copyright (c) 2018 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package webhooks

import (
	"k8s.io/api/admission/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func admissionResponse(allowed bool, msg string) *v1beta1.AdmissionResponse {
	response := &v1beta1.AdmissionResponse{
		Allowed: allowed,
	}

	if msg != "" {
		response.Result = &metav1.Status{
			Message: msg,
		}
	}

	return response
}

func errToAdmissionResponse(err error) *v1beta1.AdmissionResponse {
	return admissionResponse(false, err.Error())
}
