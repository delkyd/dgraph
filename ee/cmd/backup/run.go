// +build !oss

/*
 * Copyright 2018 Dgraph Labs, Inc. and Contributors
 *
 * Licensed under the Dgraph Community License (the "License"); you
 * may not use this file except in compliance with the License. You
 * may obtain a copy of the License at
 *
 *     https://github.com/dgraph-io/dgraph/blob/master/licenses/DCL.txt
 */

package backup

import (
	"fmt"
	"os"

	"github.com/dgraph-io/dgraph/x"
	"github.com/spf13/cobra"
)

var Backup x.SubCommand

var opt struct {
	version   bool
	loc, http string
}

func init() {
	Backup.Cmd = &cobra.Command{
		Use:   "backup",
		Short: "Dgraph Enterprise Backup",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			defer x.StartProfile(Backup.Conf).Stop()
			if err := run(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	flag := Backup.Cmd.Flags()
	flag.StringVarP(&opt.loc, "loc", "l", "", "sets the location URI to a source or target.")
	flag.BoolVar(&opt.version, "version", false, "prints the version of Dgraph Backup.")
	flag.StringVar(&opt.http, "http", "http://localhost:8080", "HTTP address to Dgraph alpha.")
	flag.Bool("debugmode", false, "Enable debug mode for more debug information.")
}

func run() error {
	x.PrintVersion()
	if opt.version {
		return nil
	}
	x.Config.DebugMode = Backup.Conf.GetBool("debugmode")

	return runBackup()
}