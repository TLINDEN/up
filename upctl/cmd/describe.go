/*
Copyright © 2023 Thomas von Dein

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/tlinden/up/upctl/cfg"
	"github.com/tlinden/up/upctl/lib"
)

func DescribeCommand(conf *cfg.Config) *cobra.Command {
	var listCmd = &cobra.Command{
		Use:   "describe [options] upload-id",
		Long:  "Show detailed informations about an upload object.",
		Short: `Describe an upload.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("No id specified to delete!")
			}

			// errors at this stage do not cause the usage to be shown
			cmd.SilenceUsage = true

			return lib.Describe(conf, args)
		},
	}

	listCmd.Aliases = append(listCmd.Aliases, "des")
	listCmd.Aliases = append(listCmd.Aliases, "info")
	listCmd.Aliases = append(listCmd.Aliases, "i")

	return listCmd
}