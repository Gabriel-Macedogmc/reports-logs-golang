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
	bcrypt           interfaces.BcryptInterface
}

func NewClientsService(clientRepository interfaces.ClientRepository, jwtLogin interfaces.JsonWebToken, bcrypt interfaces.BcryptInterface) *ClientsService {
	return &ClientsService{clientRepository: clientRepository, jwtLogin: jwtLogin, bcrypt: bcrypt}
}

func (s *ClientsService) Create(data dto.CreateClient) (*models.Client, error) {
	password, err := s.bcrypt.HashPassword(data.Password)

	if err != nil {
		return nil, err
	}

	client, err := s.clientRepository.Create(dto.CreateClient{
		Password: password,
		Name:     data.Name,
		Email:    data.Email,
		Document: data.Document,
	})

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

	if !s.bcrypt.CheckPasswordHash(data.Password, clientExist.Password) {
		return nil, fmt.Errorf("Email or Password incorrect")
	}

	token, err := s.jwtLogin.GenerateToken(clientExist.ID)

	if err != nil {
		return nil, err
	}

	return token, nil
}
