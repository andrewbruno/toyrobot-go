package toy

import (
	"github.com/andrewbruno/go-toyrobot/src/toyrobot/compass"
	"github.com/andrewbruno/go-toyrobot/src/toyrobot/unit"
)

// Toyer definition
type Toyer interface {
	Update(unit.X, unit.Y, compass.Direction)
}
