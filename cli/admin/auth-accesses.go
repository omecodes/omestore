package admin

import (
	"encoding/json"
	"fmt"
	"github.com/omecodes/libome/crypt"
	"github.com/omecodes/store/auth"
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/omecodes/common/utils/prompt"
)

func init() {
	flags := saveAccessCMD.PersistentFlags()
	flags.StringVar(&input, "in", "", "Input file containing sequence of JSON encoded access")
	flags.StringVar(&server, "server", "", "Server API location")
	flags.StringVar(&password, "password", "", "admin password")
	if err := cobra.MarkFlagRequired(flags, "server"); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	if err := cobra.MarkFlagRequired(flags, "in"); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	flags = getAccessesCMD.PersistentFlags()
	flags.StringVar(&server, "server", "", "Server API location")
	flags.StringVar(&password, "password", "", "admin password")
	flags.StringVar(&output, "out", "", "Output file")
	if err := cobra.MarkFlagRequired(flags, "server"); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	flags = deleteAccessesCMD.PersistentFlags()
	flags.StringVar(&server, "server", "", "Server API location")
	flags.StringVar(&password, "password", "", "admin password")
	flags.StringArrayVar(&ids, "id", nil, "Access ID")
	if err := cobra.MarkFlagRequired(flags, "server"); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	if err := cobra.MarkFlagRequired(flags, "id"); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	accessCMD.AddCommand(saveAccessCMD)
	accessCMD.AddCommand(getAccessesCMD)
	accessCMD.AddCommand(deleteAccessesCMD)
}

var accessCMD = &cobra.Command{
	Use:   "access",
	Short: "Manage store accesses",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var saveAccessCMD = &cobra.Command{
	Use:   "set",
	Short: "save accesses",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if password == "" {
			password, err = prompt.Password("password")
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
		}

		file, err := os.Open(input)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		defer func() {
			_ = file.Close()
		}()

		decoder := json.NewDecoder(file)
		for {
			var access *auth.ClientApp
			err = decoder.Decode(&access)
			if err == io.EOF {
				return
			}

			if access.Secret == "" {
				access.Secret, err = crypt.GenerateVerificationCode(16)
				if err != nil {
					fmt.Println(err)
					os.Exit(-1)
				}
			}

			err = putAccess(password, access)
			if err != nil {
				fmt.Println(err)
			}
		}
	},
}

var getAccessesCMD = &cobra.Command{
	Use:   "get",
	Short: "Get all accesses",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if password == "" {
			password, err = prompt.Password("password")
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
		}

		err = getAccesses(password, output)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	},
}

var deleteAccessesCMD = &cobra.Command{
	Use:   "del",
	Short: "Delete accesses",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if password == "" {
			password, err = prompt.Password("password")
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
		}

		for _, id := range ids {
			err = deleteAccess(password, id)
			if err != nil {
				fmt.Println(err)
			}
		}
	},
}
