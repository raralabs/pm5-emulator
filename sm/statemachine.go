package sm

type State string

type TransitionEvent struct{
	event string  //event name

	Src State //source state
}

//StateMachine
type StateMachine struct {
	//Current is the state that the FSM is currently in
	Current      State
	//transition holds all possible events and transitions from a given state
	Transition   map[TransitionEvent]State
}

//SetState sets state for SM
func (s *StateMachine) SetState(state State) {
	s.Current = state
}

//GetState returns current state of the SM
func (s *StateMachine) GetState() State {
	return s.Current
}

//SetTransition
func(s *StateMachine) SetTransition(src State,dst State, name string) {
	tmap := TransitionEvent{
		event: name,
		Src:   src,
	}
	s.Transition[tmap] = dst
}

//Update checks whether the given transition is possible or not
//if yes,it updates the state of machine and returns true
//if no, returns false
func (s *StateMachine) Update(event string) bool{
	tmap:=TransitionEvent{
		event: event,
		Src:   s.Current,
	}

	if dst,ok:=s.Transition[tmap];ok{
		s.SetState(dst)  //update current state to destination state
		return true
	}
	return false
}
