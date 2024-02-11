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

func notifyRunFunc(cmd *cobra.Command, args []string) error {
	seconds := toSeconds(args)
	sleepProgress(seconds)
	return beeep.Notify("Timer", "Time's up", "assets/information.png")
}

var notifyCmd = &cobra.Command{
	Use:     "notify duration",
	Aliases: []string{"n"},
	Short:   "Show a notification after the chosen amount of time passes",
	Long: `Show a notification after the chosen amount of time passes.

Valid time units:
  s - seconds
  m - minutes
  h - hours
  d - days`,
	Example: `notify 2m 30s
notify 1d
notify 3h 15m 20s`,
	Args: timeArgsFunc,
	RunE: notifyRunFunc,
}

func init() {
	rootCmd.AddCommand(notifyCmd)
}
