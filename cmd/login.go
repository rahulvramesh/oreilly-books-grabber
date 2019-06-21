/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login to safari books",
	Long:  `login to safari books safari books using username and password`,
	Run: func(cmd *cobra.Command, args []string) {

		email, _ := cmd.Flags().GetString("email")
		password, _ := cmd.Flags().GetString("password")

		if email == "" {
			fmt.Println("email can't be empty")
			os.Exit(1)
		}

		if password == "" {
			fmt.Println("password can't be empty")
			os.Exit(1)
		}

		//save the login details to the config file

		fmt.Println("Successfully Saved!")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("email", "e", "", "Enter Your Email")
	loginCmd.Flags().StringP("password", "p", "", "Enter Your Password")

	viper.Set("Verbose", true)
}
