package main

import (
	"testing"

    "github.com/Sneaksolid/cli/cli"
    "github.com/Sneaksolid/cli/command"
)

/*********************************************************************************
*     File Name           :     test/main_test.go
*     Created By          :     jonesax
*     Creation Date       :     [2017-06-26 18:34]
*     Last Modified       :     [2017-06-26 18:34]
*     Description         :
**********************************************************************************/
func TestAddCommand(t *testing.T) {

    cli := cli.NewCli(">>>  ", "exit")

	cli.AddCommand(command.Command{})

	if len(cli.Commands) != 1 {
		t.Error("Incorrect arg count")
	}
}
