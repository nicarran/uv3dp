//
// Copyright (c) 2020 Jason S. McMullan <jason.mcmullan@gmail.com>
//

package main

import (
	"github.com/nicarran/uv3dp"
	"github.com/spf13/pflag"
)

type BottomCommand struct {
	*pflag.FlagSet

	Style        string // Style (either 'slow' or 'fade')
	LightOnTime  float32
	LightOffTime float32
	LightPWM     uint8
	LiftHeight   float32
	LiftSpeed    float32
	Count        int
}

func NewBottomCommand() (cmd *BottomCommand) {
	cmd = &BottomCommand{
		FlagSet: pflag.NewFlagSet("bottom", pflag.ContinueOnError),
	}

	cmd.IntVarP(&cmd.Count, "count", "c", 0, "Bottom layer count")
	cmd.StringVarP(&cmd.Style, "style", "y", "slow", "Bottom layer style - 'fade' or 'slow'")
	cmd.Float32VarP(&cmd.LightOnTime, "light-on", "o", 0.0, "Bottom layer light-on time in seconds")
	cmd.Float32VarP(&cmd.LightOffTime, "light-off", "f", 0.0, "Bottom layer light-off time in seconds")
	cmd.Uint8VarP(&cmd.LightPWM, "pwm", "p", 255, "Light PWM rate (0..255)")
	cmd.Float32VarP(&cmd.LiftHeight, "lift-height", "h", 0.0, "Bottom layer lift height in mm")
	cmd.Float32VarP(&cmd.LiftSpeed, "lift-speed", "s", 0.0, "Bottom layer lift speed in mm/min")

	cmd.SetInterspersed(false)

	return
}

type bottomModifier struct {
	uv3dp.Printable
	bottom uv3dp.Bottom
}

func (mod *bottomModifier) Bottom() (bottom uv3dp.Bottom) {
	// Set the bottom exposure
	bottom = mod.bottom

	return
}

func (mod *bottomModifier) LayerExposure(index int) (exposure uv3dp.Exposure) {
	bot := mod.bottom

	if index < bot.Count {
		exposure = bot.Exposure
	} else {
		exposure = mod.Printable.LayerExposure(index)
	}

	return
}

func (cmd *BottomCommand) Filter(input uv3dp.Printable) (output uv3dp.Printable, err error) {
	bot := input.Bottom()

	if cmd.Changed("count") {
		TraceVerbosef(VerbosityNotice, "  Setting default bottom layer count %v", cmd.Count)
		bot.Count = int(cmd.Count)
	}

	if cmd.Changed("light-on") {
		TraceVerbosef(VerbosityNotice, "  Setting default bottom time to %v", cmd.LightOnTime)
		bot.Exposure.LightOnTime = cmd.LightOnTime
	}

	if cmd.Changed("light-off") {
		TraceVerbosef(VerbosityNotice, "  Setting default bottom off time to %v", cmd.LightOffTime)
		bot.Exposure.LightOffTime = cmd.LightOffTime
	}

	if cmd.Changed("pwm") {
		TraceVerbosef(VerbosityNotice, "  Setting default light PWM to %v", cmd.LightPWM)
		bot.Exposure.LightPWM = cmd.LightPWM
	}

	if cmd.Changed("lift-height") {
		TraceVerbosef(VerbosityNotice, "  Setting default bottom lift height to %v", cmd.LightOnTime)
		bot.Exposure.LiftHeight = cmd.LiftHeight
	}

	if cmd.Changed("lift-speed") {
		TraceVerbosef(VerbosityNotice, "  Setting default bottom lift speed to %v", cmd.LightOffTime)
		bot.Exposure.LiftSpeed = cmd.LiftSpeed
	}

	mod := &bottomModifier{
		Printable: input,
		bottom:    bot,
	}

	output = mod

	return
}
