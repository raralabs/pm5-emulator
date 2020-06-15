package sm

import (
	"fmt"
	"pm5-emulator/config"
)

type manualState struct {
	statemachine *StateMachine
}

func (m manualState) getStateName() string {
	return config.PM5_STATE_MANUAL
}

func (m manualState) update(command byte) error {
	if command == config.CSAFE_GOIDLE_CMD {
		m.statemachine.SetState(config.PM5_STATE_IDLE)
		return nil
	}
	//todo: handle timeout case
	return fmt.Errorf("undefined command type %v", command)
}
