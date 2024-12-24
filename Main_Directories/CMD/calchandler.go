package CMD

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result,omitempty"` //если хочешь можешь убрать omitempty,но учти что поля если будут пустыми они будут отображатьс
	Error  string `json:"error,omitempty"`
}

// это обработчик для калькулятора
func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	//проверка является ли запрос post-ом,а не get-ом например
	if r.Method != http.MethodPost {
		http.Error(w, "я же тебе сказал,через всякие get нельзя", http.StatusUnprocessableEntity)
		return
	}
	//проверка чтто если не получилось задекодировать тело ответа то выдаем ошибку
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid character", http.StatusUnprocessableEntity) //если что лучше http.StatusBadRequest и другие http статусы лусчше не менять на 440,200,500,просто в первый раз когда я написал код просто с числами у меня выдалась ошибка
		return
	}
	//смотрим на ответ
	result, err := Calc(req.Expression)
	var resp Response
	if err != nil {

		if strings.Contains(err.Error(), "invalid character") {
			resp.Error = "проверь еще раз твое выражение"
			w.WriteHeader(http.StatusUnprocessableEntity)
		} else {
			resp.Error = "если ты видишь эту ошибку,поздравляю ,ты умудрился получить ошибку 500"
			w.WriteHeader(http.StatusInternalServerError)
		}

		json.NewEncoder(w).Encode(resp)
		return
	}

	resp.Result = strconv.FormatFloat(result, 'f', -1, 64)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}
