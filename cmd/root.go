// Copyright 2024 Chris Wheeler

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func runFunc(cmd *cobra.Command, args []string) {
	timeStr := time.Now().Format(validFormats["rfc3339"])
	fmt.Print(timeStr)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "timef",
	Version: "v0.0.6",
	Run:     runFunc,
	Args:    cobra.NoArgs,
	Short:   "Display the current time, set a timer, or start a stopwatch",
	Long: `Display the current time, set a timer, or start a stopwatch.

The default time format is rfc3339.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
