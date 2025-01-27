package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	if err := Cmd().Execute(); err != nil {
		os.Exit(1)
	}
}

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gg [org/repo]",
		Short: "Clone a GitHub repo into an org-specific directory on your desktop",
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			if len(args) != 1 {
				return errors.New("exactly one argument is required")
			}
			parts := strings.Split(args[0], "/")
			if len(parts) != 2 {
				return errors.New("argument must be of the form ORG/REPO")
			}
			dirname, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			org := parts[0]
			orgdir := filepath.Join(dirname, "Desktop", org)
			err = os.MkdirAll(orgdir, 0755)
			if err != nil {
				return err
			}
			c := exec.Command("git", "clone", "git@github.com:"+args[0])
			c.Dir = orgdir
			out, err := c.CombinedOutput()
			fmt.Printf("%s\n", string(out))
			return err
		},
	}
	return cmd
}
