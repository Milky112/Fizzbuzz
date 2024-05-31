package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"time"

	"github.com/julienschmidt/httprouter"

	logger "fizzbuzz/log"
	"fizzbuzz/model"
	"fizzbuzz/usecase"
)

func FizzBuzz(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	startTime := time.Now()
	fmt.Println("URL Query")
	fmt.Println(r.URL.Query())

	fromStr := r.FormValue("from")
	toStr := r.FormValue("to")

	fromInt, err := strconv.Atoi(fromStr)
	if err != nil {
		logger.ErrorLog.Print("Failed Unmarshall Input From")
		w.WriteHeader(400)
		responseURL := model.FizzbuzzResult{
			Reason: err.Error(),
		}
		e, _ := json.Marshal(responseURL)
		w.Write(e)
	}
	toInt, err := strconv.Atoi(toStr)
	if err != nil {
		logger.ErrorLog.Print("Failed Unmarshall Input From")
		w.WriteHeader(400)
		responseURL := model.FizzbuzzResult{
			Reason: err.Error(),
		}
		e, _ := json.Marshal(responseURL)
		w.Write(e)
	}

	request := model.FizzBuzzRequest{
		From: fromInt,
		To:   toInt,
	}
	logger.CommonLog.Print("Request : ", request)

	if err = validateRequest(request); err != nil {
		w.WriteHeader(400)
		responseURL := model.FizzbuzzResult{
			Reason: err.Error(),
		}
		logger.ErrorLog.Printf("Error : ", err)
		e, _ := json.Marshal(responseURL)
		w.Write(e)
		return
	}

	responseFizzBuzz := usecase.FizzbuzzCalculate(context.TODO(), request)
	tNow := time.Now()
	processTime := tNow.Sub(startTime)

	responseURL := model.FizzbuzzResult{
		ProcessTime: processTime,
		Result:      responseFizzBuzz,
	}

	logger.CommonLog.Print("Respone : ", responseURL)

	e, _ := json.Marshal(responseURL)
	w.Write(e)
}

func validateRequest(req model.FizzBuzzRequest) error {
	fmt.Println("Validate Request ", req)
	if req.To > 100 || req.From <= 0 {
		fmt.Println("Error woi")
		return fmt.Errorf("invalid request constraint")
	}
	return nil

}
