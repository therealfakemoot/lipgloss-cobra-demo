package demo

import (
	"bytes"
	"fmt"
	"io"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

func CharmUsage(c *cobra.Command) error {
	var b bytes.Buffer
	err := tmpl(&b, c.UsageTemplate(), c)
	if err != nil {
		c.PrintErrln(err)
		return err
	}
	pretty(c.ErrOrStderr(), b.String())
	return nil
}

func CharmHelp(c *cobra.Command, a []string) {
	var b bytes.Buffer
	// The help should be sent to stdout
	// See https://github.com/spf13/cobra/issues/1002
	err := tmpl(&b, c.HelpTemplate(), c)
	if err != nil {
		c.PrintErrln(err)
	}
	pretty(c.ErrOrStderr(), b.String())
}

var BaseStyle = lipgloss.NewStyle().Bold(true).BorderStyle(lipgloss.RoundedBorder()).Width(60).PaddingLeft(1).PaddingRight(1).PaddingBottom(2)

func pretty(w io.Writer, s string) {
	fmt.Fprintf(w, "%s\n", BaseStyle.Render(s))
}
