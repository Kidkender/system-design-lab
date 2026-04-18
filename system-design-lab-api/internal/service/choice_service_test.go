package service

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/kidkender/system-design-lab/internal/db"
	"github.com/kidkender/system-design-lab/internal/handler/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockChoiceQuerier struct {
	mock.Mock
}

func (m *MockChoiceQuerier) CreateChoice(ctx context.Context, arg db.CreateChoiceParams) (db.Choice, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(db.Choice), args.Error(1)
}

func TestCreateChoice_Success(t *testing.T) {
	mockQ := new(MockChoiceQuerier)
	service := NewChoiceService(mockQ)

	stepID := uuid.New()
	nextStepUUID := uuid.New()

	nextStep := nextStepUUID.String()
	req := &dto.CreateChoiceRequest{
		StepID:     stepID.String(),
		NextStepID: &nextStep,
		Label:      "Optiona A",
		IsCorrect:  true,
	}

	mockQ.On("CreateChoice", mock.Anything, mock.Anything).
		Return(db.Choice{
			ID:         uuid.New(),
			StepID:     stepID,
			Label:      req.Label,
			IsCorrect:  true,
			NextStepID: nextStepUUID,
		}, nil)

	resp, err := service.CreateChoice(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, req.Label, resp.Label)
	assert.Equal(t, true, resp.IsCorrect)
	assert.Equal(t, stepID.String(), resp.ID)
	assert.NotNil(t, resp.NextStepID)

	mockQ.AssertExpectations(t)
}

func TestCreateChoice_InvalidStepID(t *testing.T) {
	mockQ := new(MockChoiceQuerier)
	service := NewChoiceService(mockQ)

	req := &dto.CreateChoiceRequest{
		StepID:    "invalid-uuid",
		Label:     "Option A",
		IsCorrect: false,
	}

	resp, err := service.CreateChoice(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestCreateChoice_InvalidNextStepID(t *testing.T) {
	mockQ := new(MockChoiceQuerier)
	service := NewChoiceService(mockQ)

	stepID := uuid.New()
	invalid := "not-a-valid"
	req := &dto.CreateChoiceRequest{
		StepID:     stepID.String(),
		Label:      "Option A",
		IsCorrect:  false,
		NextStepID: &invalid,
	}

	resp, err := service.CreateChoice(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestCreateChoice_DBError(t *testing.T) {
	mockQ := new(MockChoiceQuerier)
	service := NewChoiceService(mockQ)

	stepID := uuid.New()

	req := &dto.CreateChoiceRequest{
		StepID:    stepID.String(),
		Label:     "Option A",
		IsCorrect: false,
	}

	mockQ.
		On("CreateChoice", mock.Anything, mock.Anything).
		Return(db.Choice{}, errors.New("db error"))

	resp, err := service.CreateChoice(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, resp)

	mockQ.AssertExpectations(t)
}

func TestCreateChoice_Table(t *testing.T) {
	tests := []struct {
		name      string
		req       *dto.CreateChoiceRequest
		setupMock func(m *MockChoiceQuerier)
		wantErr   bool
	}{
		{
			name: "success",
			req: &dto.CreateChoiceRequest{
				StepID:    uuid.New().String(),
				Label:     "A",
				IsCorrect: true,
			},
			setupMock: func(m *MockChoiceQuerier) {
				m.On("CreateChoice", mock.Anything, mock.Anything).
					Return(db.Choice{
						ID:        uuid.New(),
						StepID:    uuid.New(),
						Label:     "A",
						IsCorrect: true,
					}, nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockQ := new(MockChoiceQuerier)
			if tt.setupMock != nil {
				tt.setupMock(mockQ)
			}

			service := NewChoiceService(mockQ)

			_, err := service.CreateChoice(context.Background(), tt.req)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
