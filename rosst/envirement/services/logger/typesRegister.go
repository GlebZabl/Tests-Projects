package logger

const (
	FilePrinterName    = "file"
	ConsolePrinterName = "console"
	FullPrinterName    = "full"
)

var (
	typesMap = map[string]func() printer{
		FilePrinterName:    newFilePrinter,
		ConsolePrinterName: newConsolePrinter,
		FullPrinterName:    newFullPrinter,
	}
)

func getLoggerInitializer(key string) func() printer {
	if initializer, ok := typesMap[key]; ok {
		return initializer
	}
	return newConsolePrinter
}
