package env

import (
	"fmt"
	"os"
	"time"

	"github.com/stackrox/rox/pkg/utils/panic"
)

// DurationSetting represents an environment variable which should be parsed into a duration
type DurationSetting struct {
	envVar          string
	defaultDuration time.Duration
	opts            durationSettingOpts
}

// EnvVar returns the string name of the environment variable
func (d *DurationSetting) EnvVar() string {
	return d.envVar
}

// Setting returns the string form of the duration environment variable
func (d *DurationSetting) Setting() string {
	return d.DurationSetting().String()
}

// DurationSetting returns the Duration object represented by the environment variable
func (d *DurationSetting) DurationSetting() time.Duration {
	val := os.Getenv(d.envVar)
	if val != "" {
		dur, err := time.ParseDuration(val)
		if err == nil && validateDuration(dur, d.opts) == nil {
			return dur
		}
	}
	return d.defaultDuration
}

func registerDurationSetting(envVar string, defaultDuration time.Duration, options ...DurationSettingOption) *DurationSetting {
	var opts durationSettingOpts
	for _, o := range options {
		o.apply(&opts)
	}

	if err := validateDuration(defaultDuration, opts); err != nil {
		panic.HardPanic(fmt.Sprintf("%+v", err))
	}

	s := &DurationSetting{
		envVar:          envVar,
		defaultDuration: defaultDuration,
		opts:            opts,
	}

	Settings[s.EnvVar()] = s
	return s
}

func validateDuration(d time.Duration, opts durationSettingOpts) error {
	if d < 0 {
		return fmt.Errorf("invalid duration: %v < 0", d)
	}
	if !opts.zeroAllowed && d == 0 {
		return fmt.Errorf("invalid duration: %v == 0", d)
	}

	return nil
}
