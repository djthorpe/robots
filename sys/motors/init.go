/*
  Robots
  (c) Copyright David Thorpe 2018
  All Rights Reserved

  For Licensing and Usage information, please see LICENSE.md
  Documentation https://github.com/djthorpe/robots/

*/

package motors

import (
	// Frameworks
	"github.com/djthorpe/gopi"
	"github.com/djthorpe/gopi-hw/sys/pwm"
)

////////////////////////////////////////////////////////////////////////////////
// INIT

func init() {
	gopi.RegisterModule(gopi.Module{
		Name:     "robots/motors",
		Type:     gopi.MODULE_TYPE_OTHER,
		Requires: []string{"pwm"},
		Config: func(config *gopi.AppConfig) {
			config.AppFlags.FlagString("motors.frequency", "", "Default frequency for driving motor")
		},
		New: func(app *gopi.AppInstance) (gopi.Driver, error) {
			config := Motors{
				PWM: app.ModuleInstance("pwm").(gopi.PWM),
			}
			if frequency, _ := app.AppFlags.GetString("motors.frequency"); frequency == "" {
				// Do nothing
			} else if period, err := pwm.ParseFrequency(frequency); err != nil {
				// Return error
				return nil, err
			} else {
				config.Period = period
			}
			return gopi.Open(config, app.Logger)
		},
	})
}
