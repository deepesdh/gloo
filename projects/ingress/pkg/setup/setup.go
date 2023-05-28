package setup

import (
	"context"
	"github.com/solo-io/go-utils/contextutils"

	"github.com/solo-io/gloo/pkg/version"

	"github.com/solo-io/gloo/pkg/utils/setuputils"
)

func Main(customCtx context.Context) error {
	contextutils.LoggerFrom(customCtx).Info("(2)-> called from projects/ingress/pkg/setup/setup.go/Main")
	return setuputils.Main(setuputils.SetupOpts{
		LoggerName:  "ingress",
		Version:     version.Version,
		SetupFunc:   Setup,
		ExitOnError: true,
		CustomCtx:   customCtx,
	})
}
