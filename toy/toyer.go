package toy

import (
	"github.com/andrewbruno/go-toyrobot/compass"
	"github.com/andrewbruno/go-toyrobot/unit"
)

// Toyer definition
type Toyer interface {
	Update(unit.X, unit.Y, compass.Direction)
}
