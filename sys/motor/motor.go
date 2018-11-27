package motor

import (
	"fmt"

	"github.com/djthorpe/gopi"
)

type Motor struct {
	PWM      gopi.PWM
	Forward  gopi.GPIOPin
	Backward gopi.GPIOPin
}

type motor struct {
	forward  gopi.GPIOPin
	backward gopi.GPIOPin
	pwm      gopi.PWM
	log      gopi.Logger
}

////////////////////////////////////////////////////////////////////////////////
// OPEN AND CLOSE

// Create new Motor object
func (config Motor) Open(log gopi.Logger) (gopi.Driver, error) {
	log.Debug("<sys.robots.Motor>Open{ Forward=%v Backward=%v PWM=%v }", config.Forward, config.Backward, config.PWM)

	// Check incoming parameters
	if config.PWM == nil || config.Forward == gopi.GPIO_PIN_NONE || config.Backward == gopi.GPIO_PIN_NONE {
		return nil, gopi.ErrBadParameter
	}

	// create new PWM driver
	this := new(motor)

	// Set logging
	this.log = log
	this.pwm = config.PWM
	this.forward = config.Forward
	this.backward = config.Backward

	// Enable PWM on pins and set duty cycle to zero
	if err := this.pwm.SetDutyCycle(0, this.forward); err != nil {
		return nil, err
	} else if err := this.pwm.SetDutyCycle(0, this.backward); err != nil {
		return nil, err
	}

	// success
	return this, nil
}

// Close Motor object
func (this *motor) Close() error {
	this.log.Debug("<sys.robots.Motor>Close{}")

	// Set duty cycle to zero for motor
	if err := this.pwm.SetDutyCycle(0, this.forward); err != nil {
		return err
	} else if err := this.pwm.SetDutyCycle(0, this.backward); err != nil {
		return err
	}

	// Release resources
	this.pwm = nil

	return nil
}

////////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *motor) String() string {
	return fmt.Sprintf("<sys.robots.Motor>{ }")
}

/*
# Motor, via DRV8833PWP Dual H-Bridge
M1B = 19
M1F = 20
M2B = 21
M2F = 26

class Motor(object):
    type = 'Motor'

    def __init__(self, pin_fw, pin_bw):
        self._invert = False
        self.pin_fw = pin_fw
        self.pin_bw = pin_bw
        self._speed = 0
        self._gpio_is_setup = False

    def _setup_gpio(self):
        if self._gpio_is_setup:
            return

        self._gpio_is_setup = True
        setup_gpio(self.pin_fw, GPIO.OUT, initial=GPIO.LOW)
        setup_gpio(self.pin_bw, GPIO.OUT, initial=GPIO.LOW)

        self.pwm_fw = GPIO.PWM(self.pin_fw, 100)
        self.pwm_fw.start(0)

        self.pwm_bw = GPIO.PWM(self.pin_bw, 100)
        self.pwm_bw.start(0)

    def invert(self):
        self._invert = not self._invert
        self._speed = -self._speed
        self.speed(self._speed)
        return self._invert

    def forwards(self, speed=100):
        if speed > 100 or speed < 0:
            raise ValueError("Speed must be between 0 and 100")
        if self._invert:
            self.speed(-speed)
        else:
            self.speed(speed)

    def backwards(self, speed=100):
        if speed > 100 or speed < 0:
            raise ValueError("Speed must be between 0 and 100")
        if self._invert:
            self.speed(speed)
        else:
            self.speed(-speed)

    def speed(self, speed=100):
        self._setup_gpio()

        if speed > 100 or speed < -100:
            raise ValueError("Speed must be between -100 and 100")

        self._speed = speed
        if speed > 0:
            self.pwm_bw.ChangeDutyCycle(0)
            self.pwm_fw.ChangeDutyCycle(speed)
        if speed < 0:
            self.pwm_fw.ChangeDutyCycle(0)
            self.pwm_bw.ChangeDutyCycle(abs(speed))
        if speed == 0:
            self.pwm_fw.ChangeDutyCycle(0)
            self.pwm_bw.ChangeDutyCycle(0)

        return speed

    def stop(self):
        self.speed(0)

    forward = forwards
    backward = backwards
    reverse = invert

*/
