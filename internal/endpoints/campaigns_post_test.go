package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danilocordeirodev/go-email/internal/contract"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type serviceMock struct {
	mock.Mock
}

func (s *serviceMock) Create(newCampaignReq contract.NewCampaignReq) (string, error) {
	args := s.Called(newCampaignReq)
	return args.String(0), args.Error(1)
}

func Test_CampaignsPost_should_save_new_campaign(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaignReq{
		Name:    "teste",
		Content: "Hi content",
		Emails:  []string{"test@test.com"},
	}
	service := new(serviceMock)
	service.On("Create", mock.MatchedBy(func(request contract.NewCampaignReq) bool {
		if request.Name == body.Name {
			return true
		} else {
			return false
		}
	})).Return("56x", nil)
	handler := Handler{CampaignService: service}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignPost(rr, req)

	assert.Equal(201, status)
	assert.Nil(err)

}

func Test_CampaignsPost_should_inform_error_when_exists(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaignReq{
		Name:    "teste",
		Content: "Hi content",
		Emails:  []string{"test@test.com"},
	}
	service := new(serviceMock)
	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))
	handler := Handler{CampaignService: service}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)
	rr := httptest.NewRecorder()

	_, _, err := handler.CampaignPost(rr, req)

	assert.NotNil(err)

}
