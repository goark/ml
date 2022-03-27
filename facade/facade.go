package facade

import (
	"bufio"
	"context"
	"os"
	"runtime"

	"github.com/goark/gocli/exitcode"
	"github.com/goark/gocli/rwi"
	"github.com/goark/gocli/signal"
	"github.com/goark/ml/facade/history"
	"github.com/goark/ml/facade/interactive"
	"github.com/goark/ml/facade/options"
	"github.com/goark/ml/makelink"
	"github.com/spf13/cobra"
)

var (
	//Name is applicatin name
	Name = "ml"
	//Version is version for applicatin
	Version = "dev-version"
)

var (
	versionFlag     bool //version flag
	interactiveFlag bool //interactive mode flag
	debugFlag       bool //debug flag
)

//newRootCmd returns cobra.Command instance for root command
func newRootCmd(ui *rwi.RWI, args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use: Name + " [flags] [URL [URL]...]",
		RunE: func(cmd *cobra.Command, args []string) error {
			//parse options
			if versionFlag {
				return ui.OutputErrln(Name, Version)
			}

			strStyle, err := cmd.Flags().GetString("style")
			if err != nil {
				return debugPrint(ui, err)
			}
			style, err := makelink.GetStyle(strStyle)
			if err != nil {
				return debugPrint(ui, err)
			}
			userAgent, err := cmd.Flags().GetString("user-agent")
			if err != nil {
				return debugPrint(ui, err)
			}

			//history log
			log, err := cmd.Flags().GetInt("log") //log size
			if err != nil {
				return debugPrint(ui, err)
			}
			hist := history.NewFile(log, historyPath())
			if log > 0 {
				if err := mkdirHistory(); err != nil {
					_ = ui.OutputErrln(err)
				} else if err := hist.Load(); err != nil {
					_ = ui.OutputErrln(err)
				}
			}
			defer func() {
				if err := hist.Save(); err != nil {
					_ = debugPrint(ui, err)
				}
			}()

			//set options
			opts := options.New(style, hist, userAgent)
			if interactiveFlag {
				//interactive mode
				if err := interactive.Do(opts); err != nil {
					return debugPrint(ui, err)
				}
			} else {
				//command line
				ctx := signal.Context(context.Background(), os.Interrupt)
				if len(args) > 0 {
					var lastErr error
					for _, arg := range args {
						if r, err := opts.MakeLink(ctx, arg); err != nil {
							_ = ui.OutputErrln(err)
							lastErr = err
						} else if err := ui.WriteFrom(r); err != nil {
							_ = ui.OutputErrln(err)
							lastErr = err
						} else {
							_ = ui.Outputln()
						}
					}
					if lastErr != nil {
						return debugPrint(ui, lastErr)
					}
				} else {
					var lastErr error
					scanner := bufio.NewScanner(ui.Reader())
					for scanner.Scan() {
						if r, err := opts.MakeLink(ctx, scanner.Text()); err != nil {
							_ = ui.OutputErrln(err)
							lastErr = err
						} else if err := ui.WriteFrom(r); err != nil {
							_ = ui.OutputErrln(err)
							lastErr = err
						} else {
							_ = ui.Outputln()
						}
					}
					if lastErr != nil {
						return debugPrint(ui, lastErr)
					}
					return debugPrint(ui, scanner.Err())
				}
			}
			return nil
		},
	}
	rootCmd.SilenceUsage = true
	rootCmd.SetArgs(args)
	rootCmd.SetIn(ui.Reader())       //Stdin
	rootCmd.SetOut(ui.ErrorWriter()) //Stdout -> Stderr
	rootCmd.SetErr(ui.ErrorWriter()) //Stderr

	rootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "output version of "+Name)
	rootCmd.Flags().BoolVarP(&interactiveFlag, "interactive", "i", false, "interactive mode")
	rootCmd.Flags().BoolVarP(&debugFlag, "debug", "", false, "for debug")
	rootCmd.Flags().StringP("style", "s", makelink.StyleMarkdown.String(), "link style ["+makelink.StyleList()+"]")
	rootCmd.Flags().StringP("user-agent", "a", "", "User-Agent string")
	rootCmd.Flags().IntP("log", "l", 0, "history log size")

	return rootCmd
}

//Execute is called from main function
func Execute(ui *rwi.RWI, args []string) (exit exitcode.ExitCode) {
	defer func() {
		//panic hundling
		if r := recover(); r != nil {
			_ = ui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, _, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				_ = ui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ": line", line)
			}
			exit = exitcode.Abnormal
		}
	}()

	//execution
	exit = exitcode.Normal
	if err := newRootCmd(ui, args).Execute(); err != nil {
		exit = exitcode.Abnormal
	}
	return
}

/* Copyright 2017-2021 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
