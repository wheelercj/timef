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
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/gosuri/uilive"
	"github.com/spf13/cobra"
)

var validTimeUnits = []string{"s", "m", "h", "d"}

const (
	minuteSeconds = 60
	hourSeconds   = 60 * 60
	daySeconds    = 60 * 60 * 24
	weekSeconds   = 60 * 60 * 24 * 7
	monthSeconds  = 60 * 60 * 24 * 30
	yearSeconds   = 60 * 60 * 24 * 365
)

func timeArgsFunc(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("expected at least 1 argument")
	}

	for _, arg := range args {
		lastLetter := string(arg[len(arg)-1])
		if !slices.Contains(validTimeUnits, lastLetter) {
			return fmt.Errorf(
				"time duration `%s` does not end with a valid time unit",
				arg,
			)
		}

		numStr := arg[:len(arg)-1]
		if _, err := strconv.ParseInt(numStr, 10, 64); err != nil {
			return fmt.Errorf(
				"time duration `%s` must be an integer followed by a unit",
				arg,
			)
		}
	}

	return nil
}

// toSeconds converts strings of time durations like ["2m", "30s"] into an integer of
// seconds. The strings are expected to be valid.
func toSeconds(tokens []string) int {
	var seconds int
	for _, token := range tokens {
		lastLetter := string(token[len(token)-1])
		num := MustParseInt(token[:len(token)-1])

		switch lastLetter {
		case "s":
			seconds += num
		case "m":
			seconds += num * minuteSeconds
		case "h":
			seconds += num * hourSeconds
		case "d":
			seconds += num * daySeconds
		}
	}

	return seconds
}

// toTimeStr converts an integer of seconds to a time string like "2m 30s".
func toTimeStr(seconds int) string {
	var tokens []string

	if seconds > daySeconds {
		tokens = append(tokens, fmt.Sprintf("%dd", seconds/daySeconds))
		seconds %= daySeconds
	}
	if seconds > hourSeconds {
		tokens = append(tokens, fmt.Sprintf("%dh", seconds/hourSeconds))
		seconds %= hourSeconds
	}
	if seconds > minuteSeconds {
		tokens = append(tokens, fmt.Sprintf("%dm", seconds/minuteSeconds))
		seconds %= minuteSeconds
	}
	tokens = append(tokens, fmt.Sprintf("%ds", seconds))

	return strings.Join(tokens, " ")
}

// sleepProgress sleeps for the chosen time while showing a countdown.
func sleepProgress(seconds int) {
	writer := uilive.New()
	writer.Start()

	for i := seconds; i > 0; i-- {
		timeStr := toTimeStr(i)
		fmt.Fprintf(writer, "%s remaining...\n", timeStr)
		time.Sleep(time.Second)
	}

	fmt.Fprintln(writer, "Done")
	writer.Stop()
}

// MustParseInt converts a string to an integer, and panics if the conversion is not
// possible.
func MustParseInt(numStr string) int {
	if num64, err := strconv.ParseInt(numStr, 10, 64); err == nil {
		return int(num64)
	} else {
		panic(err)
	}
}
