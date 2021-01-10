package main

import (
	"fmt"
	"github.com/omecodes/common/utils/log"
	"github.com/omecodes/common/utils/prompt"
	oms "github.com/omecodes/store"
	"github.com/omecodes/store/cmd/admin"
	"github.com/spf13/cobra"
	"os"
	"path"
	"path/filepath"
)

var (
	dev        bool
	adminInfo  string
	workingDir string
	dsn        string
	domains    []string
	command    *cobra.Command
)

func init() {
	command = &cobra.Command{
		Use:   path.Base(os.Args[0]),
		Short: "Monolithic object backend",
		Run: func(cmd *cobra.Command, args []string) {
			_ = command.Help()
		},
	}

	runCMD := &cobra.Command{
		Use:   "run",
		Short: "Runs a objects backend application",
		Run: func(cmd *cobra.Command, args []string) {
			if !dev && len(domains) == 0 {
				fmt.Println("Flag --domains is required when --auto-cert is set")
				os.Exit(-1)
			}

			var err error
			if workingDir == "" {
				workingDir, err = filepath.Abs("./")
				if err != nil {
					fmt.Println(err)
					os.Exit(-1)
				}
			}

			s := oms.NewMNServer(oms.MNConfig{
				WorkingDir: workingDir,
				Domains:    domains,
				Dev:        dev,
				AdminInfo:  adminInfo,
				DSN:        dsn,
			})

			err = s.Start()
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
			defer s.Stop()

			select {
			case <-prompt.QuitSignal():
			case err = <-s.Errors:
				log.Error("server error", log.Err(err))
			}
		},
	}

	flags := runCMD.PersistentFlags()
	flags.BoolVar(&dev, "dev", false, "Enable development mode")
	flags.StringArrayVar(&domains, "domains", nil, "Domains name for auto cert")
	flags.StringVar(&workingDir, "dir", "", "Data directory")
	flags.StringVar(&adminInfo, "admin", "", "Admin password info")
	flags.StringVar(&dsn, "dsn", "store:store@(127.0.0.1:3306)/store?charset=utf8", "MySQL database uri")
	if err := cobra.MarkFlagRequired(flags, "admin"); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	command.AddCommand(runCMD)

	versionCMD := &cobra.Command{
		Use:   "version",
		Short: "Version info",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println()
			fmt.Println("   Version: ", oms.Version)
			fmt.Println("  Revision: ", oms.Revision)
			fmt.Println("Build date: ", oms.BuildDate)
			fmt.Println("   License: ", oms.License)
			fmt.Println()
		},
	}
	command.AddCommand(versionCMD)
	command.AddCommand(admin.Cmd)
}

func main() {
	err := command.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
