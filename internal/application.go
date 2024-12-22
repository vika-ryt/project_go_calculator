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

	//a := make([]string, 3)
	//var w []string
	//for {
		// читаем выражение для вычисления из командной строки
		//log.Println("input expression")
		//reader := bufio.NewReader(os.Stdin)
		//for { 
			//line, err := reader.ReadString('\n') 

			//if err != nil { 
			//	if err == io.EOF { 
				//	break
				//} else { 
				//	fmt.Println(err) 
				//} 
			//} 
			//w = append(w, line)
			
			//fmt.Print(line) 
		//}
		
		//t := strings.Split(w[1], " ")
		//text := string(t[1])
		// убираем пробелы, чтобы оставить только вычислемое выражение
		//text = strings.TrimSpace(text)
		// выходим, если ввели команду "exit"
		//if text == "exit" {
			//log.Println("aplication was successfully closed")
			//return nil
		//}
		//вычисляем выражение
	result, err := calculation.Calc(request.Expression)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
			//log.Println({"error": "wrong requestion"})
		answer := Mistakes{Error: err}
		json.NewEncoder(w).Encode(answer)
	} else {
		answer := {Result: result}
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