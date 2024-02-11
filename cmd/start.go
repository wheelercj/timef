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

	"github.com/eiannone/keyboard"
	"github.com/gosuri/uilive"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

func startRunFunc(cmd *cobra.Command, args []string) error {
	if err := keyboard.Open(); err != nil {
		return err
	}
	defer func() {
		_ = keyboard.Close()
	}()

	writer := uilive.New()
	writer.Start()
	defer writer.Stop()

	return runStopwatch(writer)
}

var startCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"s"},
	Short:   "Start a stopwatch",
	Long: `Start a stopwatch

While the stopwatch is running, press space to pause, enter to lap, or escape to quit.`,
	Args: cobra.NoArgs,
	RunE: startRunFunc,
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func runStopwatch(writer *uilive.Writer) error {
	if err := clipboard.Init(); err != nil {
		NoCopy = true
	}

	keyChan := NewKeyChan()
	sleepChan := make(chan bool)

	go func() {
		for {
			time.Sleep(time.Second)
			sleepChan <- true
		}
	}()

	i := 0
	var isPaused bool
	var timeStr string
	for {
		select {
		case <-sleepChan:
			if !isPaused {
				timeStr = toTimeStr(i)
				fmt.Fprintf(writer, "%s\n", timeStr)
				i++
			}
		case key := <-keyChan:
			if key.Err != nil {
				return key.Err
			}

			switch key.Key {
			case keyboard.KeyEsc:
				// exit
				if !NoCopy {
					clipboard.Write(clipboard.FmtText, []byte(timeStr))
				}
				return nil
			case keyboard.KeyEnter:
				// lap
				fmt.Println("")
				if !NoCopy {
					clipboard.Write(clipboard.FmtText, []byte(timeStr))
				}
			case keyboard.KeySpace:
				// pause
				isPaused = !isPaused
				if !NoCopy {
					clipboard.Write(clipboard.FmtText, []byte(timeStr))
				}
			}
		}
	}
}
