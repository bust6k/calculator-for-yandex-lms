package Tests

import (
	"Main_Directories/CMD"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

type responseToReadAnswer struct {
	Result1 string `json:"result"`
	Error   string `json:"error"`
}

func TestCalcHandler(t *testing.T) {
	TestesSucsess := []struct {
		expression string
		excepted   float64
		wanterr    bool
	}{
		{
			expression: "2+2",
			wanterr:    false,
			excepted:   4,
		},
		{
			expression: "2+4*2",
			wanterr:    false,
			excepted:   10,
		},
		{
			expression: "2/2",
			wanterr:    false,
			excepted:   1,
		},
	}
	for i, _ := range TestesSucsess {
		expr := TestesSucsess[i].expression
		requestBody := strings.NewReader(`{"expression":"` + expr + `"}`)
		req, err := http.NewRequest("POST", "/api/v1/calculate", requestBody)
		if err != nil {
			t.Fatal()
		}
		w := httptest.NewRecorder()
		CMD.CalculateHandler(w, req)

		resp := w.Result()
		defer resp.Body.Close()
		body, er := io.ReadAll(resp.Body)
		if er != nil {
			t.Fatal()
		}
		var responj responseToReadAnswer
		json.Unmarshal(body, &responj)
		f, _ := strconv.ParseFloat(responj.Result1, 6)
		if f == TestesSucsess[i].excepted {
			fmt.Println("все правильно,обьект типа TestSucsess не выдает что то неправильное")
		} else {
			t.Errorf("ожидал %f, а получил %v", TestesSucsess[i].excepted, string(body))
		}
	}
	TestesFail := []struct {
		expression string
		excepted   string
		wanterr    bool
	}{
		{
			expression: "2+2p",
			wanterr:    true,
			excepted:   "проверь еще раз твое выражение",
		},
		{
			expression: "2+4*)",
			wanterr:    true,
			excepted:   "проверь еще раз твое выражение",
		},
		{
			expression: "",
			wanterr:    true,
			excepted:   "проверь еще раз твое выражение",
		},
	}
	for i, _ := range TestesFail {
		expr := TestesFail[i].expression
		requestBody := strings.NewReader(`{"expression":"` + expr + `"}`)
		req, err := http.NewRequest("POST", "/api/v1/calculate", requestBody)
		if err != nil {
			t.Fatal()
		}
		w1 := httptest.NewRecorder()
		CMD.CalculateHandler(w1, req)

		resp := w1.Result()
		defer resp.Body.Close()
		body1, _ := io.ReadAll(resp.Body)
		var respin responseToReadAnswer
		json.Unmarshal(body1, &respin)

		if respin.Error == TestesFail[i].excepted {
			fmt.Println("все правильно")
		} else {
			t.Errorf("ожидал %v, а получил %v", TestesFail[i].excepted, string(body1))
		}
	}
}
