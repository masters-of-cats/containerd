/*
   Copyright The containerd Authors.

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

package oci

import (
	"context"

	"github.com/containerd/containerd/containers"
	"github.com/containerd/containerd/rootless/specconv"
)

// WithRootless applies github.com/containerd/containerd/rootless/specconv.ToRootless().
// WithRootless needs to be called in the daemon's user namespace.

// Note: Garden does not need to use this in our client as we are not running with user ns
func WithRootless() SpecOpts {
	return func(_ context.Context, _ Client, _ *containers.Container, s *Spec) error {
		return specconv.ToRootless(s)
	}
}
