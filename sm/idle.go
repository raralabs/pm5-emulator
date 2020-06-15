package sm

import (
	"fmt"
	"pm5-emulator/config"
)

type idleState struct {
	statemachine *StateMachine
}

func (r idleState) getStateName() string {
	return config.PM5_STATE_IDLE
}

func (r idleState) update(command byte) error {
	if command == config.CSAFE_GOINUSE_CMD {
		r.statemachine.SetState(config.PM5_STATE_INUSE)
		return nil
	} else if command == config.CSAFE_GOHAVEID_CMD {
		r.statemachine.SetState(config.PM5_STATE_HAVEID)
		return nil
	}
	//todo: goto paused state due to timeout
	return fmt.Errorf("undefined command type")
}
