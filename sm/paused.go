package sm

import (
	"fmt"
	"pm5-emulator/config"
)

type pausedState struct{
	statemachine *StateMachine
}

func(r pausedState) getStateName()string{
	return config.PM5_STATE_PAUSED
}


func(r pausedState) update(command byte) error{
	if command==config.CSAFE_GOFINISHED_CMD{
		r.statemachine.SetState(config.PM5_STATE_FINISHED)
		return nil
	}
	//TODO: timeout also causes transition from Paused to Finished state
	return fmt.Errorf("undefined command type %v",command)
}

