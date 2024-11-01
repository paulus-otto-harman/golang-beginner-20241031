package util

import (
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
)

func renderTimeout() string {
	if SessionTimeout != 0 {
		return fmt.Sprintf("\n%s\n", gola.Tf(gola.Bold, fmt.Sprintf("Timeout %d seconds", SessionTimeout), gola.LightYellow))
	}
	return fmt.Sprintf("\n%s\n", gola.Tf(gola.Bold, "Timeout Disabled", gola.LightGreen))
}

func H1(title string) {
	gola.ClearScreen()

	fmt.Println(renderTimeout())

	fmt.Printf("%s\n\n", gola.Tf(gola.Bold, fmt.Sprintf("%s %s - %s %s", "===", AppName, title, "==="), gola.LightBlue))
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

func GoBackOrNot() int {
	if cont, _ := gola.ToString(gola.Input(gola.Args(gola.P("label", "Press Enter to continue or [0] to return to previous menu")))); cont == "0" {
		return 0
	}
	return -1
}

func WaitToGoBack() int {
	gola.Wait("Press Enter to return to previous menu")
	return 0
}
