package util

import (
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
)

func H1(title string) {
	gola.ClearScreen()
	if SessionTimeout == 0 {
		fmt.Printf("\n%s\n", gola.Tf(gola.Bold, "Timeout Disabled", gola.LightGreen))
	}

	fmt.Printf("\n%s\n\n", gola.Tf(gola.Bold, fmt.Sprintf("%s %s - %s %s", "===", AppName, title, "==="), gola.LightBlue))
}

func Select(options map[string]interface{}, label string) string {
	for optionValue, option := range options {
		fmt.Printf("[%s] %v", optionValue, option)
	}
	selected, err := gola.ToString(gola.Input(gola.Args(gola.P("label", label))))
	if err != nil {

	}
	return selected
}
