/*
Copyright 2021 The Kubernetes Authors.

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

package conf

import (
	"testing"

	"k8s.io/client-go/rest"
)

func TestConfig_UseRESTConfig(t *testing.T) {
	cfg := New().WithConfig(&rest.Config{})
	if cfg.kubecfg == nil {
		t.Errorf("kubecfg is nil")
	}
}

func TestConfig_UseNamespace(t *testing.T) {
	cfg := New().WithNamespace("testns")

	if cfg.namespace != "testns"{
		t.Errorf("unexpected config namespace: %s", cfg.namespace)
	}
}