package rest_test

import (
	"net/http"
	"testing"

	"github.com/gastonbordet/remembrall/pkg/http/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockEventClient struct {
	mock.Mock
}

func (m *MockEventClient) GetAll(w http.ResponseWriter, r *http.Request)       {}
func (m *MockEventClient) GetByEventID(w http.ResponseWriter, r *http.Request) {}

func TestInitRouter(t *testing.T) {
	// Setup

	// Act
	router := rest.InitRouter(&MockEventClient{})

	// Assert
	assert.NotNil(t, router)
}
