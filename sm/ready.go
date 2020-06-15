package sm

import (
	"fmt"
	"pm5-emulator/config"
)

type readyState struct {
	statemachine *StateMachine
}

func (r readyState) getStateName() string {
	return config.PM5_STATE_READY
}

func (r readyState) update(command byte) error {
	if command == config.CSAFE_GOIDLE_CMD {
		r.statemachine.SetState(config.PM5_STATE_IDLE)
		return nil
	} else if command == config.CSAFE_GOINUSE_CMD {
		r.statemachine.SetState(config.PM5_STATE_INUSE)
		return nil
	}
	return fmt.Errorf("undefined command")
}
