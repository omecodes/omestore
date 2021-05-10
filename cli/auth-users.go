package cli

import (
	"encoding/json"
	"fmt"
	"github.com/omecodes/libome/logs"
	pb "github.com/omecodes/store/gen/go/proto"
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/omecodes/libome/crypt"
)

func init() {
	flags := saveCredentials.PersistentFlags()
	flags.StringVar(&input, "in", "", "Input file containing sequence of JSON encoded access")
	if err := cobra.MarkFlagRequired(flags, "in"); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	credentialsCMD.AddCommand(saveCredentials)
}

var credentialsCMD = &cobra.Command{
	Use:   "users",
	Short: "Manage user credentials",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var saveCredentials = &cobra.Command{
	Use:   "create",
	Short: "Creates user credentials",
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		file, err := os.Open(input)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		defer func() {
			_ = file.Close()
		}()

		cl := newClient()
		decoder := json.NewDecoder(file)
		for {
			var userCredentials *pb.UserCredentials
			err = decoder.Decode(&userCredentials)
			if err == io.EOF {
				return
			}

			if userCredentials.Password == "" {
				userCredentials.Password, err = crypt.GenerateVerificationCode(16)
				if err != nil {
					fmt.Println(err)
					os.Exit(-1)
				}
				logs.Info("Generated password for user", logs.Details("user", userCredentials.Username), logs.Details("password", userCredentials.Password))
			}

			err = cl.CreateUserCredentials(userCredentials)
			if err != nil {
				fmt.Println(err)
			}
		}
	},
}
