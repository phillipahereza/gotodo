package gotodo

import (
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "gotodo",
	Doc:  "reports the locations of TODO comments",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		for _, com := range f.Comments {
			if strings.HasPrefix(strings.ToLower(com.Text()), "todo") {
				pass.Reportf(com.Pos(), "%s", strings.TrimSuffix(com.Text(), "\n"))
			}
		}
	}
	return nil, nil
}
