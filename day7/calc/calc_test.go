package calc

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	n := Add(10, 20)
	assert.Equal(t, 30, n)
}

func TestSub(t *testing.T) {
	n := Sub(10, 20)
	if n != -10 {
		t.Errorf("Expect -10 but got %d", n)
	}
}

func TestMul(t *testing.T) {
	n := Mul(10, 20)
	if n != 200 {
		t.Errorf("Expect 200 but got %d", n)
	}
}

func TestDiv(t *testing.T) {
	n := Div(10, 20)
	if n != 0.5 {
		t.Errorf("Expect 0.5 but got %f", n)
	}
}

func ExampleAdd() {
	n := Add(10, 20)
	fmt.Println(n)
	//Output:
	//30
}

func TestAddHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "http://test.com/add?a=10&b=20", nil)
	w := httptest.NewRecorder()

	AddHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	assert.Equal(t, `{"result":30}`, string(body))
}
