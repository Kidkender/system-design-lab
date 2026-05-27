package service

import (
	"context"
	"testing"

	"github.com/kidkender/system-design-lab/internal/db"
	"github.com/stretchr/testify/mock"
)

type MockScenarioQuerier struct {
	mock.Mock
}

func (m *MockScenarioQuerier) CreateScenario(
	ctx context.Context,
	arg db.CreateScenarioParams,
) (db.Scenario, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(db.Scenario), args.Error(1)
}

func TestCreateScenario_Success(t *testing.T) {
	// mockS := new(MockScenarioQuerier)
	// service := NewScenarioService(mockS)

}
