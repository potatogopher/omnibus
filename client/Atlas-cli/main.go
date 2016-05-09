package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"goa-atlas/client"
	"os"
	"time"
)

// PrettyPrint is true if the tool output should be formatted for human consumption.
var PrettyPrint bool

func main() {
	// Create command line parser
	app := &cobra.Command{
		Use:   "Atlas-cli",
		Short: `CLI client for the Atlas service`,
	}
	c := client.New(nil)
	c.UserAgent = "Atlas-cli/1.0"
	app.PersistentFlags().StringVarP(&c.Scheme, "scheme", "s", "", "Set the requests scheme")
	app.PersistentFlags().StringVarP(&c.Host, "host", "H", "localhost", "API hostname")
	app.PersistentFlags().DurationVarP(&c.Timeout, "timeout", "t", time.Duration(20)*time.Second, "Set the request timeout")
	app.PersistentFlags().BoolVar(&c.Dump, "dump", false, "Dump HTTP request and response.")
	app.PersistentFlags().BoolVar(&PrettyPrint, "pp", false, "Pretty print response body")
	RegisterCommands(app, c)
	if err := app.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "request failed: %s", err)
		os.Exit(-1)
	}
}

// RegisterCommands all the resource action subcommands to the application command line.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "create",
		Short: `Create new user`,
	}
	tmp1 := new(CreateUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/users"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "delete",
		Short: ``,
	}
	tmp2 := new(DeleteUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/users/:userID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "show",
		Short: `Retrieve user with given id. IDs 1 and 2 pre-exist in the system.`,
	}
	tmp3 := new(ShowUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/users/:userID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "update",
		Short: `Change user attributes`,
	}
	tmp4 := new(UpdateUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/users/:userID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	command.AddCommand(sub)
	app.AddCommand(command)

}
