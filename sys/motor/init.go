/*
  Robots
  (c) Copyright David Thorpe 2018
  All Rights Reserved

  For Licensing and Usage information, please see LICENSE.md
  Documentation https://github.com/djthorpe/robots/

*/

package motor

import (
	// Frameworks
	"github.com/djthorpe/gopi"
)

////////////////////////////////////////////////////////////////////////////////
// INIT

func init() {
	gopi.RegisterModule(gopi.Module{
		Name:     "robots/motor",
		Type:     gopi.MODULE_TYPE_OTHER,
		Requires: []string{"pwm"},
		Config: func(config *gopi.AppConfig) {
			// TODO
		},
		New: func(app *gopi.AppInstance) (gopi.Driver, error) {
			// TODO
			return gopi.Open(Motor{
				PWM:      app.ModuleInstance("pwm").(gopi.PWM),
				Forward:  forward_pin,
				Backward: backward_pin,
			}, app.Logger)
		},
	})
}
