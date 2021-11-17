// handlers_test.go
package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestHostTelemetryHandler(t *testing.T) {
	storageAdapterMock := new(MockStorageAdapter)
	storageAdapterMock.On("StoreHostTelemetry", mock.Anything).Return(nil)

	body, _ := json.Marshal(&HostTelemetry{
		AgentID:       "agent_id",
		SLESVersion:   "15.2",
		CPUCount:      16,
		SocketCount:   32,
		TotalMemoryMB: 32000,
		CloudProvider: "azure",
		Time:          time.Now(),
	})
	req, err := http.NewRequest("POST", "/api/collect/hosts", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(hostTelemetryHandler(storageAdapterMock))
	handler.ServeHTTP(rec, req)

	storageAdapterMock.AssertExpectations(t)
	assert.Equal(t, http.StatusAccepted, rec.Code)
}