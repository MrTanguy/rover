package communication

import (
	"github.com/MrTanguy/rover/missionControl"
	"github.com/MrTanguy/rover/rover"
)

type CommunicationImpl struct {
	rover          rover.Rover
	missionControl missionControl.MissionControl
}
