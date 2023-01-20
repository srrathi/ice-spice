package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/srrathi/ice-spice/models"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/station_data", r.GetStationData)
	api.Get("/analysis_data", r.GetAnalysisData)
}

// Function to fetch station baltic sea ice data with the analysis data of that station
func (r *Repository) GetStationData(context *fiber.Ctx) error {
	stationsModels := &[]models.StationData{}
	analysisModels := &[]models.AnalysisData{}
	stationsDataResponse := []models.StationDataResponse{}
	analysisDataResponse := []models.AnalysisDataResponse{}
	stationCode := context.Query("stationCode")

	if stationCode == "" {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "stationCode query param cannot be empty",
		})
		return nil
	}

	err := r.DB.Table("stations_data").Where("station_code= ?", stationCode).Find(stationsModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get stations"})
		return err
	}

	if len(*stationsModels) == 0 {
		context.Status(http.StatusOK).JSON(
			&fiber.Map{"message": "No data found for this metro station among metro stations in Finland"})
		return err
	}

	for _, station := range *stationsModels {
		restrictedDaysModel := &models.RestrictedDays{}
		err := r.DB.Table("restricted_days").Where("rd_id= ?", *station.RestrictedDays).Find(restrictedDaysModel).Error
		if err != nil {
			context.Status(http.StatusBadRequest).JSON(
				&fiber.Map{"message": "could not get stations"})
			return err
		}

		var maxThickness float32
		if station.MaxThickness != nil {
			maxThickness = *station.MaxThickness
		}

		rds := []int{*restrictedDaysModel.RD1, *restrictedDaysModel.RD2, *restrictedDaysModel.RD3, *restrictedDaysModel.RD4,
			*restrictedDaysModel.RD5, *restrictedDaysModel.RD6, *restrictedDaysModel.RD7, *restrictedDaysModel.RD8, *restrictedDaysModel.RD9}

		stationData := models.StationDataResponse{
			Beginning:      *station.Beginning,
			End:            *station.End,
			SeasonLength:   *station.SeasonLength,
			IceDays:        *station.IceDays,
			FlTotal:        *station.FlTotal,
			VolSum:         *station.VolSum,
			MaxThickness:   maxThickness,
			MaxThickness2:  *station.MaxThickness2,
			Country:        *station.Country,
			CountryCode:    *station.CountryCode,
			StationName:    *station.StationName,
			StationCode:    *station.StationCode,
			RestrictedDays: rds,
		}
		stationsDataResponse = append(stationsDataResponse, stationData)
	}

	err = r.DB.Table("analysis_data").Where("station_code= ?", stationCode).Find(analysisModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get stations"})
		return err
	}

	for _, stationAnalysis := range *analysisModels {
		analysisResponse := models.AnalysisDataResponse{
			AnalysisBasis: *stationAnalysis.AnalysisBasis,
			Beginning:     *stationAnalysis.Beginning,
			End:           *stationAnalysis.End,
			IceDays:       *stationAnalysis.IceDays,
			FlTotal:       *stationAnalysis.FlTotal,
			VolSum:        *stationAnalysis.VolSum,
			Country:       *stationAnalysis.Country,
			CountryCode:   *stationAnalysis.CountryCode,
			StationName:   *stationAnalysis.StationName,
			StationCode:   *stationAnalysis.StationCode,
		}
		analysisDataResponse = append(analysisDataResponse, analysisResponse)
	}

	if len(*analysisModels) == 0 {
		context.Status(http.StatusOK).JSON(
			&fiber.Map{"message": "No data found for this metro station among metro stations in Finland"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "metro station data fetched successfully",
		"data": &fiber.Map{
			"stations-data":         stationsDataResponse,
			"analysis-data-station": analysisDataResponse,
		},
	})
	return nil
}

// Function to fetch only analysis data of that station
func (r *Repository) GetAnalysisData(context *fiber.Ctx) error {
	analysisModels := &[]models.AnalysisData{}
	analysisDataResponse := []models.AnalysisDataResponse{}
	stationCode := context.Query("stationCode")

	if stationCode == "" {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "stationCode query param cannot be empty",
		})
		return nil
	}

	err := r.DB.Table("analysis_data").Where("station_code= ?", stationCode).Find(analysisModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get analysis data for the station"})
		return err
	}

	for _, stationAnalysis := range *analysisModels {
		analysisResponse := models.AnalysisDataResponse{
			AnalysisBasis: *stationAnalysis.AnalysisBasis,
			Beginning:     *stationAnalysis.Beginning,
			End:           *stationAnalysis.End,
			IceDays:       *stationAnalysis.IceDays,
			FlTotal:       *stationAnalysis.FlTotal,
			VolSum:        *stationAnalysis.VolSum,
			Country:       *stationAnalysis.Country,
			CountryCode:   *stationAnalysis.CountryCode,
			StationName:   *stationAnalysis.StationName,
			StationCode:   *stationAnalysis.StationCode,
		}
		analysisDataResponse = append(analysisDataResponse, analysisResponse)
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "metro station analysis data fetched successfully",
		"data": &fiber.Map{
			"analysis-data-station": analysisDataResponse,
		},
	})
	return nil
}
