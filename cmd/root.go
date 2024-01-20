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

var validFormats = map[string]string{
	"layout":      time.Layout,
	"ansic":       time.ANSIC,
	"unixdate":    time.UnixDate,
	"rubydate":    time.RubyDate,
	"rfc822":      time.RFC822,
	"rfc822z":     time.RFC822Z,
	"rfc850":      time.RFC850,
	"rfc1123":     time.RFC1123,
	"rfc1123z":    time.RFC1123Z,
	"rfc3339":     time.RFC3339,
	"rfc3339nano": time.RFC3339Nano,
	"kitchen":     time.Kitchen,
	"stamp":       time.Stamp,
	"stampmilli":  time.StampMilli,
	"stampmicro":  time.StampMicro,
	"stampnano":   time.StampNano,
	"datetime":    time.DateTime,
	"dateonly":    time.DateOnly,
	"timeonly":    time.TimeOnly,
}

func argsFunc(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return nil
	}
	if len(args) > 1 {
		return fmt.Errorf("too many arguments")
	}
	format := args[0]

	if _, ok := validFormats[format]; !ok {
		return fmt.Errorf("invalid format")
	}

	return nil
}

func runFunc(cmd *cobra.Command, args []string) {
	format := "rfc3339"
	if len(args) > 0 {
		format = args[0]
	}
	fmt.Println(time.Now().Format(validFormats[format]))
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "timef [format]",
	Version: "v0.0.1",
	Run:    runFunc,
	Args:    argsFunc,
	Short:   "Displays the current time",
	Long: `Displays the current time. The default format is rfc3339.

Format options: layout, ansic, unixdate, rubydate, rfc822, rfc822z, rfc850,
rfc1123, rfc1123z, rfc3339, rfc3339nano, kitchen, stamp, stampmilli, stampmicro,
stampnano, datetime, dateonly, timeonly.

Examples and details: https://pkg.go.dev/time#pkg-constants
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Define any flags needed here.
}
