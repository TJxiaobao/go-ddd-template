package option

import (
	"flag"

	"github.com/TJxiaobao/go-ddd-template/pkg/version"
)

type Options struct {
	ShowVersion *bool
	ShowConfig  *bool
	ConfigFile  *string
}

func New() *Options {
	opt := &Options{}

	opt.ShowVersion = flag.Bool("V", false, "Print the version and exit.")
	opt.ShowConfig = flag.Bool("c", false, "Print config with yaml format")
	opt.ConfigFile = flag.String("C", "", "Load configuration from a file(yaml or json format)")

	flag.Parse()

	return opt
}

func (opt *Options) ShowMessage() string {
	if *opt.ShowVersion {
		return version.Info
	}
	return ""
}
