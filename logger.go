package chromedp

import logger "github.com/honeyscience/honey-logger-go"

var Logger *logger.Logger

func init() {
	var err error

	Logger, err = logger.New()
	if err != nil {
		panic(err)
	}
}
