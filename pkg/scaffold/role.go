// Copyright 2018 The Operator-SDK Authors
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

package scaffold

import (
	"path/filepath"

	"github.com/operator-framework/operator-sdk/pkg/scaffold/input"
)

const RoleYamlFile = "role.yaml"

type Role struct {
	input.Input
}

func (s *Role) GetInput() (input.Input, error) {
	if s.Path == "" {
		s.Path = filepath.Join(DeployDir, RoleYamlFile)
	}
	s.TemplateBody = roleTemplate
	return s.Input, nil
}

const roleTemplate = `kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: {{.ProjectName}}
rules:
- apiGroups:
  - ""
  resources:
  - pods
  - services
  - endpoints
  - persistentvolumeclaims
  - events
  - configmaps
  - secrets
  verbs:
  - "*"
- apiGroups:
  - apps
  resources:
  - deployments
  - daemonsets
  - replicasets
  - statefulsets
  verbs:
  - "*"
- apiGroups:
  - monitoring.coreos.com
  resources:
  - servicemonitors
  verbs:
  - "get"
  - "create"
`
