package service_test

import (
	"errors"
	"library/internal/features/recomendation"
	"library/internal/features/recomendation/service"
	"library/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddRecommend(t *testing.T) {
	mockQry := mocks.NewQuery(t) // Mock untuk recomendation.Query
	recommendSrv := service.NewRecommendServices(mockQry)

	t.Run("Success Add Recommend", func(t *testing.T) {
		newRecommend := recomendation.Recomendation{
			UserID: 1,
			BookID: 1,
			Reason: "Great book",
		}

		// Mocking query behavior
		mockQry.On("AddRecommend", mock.Anything).Return(nil).Once()

		err := recommendSrv.AddRecommend(newRecommend)

		assert.NoError(t, err)
		mockQry.AssertExpectations(t)
	})

	t.Run("Error Add Recommend", func(t *testing.T) {
		newRecommend := recomendation.Recomendation{
			UserID: 1,
			BookID: 1,
			Reason: "Great book",
		}

		// Mocking query behavior with error
		mockQry.On("AddRecommend", mock.Anything).Return(errors.New("query error")).Once()

		err := recommendSrv.AddRecommend(newRecommend)

		assert.Error(t, err)
		assert.Equal(t, "error in server", err.Error())
		mockQry.AssertExpectations(t)
	})
}

func TestGetAllRecommend(t *testing.T) {
	mockQry := mocks.NewQuery(t) // Mock untuk recomendation.Query
	recommendSrv := service.NewRecommendServices(mockQry)

	t.Run("Success Get All Recommend", func(t *testing.T) {
		mockRecommendations := []recomendation.Recomendation{
			{
				ID:     1,
				UserID: 1,
				BookID: 1,
				Reason: "Great book",
			},
			{
				ID:     2,
				UserID: 2,
				BookID: 2,
				Reason: "Inspiring story",
			},
		}

		// Mocking query behavior
		mockQry.On("GetAllRecommend").Return(mockRecommendations, nil).Once()

		result, err := recommendSrv.GetAllRecommend()

		assert.NoError(t, err)
		assert.Equal(t, mockRecommendations, result)
		mockQry.AssertExpectations(t)
	})

	t.Run("Error Get All Recommend", func(t *testing.T) {
		// Mocking query behavior with error
		mockQry.On("GetAllRecommend").Return([]recomendation.Recomendation{}, errors.New("query error")).Once()

		result, err := recommendSrv.GetAllRecommend()

		assert.Error(t, err)
		assert.Equal(t, "error in server", err.Error())
		assert.Equal(t, []recomendation.Recomendation{}, result)
		mockQry.AssertExpectations(t)
	})
}

func TestUpdateRecommend(t *testing.T) {
	mockQry := mocks.NewQuery(t) // Mock untuk recomendation.Query
	recommendSrv := service.NewRecommendServices(mockQry)

	t.Run("Success Update Recommend", func(t *testing.T) {
		recommendID := uint(1)
		updateRecommend := recomendation.Recomendation{
			Reason: "Updated reason",
		}

		// Mocking query behavior
		mockQry.On("UpdateRecommend", recommendID, mock.Anything).Return(nil).Once()

		err := recommendSrv.UpdateRecommend(recommendID, updateRecommend)

		assert.NoError(t, err)
		mockQry.AssertExpectations(t)
	})

	t.Run("Error Update Recommend", func(t *testing.T) {
		recommendID := uint(1)
		updateRecommend := recomendation.Recomendation{
			Reason: "Updated reason",
		}

		// Mocking query behavior with error
		mockQry.On("UpdateRecommend", recommendID, mock.Anything).Return(errors.New("query error")).Once()

		err := recommendSrv.UpdateRecommend(recommendID, updateRecommend)

		assert.Error(t, err)
		assert.Equal(t, "error in server", err.Error())
		mockQry.AssertExpectations(t)
	})
}

func TestDeleteRecommend(t *testing.T) {
	mockQry := mocks.NewQuery(t) // Mock untuk recomendation.Query
	recommendSrv := service.NewRecommendServices(mockQry)

	t.Run("Success Delete Recommend", func(t *testing.T) {
		recommendID := uint(1)

		// Mocking query behavior
		mockQry.On("DeleteRecommend", recommendID).Return(nil).Once()

		err := recommendSrv.DeleteRecommend(recommendID)

		assert.NoError(t, err)
		mockQry.AssertExpectations(t)
	})

	t.Run("Error Delete Recommend", func(t *testing.T) {
		recommendID := uint(1)

		// Mocking query behavior with error
		mockQry.On("DeleteRecommend", recommendID).Return(errors.New("query error")).Once()

		err := recommendSrv.DeleteRecommend(recommendID)

		assert.Error(t, err)
		assert.Equal(t, "error in server", err.Error())
		mockQry.AssertExpectations(t)
	})
}
