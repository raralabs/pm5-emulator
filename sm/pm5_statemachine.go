package sm

import "pm5-emulator/config"

type state interface {
	getStateName() string
	update(command byte) error
}

type StateMachine struct {
	READY    state
	OFFLINE  state
	IDLE     state
	MANUAL   state
	INUSE    state
	FINISHED state
	HAVEID   state
	PAUSED   state

	currentState state
}

func NewStateMachine() *StateMachine {
	pm := &StateMachine{}

	pm.READY = &readyState{statemachine: pm}
	pm.IDLE = &idleState{statemachine: pm}
	pm.HAVEID = &haveIDState{statemachine: pm}
	pm.PAUSED = &pausedState{statemachine: pm}
	pm.MANUAL = &manualState{statemachine: pm}
	pm.INUSE = &inUseState{statemachine: pm}
	pm.FINISHED = &finishedState{statemachine: pm}

	return pm
}

func (sm *StateMachine) GetStateName() string {
	return sm.currentState.getStateName()
}

func (sm *StateMachine) GetState() state {
	return sm.currentState
}

func (sm *StateMachine) SetState(s string) {
	switch s {
	case config.PM5_STATE_IDLE:
		sm.currentState = sm.IDLE
	case config.PM5_STATE_FINISHED:
		sm.currentState = sm.FINISHED
	case config.PM5_STATE_HAVEID:
		sm.currentState = sm.HAVEID
	case config.PM5_STATE_INUSE:
		sm.currentState = sm.INUSE
	case config.PM5_STATE_MANUAL:
		sm.currentState = sm.MANUAL
	case config.PM5_STATE_PAUSED:
		sm.currentState = sm.PAUSED
	default:
		sm.currentState = sm.READY
	}
}

func (sm *StateMachine) Reset() {
	sm.SetState(config.PM5_STATE_READY)
}

func (sm *StateMachine) Update(command byte) error {
	if command == config.CSAFE_RESET_CMD {
		sm.Reset()
		return nil
	}
	return sm.currentState.update(command)
}

func (sm *StateMachine) IsIdle() bool {
	if sm.GetState() == sm.IDLE {
		return true
	}
	return false
}

func (sm *StateMachine) HaveID() bool {
	if sm.GetState() == sm.HAVEID {
		return true
	}
	return false
}

func (sm *StateMachine) IsFinished() bool {
	if sm.GetState() == sm.FINISHED {
		return true
	}
	return false
}

func (sm *StateMachine) IsReady() bool {
	if sm.GetState() == sm.READY {
		return true
	}
	return false
}
