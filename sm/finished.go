package sm

import (
	"fmt"
	"pm5-emulator/config"
)

type finishedState struct{
	statemachine *StateMachine
}

func(r finishedState) getStateName()string{
	return config.PM5_STATE_FINISHED
}

func(r finishedState) update(command byte) error{
	if command==config.CSAFE_GOIDLE_CMD{
		r.statemachine.SetState(config.PM5_STATE_IDLE)
	}
	//todo: handle timeout case
	return fmt.Errorf("undefined command type %v",command)
}

