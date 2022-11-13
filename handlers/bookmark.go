package handlers

import (
	bookmarkdto "backend-journey/dto/bookmark"
	dto "backend-journey/dto/result"
	"backend-journey/models"
	"backend-journey/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerBookmark struct {
	BookmarkRepository repositories.BookmarkRepository
}

func HandlerBookmark(BookmarkRepository repositories.BookmarkRepository) *handlerBookmark {
	return &handlerBookmark{BookmarkRepository}
}

func (h *handlerBookmark) CreateBookmark(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["user_id"].(float64))

	var request bookmarkdto.BookmarkRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	bookmark := models.Bookmark{
		UserId:    userId,
		JourneyId: request.JourneyId,
	}

	data, err := h.BookmarkRepository.CreateBookmark(bookmark)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	bookmark, _ = h.BookmarkRepository.GetBookmark(data.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseBookmark(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBookmark) FindBookmarks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bookmarks, err := h.BookmarkRepository.FindBookmarks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	// for i, p := range bookmarks {
	// 	bookmarks[i].Image = os.Getenv("PATH_FILE") + p.Image
	// }

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: bookmarks}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBookmark) GetBookmark(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var bookmark models.Bookmark

	bookmark, err := h.BookmarkRepository.GetBookmark(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: bookmark}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBookmark) DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	bookmark, err := h.BookmarkRepository.GetBookmark(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.BookmarkRepository.DeleteBookmark(bookmark)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseBookmark(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseBookmark(u models.Bookmark) models.Bookmark {
	return models.Bookmark{
		ID:        u.ID,
		UserId:    u.UserId,
		User:      u.User,
		JourneyId: u.JourneyId,
		Journey:   u.Journey,
	}
}
