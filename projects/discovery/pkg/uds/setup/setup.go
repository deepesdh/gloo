package setup

import (
	"context"
	"github.com/solo-io/go-utils/contextutils"

	"github.com/solo-io/gloo/pkg/utils/setuputils"
	"github.com/solo-io/gloo/pkg/version"
	"github.com/solo-io/gloo/projects/discovery/pkg/uds/syncer"
	"github.com/solo-io/gloo/projects/gloo/pkg/syncer/setup"
)

func Main(customCtx context.Context) error {
	contextutils.LoggerFrom(customCtx).Info("(2)-> called from projects/discovery/pkg/uds/setup/setup.go/Main")
	return setuputils.Main(setuputils.SetupOpts{
		LoggerName:  "uds",
		Version:     version.Version,
		SetupFunc:   setup.NewSetupFuncWithRun(syncer.RunUDS),
		ExitOnError: true,
		CustomCtx:   customCtx,
	})
}
