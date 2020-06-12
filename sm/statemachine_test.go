package sm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"pm5-emulator/config"
	"testing"
)

func TestSetState(t *testing.T){
	sm:=NewStateMachine()
	sm.SetState(config.PM5_STATE_READY)
	currentState:=sm.GetState().getStateName()
	if currentState!=config.PM5_STATE_READY{
		t.Errorf("statemachine state, got: %s, want: %s",currentState,config.PM5_STATE_READY)
	}
}

func TestReset(t *testing.T){
	sm:=NewStateMachine()
	sm.SetState(config.PM5_STATE_READY)

	sm.Reset()  //reset statemachine
	currentState:=sm.GetState().getStateName()

	if currentState!=config.PM5_STATE_READY{
		t.Errorf("statemachine state, got: %s, want: %s",currentState,config.PM5_STATE_READY)
	}
}


func TestUpdate(t *testing.T){
	sm:=NewStateMachine()
	sm.SetState(config.PM5_STATE_READY)

	err:=sm.Update(config.CSAFE_GOFINISHED_CMD)  //should fail
	if assert.Error(t,err){
		assert.Equal(t,fmt.Errorf("undefined command"),err)
	}


	err=sm.Update(config.CSAFE_GOIDLE_CMD) //should pass
	if assert.NoError(t,err){
		assert.Equal(t,nil,err)
	}

	sm.Reset() //reset statemachine

	//walk through all the states
	stepUpdate:=[]struct{
		path string
		command byte

	}{
		{"ready2idle",config.CSAFE_GOIDLE_CMD},
		{"idle2haveID",config.CSAFE_GOHAVEID_CMD},
		{"haveID2InUse",config.CSAFE_GOINUSE_CMD},
		{"inUse2Finished",config.CSAFE_GOFINISHED_CMD},
	}

	for _,step :=range stepUpdate{
		t.Run(step.path,func(t *testing.T){
			err=sm.Update(step.command)
			if assert.NoError(t,err){
				assert.Equal(t,nil,err)
			}
		})
	}
}

func TestIsIdle(t *testing.T){
	sm:=NewStateMachine()
	sm.SetState(config.PM5_STATE_READY)  //set to ready state

	if sm.IsIdle(){
		t.Errorf("statemachine idle state, got: %v, want: %s",sm.IsIdle(),"false")
	}
}

func TestHaveID(t *testing.T){
	sm:=NewStateMachine()
	sm.SetState(config.PM5_STATE_HAVEID)

	if !sm.HaveID(){
		t.Errorf("statemachine HaveID state, got: %v, want: %s",sm.HaveID(),"true")
	}
}

func TestIsFinished(t *testing.T){
	sm:=NewStateMachine()
	sm.SetState(config.PM5_STATE_FINISHED)
	if !sm.IsFinished(){
		t.Errorf("statemachine IsFinished state, got: %v, want: %s",sm.IsFinished(),"true")
	}
}

func TestIsReady(t *testing.T){
	sm:=NewStateMachine()
	sm.SetState(config.PM5_STATE_READY)
	if !sm.IsReady(){
		t.Errorf("statemachine ready state, got %v, want %s",sm.IsReady(),"true")
	}
}
