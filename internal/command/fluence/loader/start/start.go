package start

import (
	"github.com/urfave/cli/v2"

	"iusworks.com/p/polar/internal/web"
)

func Loader(context *cli.Context) error {
	web.Loader(context)
	return nil
}
