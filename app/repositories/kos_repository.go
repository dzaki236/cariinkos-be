package repositories

import (
	"errors"
	"strings"

	"cariinkos/app/facades"
	"cariinkos/app/helpers"
	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"

	"gorm.io/gorm"
)

type KosRepository struct{}

func NewKosRepository() *KosRepository {
	return &KosRepository{}
}

func (r *KosRepository) GetAllData(page, perPage int) ([]models.Kos, int64, error) {
	return r.getData(page, perPage, false)
}

func (r *KosRepository) GetPublishedData(page, perPage int) ([]models.Kos, int64, error) {
	return r.getData(page, perPage, true)
}

func (r *KosRepository) getData(page, perPage int, publishedOnly bool) ([]models.Kos, int64, error) {
	var dataKos []models.Kos
	var total int64
	query := facades.Orm().Query().With("User")
	if publishedOnly == true {
		query = query.Where("status = ?", "published")
	}
	err := query.Order("created_at desc").Paginate(page, perPage, &dataKos, &total)
	if err != nil {
		return nil, 0, err
	}
	if err := r.getAvgRatingsAndFavoriteCountForKos(dataKos); err != nil {
		return nil, 0, err
	}
	for index := range dataKos {
		var photo models.KosPhoto
		err := facades.Orm().Query().Where("kos_id = ?", dataKos[index].ID).Order("created_at asc").Order("id asc").First(&photo)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			baseURL := strings.TrimRight(facades.Config().GetString("http.url"), "/") + "/public/images_not_found.jpg"
			dataKos[index].CoverPhoto = &baseURL
			continue
		}
		if err != nil {
			return nil, 0, err
		}
		dataKos[index].CoverPhoto = &photo.ImageURL
	}
	return dataKos, total, nil
}

func (r *KosRepository) GetByID(id string) (models.Kos, error) {
	return r.getByID(id, false)
}

func (r *KosRepository) GetPublishedByID(id string) (models.Kos, error) {
	return r.getByID(id, true)
}

func (r *KosRepository) getByID(id string, publishedOnly bool) (models.Kos, error) {
	var kos models.Kos
	query := facades.Orm().Query().Where("id = ?", id)
	if publishedOnly {
		query = query.Where("status = ?", "published")
	}
	err := query.With("User").
		With("Facilities").
		With("Contact").
		With("Photos").
		FirstOrFail(&kos)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Kos{}, helpers.ErrDataNotFound
	}
	if err != nil {
		return models.Kos{}, err
	}
	if err := r.getAvgRatingsAndFavoriteCount(&kos); err != nil {
		return models.Kos{}, err
	}
	return kos, err
}

func (r *KosRepository) getAvgRatingsAndFavoriteCount(kos *models.Kos) error {
	var tablename string = kos.TableName()
	var aggregate struct {
		AverageRating *float64 `gorm:"column:average_rating"`
		FavoriteCount int64    `gorm:"column:favorite_count"`
	}
	err := facades.Orm().Query().
		Table(tablename).Select("kos.id, "+ratingAndFavoriteAggregates).
		Where("kos.id = ?", kos.ID).
		First(&aggregate)
	if err != nil {
		return err
	}
	kos.Rattings = aggregate.AverageRating
	kos.FavoriteCount = aggregate.FavoriteCount
	return nil
}

const ratingAndFavoriteAggregates = "(SELECT AVG(rating) FROM kos_reviews WHERE kos_reviews.kos_id = kos.id) AS average_rating, (SELECT COUNT(*) FROM kos_favorites WHERE kos_favorites.kos_id = kos.id) AS favorite_count"

func (r *KosRepository) getAvgRatingsAndFavoriteCountForKos(kosList []models.Kos) error {
	var tablename string = models.Kos{}.TableName()
	if len(kosList) == 0 {
		return nil
	}
	type aggregate struct {
		KosID         string   `gorm:"column:kos_id"`
		AverageRating *float64 `gorm:"column:average_rating"`
		FavoriteCount int64    `gorm:"column:favorite_count"`
	}
	var aggregates []aggregate
	if err := facades.Orm().Query().Table(tablename).
		Select("kos.id AS kos_id, "+ratingAndFavoriteAggregates).
		Where("kos.id IN ?", kosIDs(kosList)).
		Find(&aggregates); err != nil {
		return err
	}
	statisticsByKosID := make(map[string]aggregate, len(aggregates))
	for _, item := range aggregates {
		statisticsByKosID[item.KosID] = item
	}
	for index := range kosList {
		if item, exists := statisticsByKosID[kosList[index].ID]; exists {
			kosList[index].Rattings = item.AverageRating
			kosList[index].FavoriteCount = item.FavoriteCount
		}
	}
	return nil
}

func kosIDs(kosList []models.Kos) []string {
	ids := make([]string, len(kosList))
	for index := range kosList {
		ids[index] = kosList[index].ID
	}
	return ids
}

func (r *KosRepository) Create(kos *models.Kos) error {
	return facades.Orm().Query().Create(kos)
}

func (r *KosRepository) Update(kos *models.Kos) error {
	_, err := facades.Orm().Query().Where("id = ?", kos.ID).Update(&kos)
	return err
}

func (r *KosRepository) Delete(id string) error {
	var Kos models.Kos
	result, err := facades.Orm().Query().Where("id = ?", id).Delete(&Kos)
	if err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return helpers.ErrDataNotFound
	}
	return nil
}

var _ repositoryContracts.KosRepositoryContract = (*KosRepository)(nil)
