package runner

import (
	"fmt"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/nuclei/v2/pkg/catalog/config"
)

var banner = fmt.Sprintf(`
  __   ___
 |  |   |
 |  |   |
 |__|   | /   %s
`, config.Version)

// showBanner is used to show the banner to the user
func showBanner() {
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("\t\tQualitestProduct\n\n"

}
