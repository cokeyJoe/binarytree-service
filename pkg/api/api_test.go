package api

import (
	"binarytree/pkg/tree/binary"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_getIntFromRequest(t *testing.T) {

	req1, _ := http.NewRequest(http.MethodGet, "localhost", nil)
	q1 := req1.URL.Query()
	q1.Add("val", "5")
	req1.URL.RawQuery = q1.Encode()

	req2, _ := http.NewRequest(http.MethodGet, "localhost", nil)
	q2 := req2.URL.Query()
	q1.Add("vadd", "5")
	req2.URL.RawQuery = q2.Encode()

	req3, _ := http.NewRequest(http.MethodGet, "localhost", nil)
	q3 := req3.URL.Query()
	q3.Add("val", "5ss")
	req3.URL.RawQuery = q3.Encode()

	type args struct {
		r   *http.Request
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "valid query, must return 5",
			args:    args{r: req1, key: "val"},
			want:    5,
			wantErr: false,
		},
		{
			name:    "has no requered query, must return err",
			args:    args{r: req2, key: "val"},
			want:    0,
			wantErr: true,
		},
		{
			name:    "invalid query, must return err",
			args:    args{r: req3, key: "val"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getIntFromRequest(tt.args.r, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("getIntFromRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getIntFromRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_searchHandler(t *testing.T) {

	rww := make([]*httptest.ResponseRecorder, 3)
	rww[0] = &httptest.ResponseRecorder{}
	rww[1] = &httptest.ResponseRecorder{}
	rww[2] = &httptest.ResponseRecorder{}

	req1, _ := http.NewRequest(http.MethodGet, "localhost", nil)
	q1 := req1.URL.Query()
	q1.Add("vals", "6")
	req1.URL.RawQuery = q1.Encode()

	req2, _ := http.NewRequest(http.MethodGet, "localhost", nil)
	q2 := req2.URL.Query()
	q2.Add("val", "56")
	req2.URL.RawQuery = q2.Encode()

	req3, _ := http.NewRequest(http.MethodGet, "localhost", nil)
	q3 := req3.URL.Query()
	q3.Add("val", "5")
	req3.URL.RawQuery = q3.Encode()

	type fields struct {
		bstFinder  Finder
		wantStatus int
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "request has no valid query, must return bad request(400)",
			args:   args{rww[0], req1},
			fields: fields{bstFinder: &test1Finder{}, wantStatus: 400},
		},

		{
			name:   "cant find bst value, must return not found(404)",
			args:   args{rww[1], req2},
			fields: fields{bstFinder: &test1Finder{}, wantStatus: 404},
		},

		{
			name:   "bst returns value, must return OK(200) and value in body",
			args:   args{rww[2], req3},
			fields: fields{bstFinder: &test1Finder{}, wantStatus: 200},
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				bstFinder: tt.fields.bstFinder,
			}
			api.searchHandler(tt.args.w, tt.args.r, nil)

			if rww[i].Code != tt.fields.wantStatus {
				t.Errorf("expected status code %d, got %d", tt.fields.wantStatus, rww[i].Code)
			}
		})
	}
}

func TestAPI_deleteHandler(t *testing.T) {

	rww := make([]*httptest.ResponseRecorder, 2)
	rww[0] = &httptest.ResponseRecorder{}
	rww[1] = &httptest.ResponseRecorder{}

	req1, _ := http.NewRequest(http.MethodDelete, "localhost", nil)
	q1 := req1.URL.Query()
	q1.Add("vals", "6")
	req1.URL.RawQuery = q1.Encode()

	req2, _ := http.NewRequest(http.MethodDelete, "localhost", nil)
	q2 := req2.URL.Query()
	q2.Add("val", "5")
	req2.URL.RawQuery = q2.Encode()

	type fields struct {
		bstRemover Remover
		wantStatus int
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "request has no valid query, must return bad request(400)",
			args:   args{rww[0], req1},
			fields: fields{bstRemover: &test1Finder{}, wantStatus: 400},
		},

		{
			name:   "cant find bst value, must return OK(200)",
			args:   args{rww[1], req2},
			fields: fields{bstRemover: &test1Finder{}, wantStatus: 200},
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				bstRemover: tt.fields.bstRemover,
			}
			api.deleteHandler(tt.args.w, tt.args.r, nil)

			if rww[i].Code != tt.fields.wantStatus {
				t.Errorf("expected status code %d, got %d", tt.fields.wantStatus, rww[i].Code)
			}
		})
	}
}

type test1Finder struct {
}

func (t *test1Finder) Find(value int) *binary.TreeNode {
	if value == 5 {
		return &binary.TreeNode{}
	}
	return nil
}

func (t *test1Finder) Remove(int) {}

func Test_getValueFromRequest(t *testing.T) {
	t.Run("has no body, must return error", func(t *testing.T) {

		buf := bytes.NewBuffer([]byte{})
		req, _ := http.NewRequest("POST", "localhost", buf)
		_, err := getValueFromRequest(req)
		if err == nil {
			t.Errorf("expected error, got nil")
		}

	})

	t.Run("body has invalid int field, must return error", func(t *testing.T) {
		var fakeBody struct {
			Val string `json:"value"`
		}
		fakeBody.Val = "555"

		bb, _ := json.Marshal(&fakeBody)

		buf := bytes.NewBuffer(bb)
		req, _ := http.NewRequest("POST", "localhost", buf)
		_, err := getValueFromRequest(req)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("body has valid int, must return no error", func(t *testing.T) {
		var fakeBody struct {
			Val int `json:"value"`
		}
		fakeBody.Val = 555

		bb, _ := json.Marshal(&fakeBody)

		buf := bytes.NewBuffer(bb)
		req, _ := http.NewRequest("POST", "localhost", buf)
		_, err := getValueFromRequest(req)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
	})
}

func TestAPI_insertHandler(t *testing.T) {

	t.Run("request is not valid, must return status 400", func(t *testing.T) {

		rr := &httptest.ResponseRecorder{}

		inserter := &test2Bst{}

		api := &API{
			bstInserter: inserter,
		}
		buf := bytes.NewBuffer([]byte{1, 2, 3})
		req1, _ := http.NewRequest(http.MethodPost, "localhost", buf)

		api.insertHandler(rr, req1, nil)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("request is valid, must return status 200, record valid value", func(t *testing.T) {

		expectedVal := 555

		rr := &httptest.ResponseRecorder{}

		inserter := &test2Bst{}

		api := &API{
			bstInserter: inserter,
		}

		var fakeBody struct {
			Val int `json:"value"`
		}
		fakeBody.Val = expectedVal

		bb, _ := json.Marshal(&fakeBody)

		buf := bytes.NewBuffer(bb)
		req1, _ := http.NewRequest(http.MethodPost, "localhost", buf)

		api.insertHandler(rr, req1, nil)
		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}

		if inserter.value != expectedVal {
			t.Errorf("expected bst inserts value %d, got %d", expectedVal, inserter.value)
		}
	})
}

type test2Bst struct {
	value int
}

func (t *test2Bst) Insert(v int) {
	t.value = v
}

func TestAPI_initHandlers(t *testing.T) {
	t.Run("http handlers must be not nil", func(t *testing.T) {
		api := &API{
			httpServer: &http.Server{},
		}

		api.initHandlers()

		if api.httpServer.Handler == nil {
			t.Error("expected api.httpServer.Handler not nil, got nil")
		}
	})
}
