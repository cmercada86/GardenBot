package communication

import (
    "encoding/json"
)

type (
     GardenBotCommunication struct{}
)

type msg struct{
    msgType string
    value int
}

func NewGardenBotCommuncation() *GardenBotCommunication{
    return &GardenBotCommunication{}
}

func (gbc GardenBotCommunication) QueryStatus(){

}

func (gbc GardenBotCommunication) StartWater(timeMin int){

}

func (gbc GardenBotCommunication) StopWater(){
    
}

func (gbc GardenBotCommunication) StartPump(timeMin int){

}
func (gbc GardenBotCommunication) StopPump(timeMin int){

}

func (gbc GardenBotCommunication) sendMsg()