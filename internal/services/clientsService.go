package services

import (
	"fmt"

	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/dto"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/interfaces"
	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/models"
)

type ClientsService struct {
	clientRepository interfaces.ClientRepository
	jwtLogin         interfaces.JsonWebToken
}

func NewClientsService(clientRepository interfaces.ClientRepository, jwtLogin interfaces.JsonWebToken) *ClientsService {
	return &ClientsService{clientRepository: clientRepository, jwtLogin: jwtLogin}
}

func (s *ClientsService) Create(data dto.CreateClient) (*models.Client, error) {
	client, err := s.clientRepository.Create(data)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (s *ClientsService) Login(data dto.CreateLoginDTO) (interface{}, error) {
	clientExist := s.clientRepository.GetByEmail(data.Email)

	if clientExist == nil {
		return nil, fmt.Errorf("Email or Password incorrect")
	}

	token, err := s.jwtLogin.GenerateToken(clientExist.ID)

	if err != nil {
		return nil, err
	}

	return token, nil
}
