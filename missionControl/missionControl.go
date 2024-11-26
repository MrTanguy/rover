package missionControl

import (
	"github.com/MrTanguy/rover/rover"
)

type MissionControl interface {
	Run()
}

type MissionControlImpl struct {
	rover rover.Rover
}
