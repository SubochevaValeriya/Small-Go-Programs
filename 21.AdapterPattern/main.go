package main

import (
	"fmt"
)

//21.	Реализовать паттерн «адаптер» на любом примере

//Adapter pattern works as a bridge between two incompatible interfaces, like for example if we want that our old Printer satisfied new interface for modern printer.

type LegacyPrinter interface {
	Print(s string) string
}

type MyLegacyPrinter struct {
}

func (l *MyLegacyPrinter) Print(s string) string {
	newMsg := fmt.Sprintf("Legacy Printer: %s", s)
	println(newMsg)
	return newMsg
}

type ModernPrinter interface {
	PrintStored() string
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (p *PrinterAdapter) PrintStored() string {
	var newMsg string
	if p.OldPrinter != nil {
		newMsg = fmt.Sprintf("Legacy Printer: %s", p.Msg)
		println(newMsg)
	} else {
		newMsg = p.Msg
	}

	return newMsg
}

func main() {
	oldPrinter := MyLegacyPrinter{}

	adapter := PrinterAdapter{
		OldPrinter: &oldPrinter,
		Msg:        "Test",
	}

	adapter.PrintStored()
}
