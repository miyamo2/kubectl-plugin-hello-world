/*
Copyright (c) 2024 miyamo2

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/klog"
)

// helloWorldCmd represents the base command when called without any subcommands
var helloWorldCmd = &cobra.Command{
	Use:   "hello_world",
	Short: "Prints 'Hello, World!'",
	Long:  `Prints 'Hello, World!'`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) > 0 {
			return fmt.Errorf("unknown argments: %+v", args)
		}
		cmd.Print("Hello, World!\n")
		return
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the helloWorldCmd.
func Execute() {
	if err := helloWorldCmd.Execute(); err != nil {
		klog.Exit(err.Error())
	}
}
