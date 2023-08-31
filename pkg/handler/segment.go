package handler

import (
	"encoding/json"
	"log"
	"net/http"

	dynamicsegmentation "github.com/RCNRC/dynamic_segmentation"
)

func (h *Handler) createSegment(w http.ResponseWriter, r *http.Request) {
	var err error

	segment := dynamicsegmentation.Segment{}
	err = json.NewDecoder(r.Body).Decode(&segment)
	if err != nil {
		log.Fatalf("error while unmarshalling segment: %s", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["success"] = "false"

	if err = h.services.Segment.CreateSegment(segment.Title); err != nil {
		log.Fatalf("error while insrteing segment: %s", err.Error())
		resp["error"] = "cannot create segment"
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		resp["success"] = "true"
		resp["new_segment"] = segment.Title
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

func (h *Handler) deleteSegment(w http.ResponseWriter, r *http.Request) {
	var err error

	segment := dynamicsegmentation.Segment{}
	err = json.NewDecoder(r.Body).Decode(&segment)
	if err != nil {
		log.Fatalf("error while unmarshalling segment: %s", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["success"] = "false"

	if err = h.services.Segment.DeleteSegment(segment.Title); err != nil {
		log.Fatalf("error while deleting segment: %s", err.Error())
		resp["error"] = "cannot delete segment"
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		resp["success"] = "true"
		resp["deleted_segment"] = segment.Title
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
