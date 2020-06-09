package sm

var (
	READY    State = "ready"
	OFFLINE  State = "offline"
	IDLE     State = "idle"
	MANUAL   State = "manual"
	INUSE    State = "inuse"
	FINISHED State = "finished"
	HAVEID   State = "haveid"
	PAUSED   State = "paused"
)


func (s *StateMachine) Reset() {
	s.SetState(READY)
}

func (s *StateMachine) IsIdle() bool {
	if s.GetState() == IDLE {
		return true
	}
	return false
}

func (s *StateMachine) HaveID() bool {
	if s.GetState()==HAVEID {
		return true
	}
	return false
}

func (s *StateMachine) IsFinished() bool {
	if s.GetState() == FINISHED {
		return true
	}
	return false
}

func (s *StateMachine) IsReady() bool {
	if s.GetState() == READY {
		return true
	}
	return false
}

