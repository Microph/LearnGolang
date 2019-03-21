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

func TestSubHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(SubHandler))
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/sub?a=30&b=20", nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	assert.Equal(t, `{"result":10}`, string(body))
}

/*
func TestAddRouter(t *testing.T) {
	SetAddRouter()

	ts := httptest.NewServer(http.DefaultServeMux)
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/add?a=10&b=20", nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	assert.Equal(t, `{"result":30}`, string(body))
}

func TestSubRouter(t *testing.T) {
	SetSubRouter()

	ts := httptest.NewServer(http.DefaultServeMux)
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/sub?a=10&b=20", nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	assert.Equal(t, `{"result":30}`, string(body))
}
*/

//Good practice for multiple routes test (which trigger panic)
func TestRouter(t *testing.T) {
	SetRouter()

	testCases := []struct {
		name   string
		path   string
		result int
	}{
		{"Add", "/add?a=10&b=20", 30},
		{"Sub", "/sub?a=20&b=10", 10},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testRouter(t, tc.path, tc.result)
		})
	}
}

//Trick: use lowercase to prevent GO from automatically run the method for testing
func testRouter(t *testing.T, path string, result int) {
	ts := httptest.NewServer(http.DefaultServeMux)
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+path, nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	expectJSON := fmt.Sprintf(`{"result":%d}`, result)
	assert.Equal(t, expectJSON, string(body))
}
