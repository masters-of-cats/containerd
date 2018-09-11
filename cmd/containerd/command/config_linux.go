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

package command

import (
	"github.com/containerd/containerd/defaults"
	"github.com/containerd/containerd/services/server"
)

// Note: We weren't sure of the reason for having separate Default and User config things
// and then setting them to the same thing in defaults/defaults.go
// We have done this for now but there is probably no need for the User* ones
// You will see that we have switched all calls of defaults.Default* to defaults.User*
func defaultConfig() *server.Config {
	c := &server.Config{
		Root:  defaults.UserRootDir,
		State: defaults.UserStateDir,
		Debug: server.Debug{
			Address: defaults.UserDebugAddress,
		},
		GRPC: server.GRPCConfig{
			Address:        defaults.UserAddress,
			MaxRecvMsgSize: defaults.DefaultMaxRecvMsgSize,
			MaxSendMsgSize: defaults.DefaultMaxSendMsgSize,
		},
	}
	return c
}
