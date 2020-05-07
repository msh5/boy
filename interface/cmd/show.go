/*
Copyright © 2019 Sho Minagawa <msh5.global@gmail.com>

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
	"log"

	"github.com/spf13/cobra"

	"github.com/msh5/boy/interface/dependency"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "TODO",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		config := loadCommandConfig()

		ref := args[0]
    params := config.toDIContainerBuildParams()
    params.Ref = ref

		dependencies := dependency.NewCLIDependencies(params)
		if err := dependencies.ShowController.Handle(ref); err != nil {
			log.Fatal(err)
		}

		dependencies.ShowView.Update()
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
