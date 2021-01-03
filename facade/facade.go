package facade

import (
	"bufio"
	"context"
	"io"
	"os"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/gocli/exitcode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
	"github.com/spiegel-im-spiegel/gocli/signal"
	"github.com/spiegel-im-spiegel/ml/facade/interactive"
	"github.com/spiegel-im-spiegel/ml/facade/options"
	"github.com/spiegel-im-spiegel/ml/makelink"
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

			logfile, err := cmd.Flags().GetString("log")
			if err != nil {
				return debugPrint(ui, err)
			}
			var log io.Writer
			if len(logfile) > 0 {
				file, err := os.Create(logfile)
				if err != nil {
					return debugPrint(ui, err)
				}
				defer file.Close()
				log = file
			}
			opts := options.New(signal.Context(context.Background(), os.Interrupt), style, log)

			//interactive mode
			if interactiveFlag {
				return interactive.Do(opts)
			}

			//command line
			if len(args) > 0 {
				for _, arg := range args {
					r, err := opts.MakeLink(arg)
					if err != nil {
						return debugPrint(ui, err)
					}
					if err := ui.WriteFrom(r); err != nil {
						return debugPrint(ui, err)
					}
					_ = ui.Outputln()
				}
			} else {
				scanner := bufio.NewScanner(ui.Reader())
				for scanner.Scan() {
					r, err := opts.MakeLink(scanner.Text())
					if err != nil {
						return debugPrint(ui, err)
					}
					if err := ui.WriteFrom(r); err != nil {
						return debugPrint(ui, err)
					}
					_ = ui.Outputln()
				}
				return scanner.Err()
			}
			return nil
		},
	}
	rootCmd.SetArgs(args)
	rootCmd.SetOutput(ui.ErrorWriter())

	rootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "output version of "+Name)
	rootCmd.Flags().BoolVarP(&interactiveFlag, "interactive", "i", false, "interactive mode")
	rootCmd.Flags().BoolVarP(&debugFlag, "debug", "", false, "for debug")
	rootCmd.Flags().StringP("style", "s", makelink.StyleMarkdown.String(), "link style ["+makelink.StyleList()+"]")
	rootCmd.Flags().StringP("log", "", "", "output log")

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
