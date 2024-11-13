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
	"k8s.io/cli-runtime/pkg/genericiooptions"
)

// helloWorldExample is an example of how to use the helloWorld command.
var helloWorldExample = `
# Prints 'Hello, World!'
kubectl hello_world`

// HelloWorldOptions provides information required to print 'Hello, World!'.
type HelloWorldOptions struct {
	args []string
	genericiooptions.IOStreams
}

// NewHelloWorldOptions provides an instance of HelloWorldOptions with default values.
func NewHelloWorldOptions(streams genericiooptions.IOStreams) *HelloWorldOptions {
	return &HelloWorldOptions{
		IOStreams: streams,
	}
}

// NewCmdHelloWorld provides a cobra command wrapping HelloWorldOptions.
func NewCmdHelloWorld(streams genericiooptions.IOStreams) *cobra.Command {
	o := NewHelloWorldOptions(streams)
	cmd := &cobra.Command{
		Use:          "hello_world [flags]",
		Short:        "Prints 'Hello, World!'",
		Example:      helloWorldExample,
		SilenceUsage: true,
		Annotations: map[string]string{
			cobra.CommandDisplayNameAnnotation: "kubectl hello_world",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := o.Complete(args); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			return o.Run(cmd)
		},
	}
	cmd.SetIn(streams.In)
	cmd.SetOut(streams.Out)
	cmd.SetErr(streams.ErrOut)

	return cmd
}

// Complete completes all the required options
func (o *HelloWorldOptions) Complete(args []string) error {
	o.args = args
	return nil
}

// Validate validates the provided options
func (o *HelloWorldOptions) Validate() error {
	if len(o.args) > 0 {
		return fmt.Errorf("unknown argments: %+v", o.args)
	}
	return nil
}

// Run prints 'Hello, World!'
func (o *HelloWorldOptions) Run(cmd *cobra.Command) error {
	cmd.Printf("Hello, World!\n")
	return nil
}
