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
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

var validFormats = map[string]string{
	"layout":      time.Layout,
	"l":           time.Layout,
	"ansic":       time.ANSIC,
	"a":           time.ANSIC,
	"unixdate":    time.UnixDate,
	"u":           time.UnixDate,
	"rubydate":    time.RubyDate,
	"r":           time.RubyDate,
	"rfc822":      time.RFC822,
	"rfc822z":     time.RFC822Z,
	"rfc850":      time.RFC850,
	"rfc1123":     time.RFC1123,
	"rfc1123z":    time.RFC1123Z,
	"rfc3339":     time.RFC3339,
	"rfc3339nano": time.RFC3339Nano,
	"kitchen":     time.Kitchen,
	"k":           time.Kitchen,
	"stamp":       time.Stamp,
	"stampmilli":  time.StampMilli,
	"stampmicro":  time.StampMicro,
	"stampnano":   time.StampNano,
	"datetime":    time.DateTime,
	"dt":          time.DateTime,
	"dateonly":    time.DateOnly,
	"d":           time.DateOnly,
	"timeonly":    time.TimeOnly,
	"t":           time.TimeOnly,
}

func formatArgsFunc(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected 1 argument")
	}
	format := args[0]

	if _, ok := validFormats[format]; !ok {
		return fmt.Errorf("invalid format")
	}

	return nil
}

func formatRunFunc(cmd *cobra.Command, args []string) {
	if err := clipboard.Init(); err != nil {
		NoCopy = true
	}

	timeStr := time.Now().Format(validFormats[args[0]])
	if !NoCopy {
		clipboard.Write(clipboard.FmtText, []byte(timeStr))
	}
	fmt.Println(timeStr)
}

var formatCmd = &cobra.Command{
	Use:     "format",
	Aliases: []string{"f"},
	Short:   "Display the current time in a different format",
	Long: `Display the current time in a different format. The default format is rfc3339.

	Format options: layout, ansic, unixdate, rubydate, rfc822, rfc822z, rfc850,
	rfc1123, rfc1123z, rfc3339, rfc3339nano, kitchen, stamp, stampmilli, stampmicro,
	stampnano, datetime, dateonly, timeonly.
	
	Examples and details: https://pkg.go.dev/time#pkg-constants`,
	Run:  formatRunFunc,
	Args: formatArgsFunc,
}

func init() {
	rootCmd.AddCommand(formatCmd)
}
