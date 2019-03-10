package logger

const (
	filePrinterName    = "file"
	consolePrinterName = "console"
	fullPrinterName    = "full"
)

var (
	typesMap = map[string]func() printer{
		filePrinterName:    newFilePrinter,
		consolePrinterName: newConsolePrinter,
		fullPrinterName:    newFullPrinter,
	}
)

func getLoggerInitializer(key string) func() printer {
	if initializer, ok := typesMap[key]; ok {
		return initializer
	}
	return newConsolePrinter
}
