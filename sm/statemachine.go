package sm

type Command []byte

//StateMachine implements PM5 states
type StateMachine struct{

}


func (s *StateMachine) SetStatus(cmd Command){

}

func (s *StateMachine) Reset(cmd Command){}

func (s *StateMachine) IsIdle(cmd Command){}

func (s *StateMachine) HaveID(cmd Command){}

func (s *StateMachine) IsFinished(cmd Command){}

func (s *StateMachine) IsReady(cmd Command){}

func(s *StateMachine) SetTime(cmd Command){}

func (s *StateMachine) SetDate(cmd Command){}

func(s *StateMachine) SetProgram(cmd Command){}