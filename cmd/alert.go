/*
Copyright © 2024 Chris Wheeler <mail@chriswheeler.dev>

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
package cmd

import (
	"github.com/gen2brain/beeep"
	"github.com/spf13/cobra"
)

func alertRunFunc(cmd *cobra.Command, args []string) error {
	seconds := toSeconds(args)
	sleepProgress(seconds)
	return beeep.Alert("Timer", "Time's up", "assets/warning.png")
}

var alertCmd = &cobra.Command{
	Use:     "alert duration",
	Aliases: []string{"a"},
	Short:   "Shows an alert after the chosen amount of time passes",
	Long: `Shows an alert after the chosen amount of time passes.

Valid time units:
  s - seconds
  m - minutes
  h - hours
  d - days`,
	Example: `alert 2m 30s
alert 1d
alert 3h 15m 20s`,
	Args: timeArgsFunc,
	RunE: alertRunFunc,
}

func init() {
	rootCmd.AddCommand(alertCmd)
}
