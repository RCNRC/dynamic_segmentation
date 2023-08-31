package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	dynamicsegmentation "github.com/RCNRC/dynamic_segmentation"
	"github.com/gorilla/mux"
)

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
	var err error

	update := dynamicsegmentation.UserUpdate{}
	err = json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		log.Fatalf("error while unmarshalling segment: %s", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["success"] = "false"

	if err = h.services.User.Update(update); err != nil {
		log.Fatalf("error while insrteing segment: %s", err.Error())
		resp["error"] = "cannot update users segments"
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		resp["success"] = "true"
	}
	var jsonResp []byte
	if jsonResp, err = json.Marshal(resp); err != nil {
		log.Fatalf("error while marshalling segment: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(jsonResp); err != nil {
		log.Fatalf("error while writing response: %s", err)
		return
	}
}

func (h *Handler) history(w http.ResponseWriter, r *http.Request) {
	var err error

	dates := struct {
		DateFrom string `json:"date_from"`
		DateTo   string `json:"date_to"`
	}{}
	err = json.NewDecoder(r.Body).Decode(&dates)
	if err != nil {
		log.Fatalf("error while unmarshalling segment: %s", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp_seccess := "false"
	resp_error := ""
	var fileNmae string
	url := fmt.Sprintf("%s/user/report/", r.Host)

	scheme := "http"
	if r.TLS != nil {
		scheme = scheme + "s"
	}

	if fileNmae, err = h.services.User.GetUsersSegmentsHistory(dates.DateFrom, dates.DateTo); err != nil {
		log.Fatalf("error while getting segments: %s", err.Error())
		resp_error = "cannot update users segments"
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		resp_seccess = "true"
	}
	var jsonResp []byte
	if jsonResp, err = json.Marshal(struct {
		Success   string `json:"success"`
		Error     string `json:"error"`
		Reference string `json:"reference"`
	}{
		Success:   resp_seccess,
		Error:     resp_error,
		Reference: scheme + "://" + url + fileNmae,
	}); err != nil {
		log.Fatalf("error while marshalling segment: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(jsonResp); err != nil {
		log.Fatalf("error while writing response: %s", err)
		return
	}
}

func (h *Handler) report(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileName := vars["fileName"]
	filePath := h.services.User.GetReportsPath() + "/" + fileName
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error cannot read file: %s", err)
	} else {
		w.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filePath))
		if _, err = w.Write(file); err != nil {
			log.Fatalf("error while writing response: %s", err)
			return
		}
	}
}

func (h *Handler) current(w http.ResponseWriter, r *http.Request) {
	var err error

	user := dynamicsegmentation.User{}
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalf("error while unmarshalling segment: %s", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp_seccess := "false"
	resp_error := ""

	resp_segments, err := h.services.User.GetUsersCurrentSegments(user.Id)
	if err != nil {
		log.Fatalf("error while insrteing segment: %s", err.Error())
		resp_error = "cannot update users segments"
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		resp_seccess = "true"
	}
	var jsonResp []byte
	if jsonResp, err = json.Marshal(struct {
		Success  string   `json:"success"`
		Error    string   `json:"error"`
		UserId   int      `json:"id"`
		Segments []string `json:"segments"`
	}{
		Success:  resp_seccess,
		Error:    resp_error,
		UserId:   user.Id,
		Segments: resp_segments,
	}); err != nil {
		log.Fatalf("error while marshalling segment: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(jsonResp); err != nil {
		log.Fatalf("error while writing response: %s", err)
		return
	}
}
