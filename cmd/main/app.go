package main

import (
	"fmt"
	"homevideo_for_xlsx_parser/pkg/logging"
)

func main() {
	fmt.Println("Hello, World!")

	var logger = logging.GetLogger()
	logger.Info("Hello, World!")
}
