package main

//если есть какие то вопросы пишите в мой лс в тг: @java_shcolo
import (
	"Main_Directories/CMD"
	"bytes"

	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	//создаается хендлер для калькулятора
	http.HandleFunc("/api/v1/calculate", CMD.CalculateHandler)

	http.ListenAndServe(":8080", nil)

	//это вспомогательное сообщение
	fmt.Println("сервер с портом 8080 был запущен(знай это)")
	//url по которому post запрос будет делать запрос
	url := "http://localhost:8080/api/v1/calculate"
	//переменная для хранения запроса
	var req CMD.Request
	//сам запрос
	req.Expression = "2+2*(6*6)"
	//маршалим пример для отправки на сервак
	jsonreq, err := json.Marshal(req)
	//обработка
	if err != nil {
		fmt.Println("что то пошло не так при кодировании в json")
	}
	//сам запрос
	as, err1 := http.Post(url, "application/json/json", bytes.NewBuffer(jsonreq))
	if err1 != nil {
		fmt.Println("при отправке запрома что то пошло не так")
	}
	//после выхода из области функции закрываем тело запроса
	defer as.Body.Close()
	//проверка что запрос выаолнился успешно
	if as.StatusCode == http.StatusOK {
		fmt.Println("запрос выполнен успешно")
	} else {
		fmt.Println("запрос провалился")
	}
	//читаем тело ответа для вывода на экран
	Respo, errorq := io.ReadAll(as.Body)
	if errorq != nil {
		fmt.Println("случилась какая то ошибка")
	}
	//я думаю 2 эти строки можно не обьяснять
	fmt.Println("вот что сервер выдал:")
	fmt.Println(string(Respo))
}
