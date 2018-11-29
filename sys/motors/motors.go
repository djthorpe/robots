package motors

import (
	"context"
	"fmt"
	"time"

	// Frameworks
	"github.com/djthorpe/gopi"
	"github.com/djthorpe/robots"
)

////////////////////////////////////////////////////////////////////////////////
// TYPES

type Motors struct {
	PWM    gopi.PWM
	Period time.Duration
}

type motors struct {
	pwm    gopi.PWM
	period time.Duration
	log    gopi.Logger
	motors []robots.Motor
}

////////////////////////////////////////////////////////////////////////////////
// OPEN AND CLOSE

// Open creates a new Motors object
func (config Motors) Open(log gopi.Logger) (gopi.Driver, error) {
	log.Debug("<sys.robots.Motors>Open{ Period=%v PWM=%v }", config.Period, config.PWM)

	// Check incoming parameters
	if config.PWM == nil {
		return nil, gopi.ErrBadParameter
	}

	// create new PWM driver
	this := new(motors)

	// Set logging
	this.log = log
	this.pwm = config.PWM
	this.period = config.Period
	this.motors = make([]robots.Motor, 0)

	// success
	return this, nil
}

// Close Motor object
func (this *motors) Close() error {
	this.log.Debug("<sys.robots.Motors>Close{}")

	// Stop all motors
	if err := this.Stop(); err != nil {
		return err
	}

	// Release resources
	this.pwm = nil
	this.motors = nil

	return nil
}

////////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *motors) String() string {
	return fmt.Sprintf("<sys.robots.Motors>{ pwm=%v period=%v motors=%v }", this.pwm, this.period, this.motors)
}

////////////////////////////////////////////////////////////////////////////////
// IMPLEMENTATION

func (this *motors) Add(minus, plus gopi.GPIOPin, invert bool) (robots.Motor, error) {
	this.log.Debug2("<sys.robots.Motors>Add{ minus=%v plus=%v invert=%v }", minus, plus, invert)
	if plus == gopi.GPIO_PIN_NONE || minus == gopi.GPIO_PIN_NONE {
		return nil, gopi.ErrBadParameter
	}
	// Create a motor object
	motor := NewMotor(plus, minus, invert)
	if motor == nil {
		return nil, gopi.ErrBadParameter
	}
	// Set period
	if this.period != 0 {
		if err := this.pwm.SetPeriod(this.period, plus); err != nil {
			return nil, err
		}
		if err := this.pwm.SetPeriod(this.period, minus); err != nil {
			return nil, err
		}
	}
	// Set duty cycle to zero (off)
	if err := this.pwm.SetDutyCycle(0, plus); err != nil {
		return nil, err
	}
	if err := this.pwm.SetDutyCycle(0, minus); err != nil {
		return nil, err
	}
	// Append motor and return success
	this.motors = append(this.motors, motor)
	return motor, nil
}

func (this *motors) Run(ctx context.Context, speed float32, motors ...robots.Motor) error {
	this.log.Debug2("<sys.robots.Motors>Run{ speed=%v motors=%v }", speed, motors)
	if ctx == nil {
		return gopi.ErrBadParameter
	}
	return gopi.ErrNotImplemented
}

func (this *motors) Stop(motors ...robots.Motor) error {
	this.log.Debug2("<sys.robots.Motors>Cancel{ motors=%v }", motors)

	// All motors if no argument is provided
	if len(motors) == 0 {
		motors = this.motors
	}

	// Cancel all
	for _, m := range motors {
		if err := this.stop_motor(m.(*motor)); err != nil {
			return err
		}
	}

	// Return success
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *motors) stop_motor(motor *motor) error {
	if motor.cancel != nil {
		motor.cancel()
		// TODO: do not return from here until cancel is completed
	}
	return nil
}

func (this *motors) start_motor(motor *motor, speed float32) error {
	return gopi.ErrNotImplemented
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
