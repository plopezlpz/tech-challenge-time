package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
	"time-tracker-be/internal/mocks"
)

func TestGetRecords(t *testing.T) {
	tests := []struct {
		name string
		path string
		wantedInvocation string
		wantedArg int64
		wantedStatus int
	}{
		{
			name: "OK",
			path: "/records?start=10",
			wantedInvocation: "List",
			wantedArg: 10,
			wantedStatus: http.StatusOK,
		},
		{
			name: "OK zero",
			path: "/records?start=0",
			wantedInvocation: "List",
			wantedArg: 0,
			wantedStatus: http.StatusOK,
		},
		{
			name: "Not OK negative",
			path: "/records?start=-1",
			wantedStatus: http.StatusBadRequest,
		},
		{
			name: "Not OK missing 'start'",
			path: "/records",
			wantedStatus: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a response recorder
			w := httptest.NewRecorder()
			r := gin.Default()

			recordsStore := mocks.NewRecordStoreMock()
			// Routes setup
			r.GET("/records", GetRecords(&recordsStore))

			// Create a request to send to the above route
			req, _ := http.NewRequest("GET", tt.path, nil)
			r.ServeHTTP(w, req)

			if w.Code != tt.wantedStatus {
				t.Errorf("Should return %v status", tt.wantedStatus)
			}
			if tt.wantedStatus >= 400 {
				return
			}
			if len(recordsStore.Invocations[tt.wantedInvocation]) != 1 {
				t.Errorf("Should have invoked 'List'")
			}
			arg, ok := recordsStore.Invocations[tt.wantedInvocation][0].(int64)
			if !ok {
				t.Errorf("cannot cast argument of '%v(int64)' to int64", tt.wantedInvocation)
			}
			if arg != tt.wantedArg {
				t.Errorf("should have called %v(%v) on the store", tt.wantedInvocation, tt.wantedArg)
			}

		})
	}
}