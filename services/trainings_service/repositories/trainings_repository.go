package repositories

import (
	"database/sql"
	"errors"
	"trainings_service/models"
)

// TrainingRepository представляет репозиторий для работы с тренировками
type TrainingRepository struct {
	DB *sql.DB // Подключение к базе данных
}

func NewTrainingRepository(db *sql.DB) *TrainingRepository {
	return &TrainingRepository{DB: db}
}

// GetAllTrainings возвращает все тренировки
func (r *TrainingRepository) GetAllTrainings() ([]*models.Training, error) {
	query := "SELECT id, title, description, rating, favorite FROM trainings"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trainings []*models.Training

	for rows.Next() {
		var training models.Training
		err := rows.Scan(&training.ID, &training.Title, &training.Description, &training.Rating, &training.Favorite)
		if err != nil {
			return nil, err
		}
		trainings = append(trainings, &training)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return trainings, nil
}

// GetTrainingByName возвращает тренировку по ее названию
func (r *TrainingRepository) GetTrainingByName(trainingName string) (*models.Training, error) {
	query := "SELECT id, title, description, rating, favorite FROM trainings WHERE title = $1"
	var training models.Training
	err := r.DB.QueryRow(query, trainingName).Scan(&training.ID, &training.Title, &training.Description, &training.Rating, &training.Favorite)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Возвращаем nil и ошибку, если тренировка не найдена
			return nil, errors.New("тренировка не найдена")
		}
		// Возвращаем ошибку в случае другой ошибки запроса
		return nil, err
	}
	return &training, nil
}

// AddTraining добавляет новую тренировку
func (r *TrainingRepository) AddTraining(training *models.Training) error {
	query := "INSERT INTO trainings (title, description, rating, favorite) VALUES ($1, $2, $3, $4) RETURNING id"
	err := r.DB.QueryRow(query, training.Title, training.Description, training.Rating, training.Favorite).Scan(&training.ID)
	if err != nil {
		// Возвращаем ошибку, если произошла ошибка при выполнении запроса
		return errors.New("ошибка при добавлении тренировки")
	}
	return nil
}
