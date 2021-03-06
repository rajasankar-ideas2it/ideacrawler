// Code generated by cdpgen. DO NOT EDIT.

package headlessexperimental

import (
	"github.com/mafredri/cdp/protocol/runtime"
)

// BeginFrameArgs represents the arguments for BeginFrame in the HeadlessExperimental domain.
type BeginFrameArgs struct {
	FrameTime        *runtime.Timestamp `json:"frameTime,omitempty"`        // Timestamp of this BeginFrame (milliseconds since epoch). If not set, the current time will be used unless frameTicks is specified.
	FrameTimeTicks   *float64           `json:"frameTimeTicks,omitempty"`   // Timestamp of this BeginFrame in Renderer TimeTicks (milliseconds of uptime). If not set, the current time will be used unless frameTime is specified.
	Deadline         *runtime.Timestamp `json:"deadline,omitempty"`         // Deadline of this BeginFrame (milliseconds since epoch). If not set, the deadline will be calculated from the frameTime and interval unless deadlineTicks is specified.
	DeadlineTicks    *float64           `json:"deadlineTicks,omitempty"`    // Deadline of this BeginFrame in Renderer TimeTicks (milliseconds of uptime). If not set, the deadline will be calculated from the frameTime and interval unless deadline is specified.
	Interval         *float64           `json:"interval,omitempty"`         // The interval between BeginFrames that is reported to the compositor, in milliseconds. Defaults to a 60 frames/second interval, i.e. about 16.666 milliseconds.
	NoDisplayUpdates *bool              `json:"noDisplayUpdates,omitempty"` // Whether updates should not be committed and drawn onto the display. False by default. If true, only side effects of the BeginFrame will be run, such as layout and animations, but any visual updates may not be visible on the display or in screenshots.
	Screenshot       *ScreenshotParams  `json:"screenshot,omitempty"`       // If set, a screenshot of the frame will be captured and returned in the response. Otherwise, no screenshot will be captured. Note that capturing a screenshot can fail, for example, during renderer initialization. In such a case, no screenshot data will be returned.
}

// NewBeginFrameArgs initializes BeginFrameArgs with the required arguments.
func NewBeginFrameArgs() *BeginFrameArgs {
	args := new(BeginFrameArgs)

	return args
}

// SetFrameTime sets the FrameTime optional argument. Timestamp of
// this BeginFrame (milliseconds since epoch). If not set, the current
// time will be used unless frameTicks is specified.
func (a *BeginFrameArgs) SetFrameTime(frameTime runtime.Timestamp) *BeginFrameArgs {
	a.FrameTime = &frameTime
	return a
}

// SetFrameTimeTicks sets the FrameTimeTicks optional argument.
// Timestamp of this BeginFrame in Renderer TimeTicks (milliseconds of
// uptime). If not set, the current time will be used unless frameTime
// is specified.
func (a *BeginFrameArgs) SetFrameTimeTicks(frameTimeTicks float64) *BeginFrameArgs {
	a.FrameTimeTicks = &frameTimeTicks
	return a
}

// SetDeadline sets the Deadline optional argument. Deadline of this
// BeginFrame (milliseconds since epoch). If not set, the deadline will
// be calculated from the frameTime and interval unless deadlineTicks
// is specified.
func (a *BeginFrameArgs) SetDeadline(deadline runtime.Timestamp) *BeginFrameArgs {
	a.Deadline = &deadline
	return a
}

// SetDeadlineTicks sets the DeadlineTicks optional argument. Deadline
// of this BeginFrame in Renderer TimeTicks (milliseconds of uptime).
// If not set, the deadline will be calculated from the frameTime and
// interval unless deadline is specified.
func (a *BeginFrameArgs) SetDeadlineTicks(deadlineTicks float64) *BeginFrameArgs {
	a.DeadlineTicks = &deadlineTicks
	return a
}

// SetInterval sets the Interval optional argument. The interval
// between BeginFrames that is reported to the compositor, in
// milliseconds. Defaults to a 60 frames/second interval, i.e. about
// 16.666 milliseconds.
func (a *BeginFrameArgs) SetInterval(interval float64) *BeginFrameArgs {
	a.Interval = &interval
	return a
}

// SetNoDisplayUpdates sets the NoDisplayUpdates optional argument.
// Whether updates should not be committed and drawn onto the display.
// False by default. If true, only side effects of the BeginFrame will
// be run, such as layout and animations, but any visual updates may
// not be visible on the display or in screenshots.
func (a *BeginFrameArgs) SetNoDisplayUpdates(noDisplayUpdates bool) *BeginFrameArgs {
	a.NoDisplayUpdates = &noDisplayUpdates
	return a
}

// SetScreenshot sets the Screenshot optional argument. If set, a
// screenshot of the frame will be captured and returned in the
// response. Otherwise, no screenshot will be captured. Note that
// capturing a screenshot can fail, for example, during renderer
// initialization. In such a case, no screenshot data will be returned.
func (a *BeginFrameArgs) SetScreenshot(screenshot ScreenshotParams) *BeginFrameArgs {
	a.Screenshot = &screenshot
	return a
}

// BeginFrameReply represents the return values for BeginFrame in the HeadlessExperimental domain.
type BeginFrameReply struct {
	HasDamage      bool   `json:"hasDamage"`                // Whether the BeginFrame resulted in damage and, thus, a new frame was committed to the display. Reported for diagnostic uses, may be removed in the future.
	ScreenshotData []byte `json:"screenshotData,omitempty"` // Base64-encoded image data of the screenshot, if one was requested and successfully taken.
}

// EnterDeterministicModeArgs represents the arguments for EnterDeterministicMode in the HeadlessExperimental domain.
type EnterDeterministicModeArgs struct {
	InitialDate *float64 `json:"initialDate,omitempty"` // Number of seconds since the Epoch
}

// NewEnterDeterministicModeArgs initializes EnterDeterministicModeArgs with the required arguments.
func NewEnterDeterministicModeArgs() *EnterDeterministicModeArgs {
	args := new(EnterDeterministicModeArgs)

	return args
}

// SetInitialDate sets the InitialDate optional argument. Number of
// seconds since the Epoch
func (a *EnterDeterministicModeArgs) SetInitialDate(initialDate float64) *EnterDeterministicModeArgs {
	a.InitialDate = &initialDate
	return a
}
