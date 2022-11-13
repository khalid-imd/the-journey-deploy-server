package handlers

import (
	journeydto "backend-journey/dto/journey"
	dto "backend-journey/dto/result"
	"backend-journey/models"
	"backend-journey/repositories"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerJourney struct {
	JourneyRepository repositories.JourneyRepository
}

func HandlerJourney(JourneyRepository repositories.JourneyRepository) *handlerJourney {
	return &handlerJourney{JourneyRepository}
}

func (h *handlerJourney) CreateJourney(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["user_id"].(float64))

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	request := journeydto.CreateJourneyRequest{
		Title:        r.FormValue("title"),
		Descriptions: r.FormValue("description"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	journey := models.Journey{
		UserId:      userId,
		Title:       request.Title,
		Image:       filename,
		Description: request.Descriptions,
	}

	data, err := h.JourneyRepository.CreateJourney(journey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerJourney) FindJourneys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	journeys, err := h.JourneyRepository.FindJourneys()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	for i, p := range journeys {
		journeys[i].Image = os.Getenv("PATH_FILE") + p.Image
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: journeys}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerJourney) GetJourney(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	journey, err := h.JourneyRepository.GetJourney(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	journey.Image = os.Getenv("PATH_FILE") + journey.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: journey}
	json.NewEncoder(w).Encode(response)
}

func convertResponseJourney(u models.Journey) journeydto.JourneyResponse {
	return journeydto.JourneyResponse{
		ID:          u.ID,
		Title:       u.Title,
		UserId:      u.User.ID,
		Description: u.Description,
	}
}
