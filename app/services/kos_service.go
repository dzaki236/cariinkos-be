package services

import (
	"cariinkos/app/dto"
	"cariinkos/app/helpers"
	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"
	serviceContracts "cariinkos/app/services/contracts"

	"github.com/google/uuid"
)

type KosService struct {
	repo repositoryContracts.KosRepositoryContract
}

func NewKosService(repo repositoryContracts.KosRepositoryContract) *KosService {
	return &KosService{repo: repo}
}

func (s *KosService) GetAllData(page, perPage int) ([]models.Kos, int64, error) {
	return s.repo.GetAllData(page, perPage)
}

func (s *KosService) GetPublishedData(page, perPage int) ([]models.Kos, int64, error) {
	return s.repo.GetPublishedData(page, perPage)
}

func (s *KosService) GetByID(id string) (models.Kos, error) {
	return s.repo.GetByID(id)
}

func (s *KosService) GetPublishedByID(id string) (models.Kos, error) {
	return s.repo.GetPublishedByID(id)
}

func (s *KosService) Create(userID string, request dto.CreateKosRequest) (models.Kos, error) {
	kos := models.Kos{
		ID:                uuid.NewString(),
		UserID:            userID,
		Name:              request.Name,
		Description:       request.Description,
		Address:           request.Address,
		Province:          request.Province,
		City:              request.City,
		District:          request.District,
		Subdistrict:       request.Subdistrict,
		GmapsLink:         request.GmapsLink,
		PriceMonthly:      request.PriceMonthly,
		Type:              request.Type,
		Status:            request.Status,
		RoomSize:          request.RoomSize,
		TotalRooms:        request.TotalRooms,
		AvailableRooms:    request.AvailableRooms,
		MinimumRentMonths: request.MinimumRentMonths,
		Rules:             request.Rules,
	}
	if kos.Type == "" {
		kos.Type = "mixed"
	}
	if kos.Status == "" {
		kos.Status = "draft"
	}
	if kos.MinimumRentMonths == 0 {
		kos.MinimumRentMonths = 1
	}
	if err := s.repo.Create(&kos); err != nil {
		return models.Kos{}, err
	}
	return kos, nil
}

func (s *KosService) Update(userID, id string, request dto.UpdateKosRequest) (models.Kos, error) {
	kos, err := s.repo.GetByID(id)
	if err != nil {
		return models.Kos{}, err
	}
	if kos.UserID != userID {
		return models.Kos{}, helpers.ErrDataNotFound
	}
	kos.Name = request.Name
	kos.Description = request.Description
	kos.Address = request.Address
	kos.Province = request.Province
	kos.City = request.City
	kos.District = request.District
	kos.Subdistrict = request.Subdistrict
	kos.GmapsLink = request.GmapsLink
	kos.PriceMonthly = request.PriceMonthly
	kos.Type = request.Type
	kos.Status = request.Status
	kos.RoomSize = request.RoomSize
	kos.TotalRooms = request.TotalRooms
	kos.AvailableRooms = request.AvailableRooms
	kos.MinimumRentMonths = request.MinimumRentMonths
	kos.Rules = request.Rules
	if err := s.repo.Update(&kos); err != nil {
		return models.Kos{}, err
	}
	return kos, nil
}

func (s *KosService) Delete(userID, id string) error {
	kos, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if kos.UserID != userID {
		return helpers.ErrDataNotFound
	}
	return s.repo.Delete(id)
}

var _ serviceContracts.KosService = (*KosService)(nil)
