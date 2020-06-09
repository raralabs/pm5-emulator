package sm

import (
	"fmt"
	"pm5-emulator/config"
)

type inUseState struct{
	statemachine *StateMachine
}

func(r inUseState) getStateName()string{
	return config.PM5_STATE_INUSE
}


func(r inUseState) update(command byte) error{
	if command==config.CSAFE_GOFINISHED_CMD{
		r.statemachine.SetState(config.PM5_STATE_FINISHED)
		return nil
	}
	//todo: handle workout cancel and timeout issues
	return fmt.Errorf("undefined command type %v",command)
}

