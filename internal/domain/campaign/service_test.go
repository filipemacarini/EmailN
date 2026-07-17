package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	newCampaign = contract.NewCampaign{
		Name:    "Teste",
		Content: "Teste2",
		Emails:  []string{"teste1@teste.com"},
	}
	repository = new(repositoryMock)
	service    = ServiceImp{}
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(c *Campaign) error {
	args := r.Called(c)
	return args.Error(0)
}
func (r *repositoryMock) GetBy(id string) (*Campaign, error) {
	args := r.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Campaign), nil
}

func (r *repositoryMock) Get() ([]Campaign, error) {
	return nil, nil
}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repository.On("Save", mock.Anything).Return(nil)
	service.Repository = repository

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_SaveCampaign(t *testing.T) {
	repository.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) ||
			campaign.Status != Pending {
			return false
		}
		return true
	})).Return(nil)
	service.Repository = repository

	service.Create(newCampaign)

	repository.AssertExpectations(t)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaign{})

	assert.False(errors.Is(internalerrors.ErrInternal, err))
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repository = new(repositoryMock)
	repository.On("Save", mock.Anything).Return(errors.New("error to save on database"))
	service.Repository = repository

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalerrors.ErrInternal, err))
}

func Test_GetBy_ReturnCampaign(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	repository = new(repositoryMock)
	repository.On("GetBy", mock.MatchedBy(func(id string) bool {
		return id == campaign.ID
	})).Return(campaign, nil)
	service.Repository = repository

	campaignReturned, _ := service.GetBy(campaign.ID)

	assert.Equal(campaign.ID, campaignReturned.ID)
	assert.Equal(campaign.Name, campaignReturned.Name)
	assert.Equal(campaign.Content, campaignReturned.Content)
	assert.Equal(campaign.Status, campaignReturned.Status)
}

func Test_GetBy_ReturnErrorWhenSomethingWrongExist(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	repository = new(repositoryMock)
	repository.On("GetBy", mock.Anything).Return(nil, errors.New("something wrong"))
	service.Repository = repository

	_, err := service.GetBy(campaign.ID)

	assert.True(errors.Is(internalerrors.ErrInternal, err))
}
