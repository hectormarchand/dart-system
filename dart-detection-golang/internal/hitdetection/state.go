package hitdetection

// State is a per-camera detection state, tracking a candidate blob from
// its first appearance to a confirmed, settled hit.
type State int

const (
	// StateIdle: no significant change from the background.
	StateIdle State = iota
	// StateMotion: a blob has appeared but hasn't been stable long enough
	// to be distinguished from a hand/arm passing through frame.
	StateMotion
	// StateSettled: the blob has stayed in place for the configured
	// number of frames and is reported as a CandidateHit.
	StateSettled
)

func (s State) String() string {
	panic("not implemented")
}
