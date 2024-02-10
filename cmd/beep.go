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

func beepRunFunc(cmd *cobra.Command, args []string) error {
	seconds := toSeconds(args)
	sleepProgress(seconds)
	return beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
}

var beepCmd = &cobra.Command{
	Use:     "beep duration",
	Aliases: []string{"b"},
	Short:   "Makes a beep sound after the chosen amount of time passes",
	Long: `Makes a beep sound after the chosen amount of time passes.

Valid time units:
  s - seconds
  m - minutes
  h - hours
  d - days`,
	Example: `beep 2m 30s
beep 1d
beep 3h 15m 20s`,
	Args: timeArgsFunc,
	RunE: beepRunFunc,
}

func init() {
	rootCmd.AddCommand(beepCmd)
}
