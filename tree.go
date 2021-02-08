package main

import (
	"github.com/thatisuday/commando"
)

/*
 *	GUIDE LINK: https://medium.com/sysf/building-simple-command-line-cli-applications-in-go-using-commando-8a8e03dbd48a
 */

func main() {
	// configure commando
	commando.
		SetExecutableName("tree").
		SetVersion("1.0.0").
		SetDescription("This tool lists the contents oc a directory in a tree-like format.\nIt can also display information about files and folders like size, permission and ownership.")

	// Configure the root command
	commando.
		Register(nil).
		AddArgument("dir", "local directory path", "./").                 // default `./`
		AddFlag("level, l", "level of depth to travel", commando.Int, 1). // default `1`
		AddFlag("size", "display size of each file", commando.Bool, nil). // default `false`
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			// fmt.Printf("Printing options for the `root` command...\n\n")

			// // print arguments
			// for k, v := range args {
			// 	fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
			// }

			// // print flags
			// for k, v := range flags {
			// 	fmt.Printf("flag -> %v: %v(%T)\n", k, v.Value, v.Value)
			// }

			dir := args["dir"].Value
			// level, _ := flags["level"].GetInt()
			// size, _ := flags["size"].GetBool()

			// list(isInfo bool, dir string, level int, displaySize bool)
			// list(false, dir, level, size)
			list(false, dir, flags)
		})

	// configure info command
	commando.
		Register("info").
		SetShortDescription("displays detailed information about the directory").
		SetDescription("This command displays more information about the contents of the directory like size, permission and ownership, etc").
		AddArgument("dir", "local directory path", "./").                  // default `./`
		AddFlag("level,l", "level of depth to travel", commando.Int, nil). //required
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			// fmt.Printf("Printing options of the `info` command...\n\n")

			// // print arguments
			// for k, v := range args {
			// 	fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
			// }

			// // print flags
			// for k, v := range flags {
			// 	fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
			// }

			dir := args["dir"].Value
			// level, _ := flags["level"].GetInt()

			// list(isInfo bool, dir string, level int, displaySize bool)
			// list(true, dir, level, false)
			list(true, dir, flags)
		})

	// parse command-line arguments
	commando.Parse(nil)
}
