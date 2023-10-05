package printer

import (
	"os"

	"github.com/asdsec/at42/internal/env"
	"github.com/fatih/color"
)

var (
	ActiveItemColor = color.New(color.FgGreen, color.Bold)
)

func init() {
	EnableOrDisableColor(ActiveItemColor)
}

func useColors() *bool {
	tr, fa := true, false
	if os.Getenv(env.EnvForceColor) != "" {
		return &tr
	} else if os.Getenv(env.EnvNoColor) != "" {
		return &fa
	}
	return nil
}

func EnableOrDisableColor(c *color.Color) {
	if v := useColors(); v != nil && *v {
		c.EnableColor()
	} else if v != nil && !*v {
		c.DisableColor()
	}
}
