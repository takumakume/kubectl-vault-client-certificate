package main

import (
	"fmt"
	"io"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
)

type printer struct {
	table  *tablewriter.Table
	stdout io.Writer
	stderr io.Writer
}

func newPrinter(stdout, stderr io.Writer) *printer {
	log.SetOutput(stderr)

	table := tablewriter.NewWriter(stdout)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)

	return &printer{
		table:  table,
		stdout: stdout,
		stderr: stderr,
	}
}

func (p *printer) print(header []string, data [][]string, errors []error) error {
	if err := p.validatePrintData(header, data); err != nil {
		return err
	}

	p.table.SetHeader(header)
	p.table.AppendBulk(data)
	p.table.Render()

	for _, e := range errors {
		log.Error(e)
	}

	return nil
}

func (p *printer) validatePrintData(header []string, data [][]string) error {
	headrLen := len(header)
	for i, d := range data {
		if len(d) != headrLen {
			return fmt.Errorf("Inconsistent number of header and data columns. index:%d data:%+v", i, d)
		}
	}
	return nil
}
