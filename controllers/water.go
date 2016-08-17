package controllers

import (  
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/julienschmidt/httprouter"
)

type (
    WaterController struct{}
)

func NewWaterController() * WaterController{
    return &WaterController{}

    
}

func (wc WaterController) StartWater(w http.ResponseWriter){
    
}