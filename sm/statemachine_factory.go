package sm

//NewStateMachine Factory methods returns a new state machine instance
func NewStateMachine() *StateMachine {
	sm := &StateMachine{
		Current: READY,
		Transition:make(map[TransitionEvent]State,0),
	}

	//TODO: define these GOIDLE, GOINUSE commands somewhere
	sm.SetTransition(READY, IDLE, "GOIDLE")
	sm.SetTransition(IDLE, INUSE, "GOINUSE")
	sm.SetTransition(IDLE, PAUSED, "TIMEOUT")
	sm.SetTransition(INUSE, PAUSED, "TIMEOUT")
	sm.SetTransition(INUSE, FINISHED, "GOFINISHED")
	sm.SetTransition(FINISHED, READY, "TIMEOUT")
	sm.SetTransition(PAUSED, FINISHED, "GOFINISHED")
	sm.SetTransition(PAUSED, FINISHED, "TIMEOUT")
	sm.SetTransition(IDLE, MANUAL, "MANUAL")
	sm.SetTransition(MANUAL, IDLE, "TIMEOUT")
	sm.SetTransition(MANUAL, IDLE, "COMPLETE")
	sm.SetTransition(FINISHED, IDLE, "GOIDLE")
	sm.SetTransition(HAVEID, INUSE, "GOINUSE")
	sm.SetTransition(HAVEID, IDLE, "GOIDLE")
	sm.SetTransition(IDLE, HAVEID, "GOHAVEID")

	return sm
}
