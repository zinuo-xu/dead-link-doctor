package reporter

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/zinuo-xu/dead-link-doctor/internal/checker"
)

var (
	green = color.New(color.FgGreen).SprintFunc()
	red   = color.New(color.FgRed).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
)

func Terminal(results []checker.Result) {
	ok, broken := 0, 0
	for _, r := range results {
		if r.OK { ok++ } else { broken++ }
	}
	fmt.Printf("\nResults: %s %s\n", green(fmt.Sprintf("%d OK", ok)), red(fmt.Sprintf("%d Broken", broken)))
	for _, r := range results {
		if !r.OK {
			fmt.Printf("  %s %s -> %s\n", red("✗"), r.URL, r.Error)
		}
	}
}
