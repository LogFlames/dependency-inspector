/*

Copyright (c) nexB Inc. and others. All rights reserved.
ScanCode is a trademark of nexB Inc.
SPDX-License-Identifier: Apache-2.0
See http://www.apache.org/licenses/LICENSE-2.0 for the license text.
See https://github.com/nexB/dependency-inspector for support or download.
See https://aboutcode.org for more information about nexB OSS projects.

*/

package cmd

import (
	"github.com/nexB/dependency-inspector/internal"
	"github.com/spf13/cobra"
)

func yarnCmd() *cobra.Command {
	lockFiles := []string{"yarn.lock"}
	lockGenCommand := []string{"yarn", "install", "--mode", "update-lockfile"}

	pnpmCmd := &cobra.Command{
		Use:   "yarn [path]",
		Short: "Generate lockfile for yarn project",
		Long: `Create lockfile for yarn project if it doesn't exist in the specified [path].
If no path is provided, the command defaults to the current directory.`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			internal.CreateLockFile(
				lockFiles,
				args,
				lockGenCommand,
			)
		},
	}

	return pnpmCmd
}
