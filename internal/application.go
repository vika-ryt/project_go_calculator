package application

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
    "io"
	"github.com/vika-ryt/project_go_calculator/pkg/calculation"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}
type Request struct {
	Expression string`json:"expression"`
 }
type Res struct {
	Result float64`json:"result"`
}
type Mistakes struct {
	Error error`json:"error"`
}
// Функция запуска приложения
// тут будем чиать введенную строку и после нажатия ENTER писать результат работы программы на экране
// если пользователь ввел exit - то останаваливаем приложение
func (a *Application) Run() error {

		//вычисляем выражение
	result, err := calculation.Calc(request.Expression)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
			//log.Println({"error": "wrong requestion"})
		answer := Mistakes{Error: err}
		json.NewEncoder(w).Encode(answer)
	} else {
		answer := Res{Result: result}
		json.NewEncoder(w).Encode(answer)
			//log.Println(text, "=", result)
	}
}


//type Request struct {
	//Expression string `json:"expression"`
//}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
    
	result, err := calculation.Calc(request.Expression)
	if err != nil {
		if errors.Is(err, calculation.ErrInvalidExpression) {
			answer := Mistakes{Error: "Expression is not valid"}
		    json.NewEncoder(w).Encode(answer)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			answer := Mistakes{Error: "Internal server error"}
		    json.NewEncoder(w).Encode(answer)
			w.WriteHeader(http.StatusInternalServerError)
		}

	} else {
		answer := {Result: result}
		json.NewEncoder(w).Encode(answer)
		http.Error(w, err.Error(), http.StatusOK)
	}
	w.WriteHeader(http.StatusOK)
}

func (a *Application) RunServer() error {
	http.HandleFunc("/", CalcHandler)
	return http.ListenAndServe(":"+a.config.Addr, nil)
}