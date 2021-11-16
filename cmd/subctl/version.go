/*
SPDX-License-Identifier: Apache-2.0

Copyright Contributors to the Submariner project.

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

package subctl

import (
	"github.com/spf13/cobra"
	"github.com/submariner-io/submariner-operator/pkg/version"
	"os"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get version information on subctl",
	Long: `This command shows the version tag, and git commit for your
subctl binary.`,
	Run: subctlVersion,
}

func subctlVersion(cmd *cobra.Command, args []string) {
	version.PrintSubctlVersion(os.Stdout)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}