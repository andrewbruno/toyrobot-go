package unit

import "fmt"

// X represents the X placement, also known as the horizontal axis
// It has been set to uint8 (0 to 255) as the requirement is to have 5x5 Units
// However, if we need to change to larger grid, or one that takes in negative
// values, or even floats we can easily change here.  I don't consider this YAGNI
// as its a small investment to future proof.
// It also and makes dealing with X and Y type safe, added bonus.
type X uint8

// Y represents the Y placement, also knows as the vertical axis
type Y uint8

func (u X) String() string {
	return fmt.Sprintf("%d", u)
}

func (u Y) String() string {
	return fmt.Sprintf("%d", u)
}
