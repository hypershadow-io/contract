package hook

import "sync"

// AddInitiator registers a new Kind as an initiator kind.
// Initiator kinds are used to identify the source or trigger type of an operation (e.g. UI, System).
func AddInitiator(kind Kind) {
	locker.Lock()
	initiatorKinds = append(initiatorKinds, kind)
	locker.Unlock()
}

var (
	// initiatorKinds holds the list of known initiator kinds (e.g., UI, System, Silent).
	// These are used to identify the origin of a request or action.
	initiatorKinds = []Kind{KindUI, KindSystem, KindSilent}
	locker         sync.RWMutex
)

// PickInitiatorKinds returns a new Kinds set containing only the kinds
// from the current set that match known initiator kinds (e.g., UI, System, Silent).
func (a Kinds) PickInitiatorKinds() Kinds {
	locker.RLock()
	result := a.Pick(initiatorKinds...)
	locker.RUnlock()
	return result
}
