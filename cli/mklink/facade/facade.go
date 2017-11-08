package facade

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/gocli"
	"github.com/spiegel-im-spiegel/mklink"
	"github.com/spiegel-im-spiegel/mklink/cli/mklink/interactive"
	"github.com/spiegel-im-spiegel/mklink/cli/mklink/makelink"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	//Name is applicatin name
	Name = "mklink"
	//Version is version for applicatin
	Version string
)

var (
	defaultStyle    = mklink.StyleMarkdown.String() //default link style
	versionFlag     bool                            //version flag
	interactiveFlag bool                            //interactive mode flag
	cui             = gocli.NewUI()                 //CUI instance
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: Name + " [flags] [URL [URL]...]",
	RunE: func(cmd *cobra.Command, args []string) error {
		//parse options
		if versionFlag {
			cui.OutputErr(Name)
			if len(Version) > 0 {
				cui.OutputErr(fmt.Sprintf(" v%s", Version))
			}
			cui.OutputErrln()
			return nil
		}

		strStyle, err := cmd.Flags().GetString("style")
		if err != nil {
			return err
		}
		style, err := mklink.GetStyle(strStyle)
		if err != nil {
			return err
		}

		logfile, err := cmd.Flags().GetString("log")
		if err != nil {
			return err
		}
		var log io.Writer
		if len(logfile) > 0 {
			file, err := os.OpenFile(logfile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
			if err != nil {
				return err
			}
			defer file.Close()
			log = file
		}
		lnk := makelink.New(style, cui.Writer(), log)

		if len(args) > 0 {
			for _, arg := range args {
				err := lnk.MakeLink(arg)
				if err != nil {
					return err
				}
			}
		} else if terminal.IsTerminal(int(syscall.Stdin)) {
			if interactiveFlag {
				interactive.New(style, log).Run()
			}
			return nil
		} else {
			scanner := bufio.NewScanner(cui.Reader())
			for scanner.Scan() {
				err := lnk.MakeLink(scanner.Text())
				if err != nil {
					return err
				}
			}
			if err := scanner.Err(); err != nil {
				return err
			}
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(ui *gocli.UI, args []string) (exit ExitCode) {
	defer func() {
		//panic hundling
		if r := recover(); r != nil {
			cui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, _, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				cui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ": line", line)
			}
			exit = ExitAbnormal
		}
	}()

	//execution
	cui = ui
	rootCmd.SetArgs(args)
	rootCmd.SetOutput(ui.ErrorWriter())
	exit = ExitNormal
	if err := rootCmd.Execute(); err != nil {
		exit = ExitAbnormal
	}
	return
}

func init() {
	rootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "output version of "+Name)
	rootCmd.Flags().BoolVarP(&interactiveFlag, "interactive", "i", false, "interactive mode")
	rootCmd.Flags().StringP("style", "s", defaultStyle, "link style")
	rootCmd.Flags().StringP("log", "", "", "output log")
}
