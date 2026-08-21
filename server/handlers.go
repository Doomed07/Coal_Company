package server

import (
	"Coal_company/factory"
	"Coal_company/factory/items"
	"Coal_company/factory/miners"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type HTTPHandlers struct {
	Factory *factory.Factory
}

func NewHTTPHandlers(factory *factory.Factory) *HTTPHandlers {
	return &HTTPHandlers{
		Factory: factory,
	}
}

func (h *HTTPHandlers) writeErr(w http.ResponseWriter, err error, code int) {
	errDTO := ErrorDTO{
		Message: err.Error(),
		Time:    time.Now(),
	}
	http.Error(w, errDTO.ToString(), code)
}

func (h *HTTPHandlers) Output(w http.ResponseWriter, t any) {
	b, err := json.MarshalIndent(t, "", "    ")
	if err != nil {
		fmt.Println("Fail to convertation information into JSON")
		h.writeErr(w, err, http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("Fail to write information into http response")
		h.writeErr(w, err, http.StatusInternalServerError)
		return
	}
}

// Main handlers

func (h *HTTPHandlers) AboutFactory(w http.ResponseWriter, r *http.Request) {
	h.Output(w, h.Factory.Allinfo())
}

func (h *HTTPHandlers) StopFactory(w http.ResponseWriter, r *http.Request) {
	h.Factory.Stop()
	w.WriteHeader(http.StatusNoContent)
}

// Miners" handlers

func (h *HTTPHandlers) ShowMiners(w http.ResponseWriter, r *http.Request) {
	h.Output(w, h.Factory.MinersCostInfo())
}

func (h *HTTPHandlers) BuyMiner(w http.ResponseWriter, r *http.Request) {
	var minerClassDTO ClassDTO
	if err := json.NewDecoder(r.Body).Decode(&minerClassDTO); err != nil {
		h.writeErr(w, err, http.StatusBadRequest)
		return
	}

	classInterfaceDTO, err := h.Factory.Miners.ClassByName(minerClassDTO.Class)
	if err != nil {
		if errors.Is(err, miners.ErrorUnknownClass) {
			h.writeErr(w, err, http.StatusBadRequest)
			return
		} else {
			h.writeErr(w, err, http.StatusInternalServerError)
			return
		}
	}

	if err := h.Factory.BuyMiner(classInterfaceDTO); err != nil {
		if errors.Is(err, factory.ErrorNotEnoughCoal) {
			h.writeErr(w, err, http.StatusConflict)
			return
		} else {
			h.writeErr(w, err, http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	h.Output(w, classInterfaceDTO)
}

func (h *HTTPHandlers) WorkingMiners(w http.ResponseWriter, r *http.Request) {
	h.Output(w, h.Factory.WorkingMiners())
}

func (h *HTTPHandlers) AllMiners(w http.ResponseWriter, r *http.Request) {
	h.Output(w, h.Factory.AllMiners())
}

//Items' handlers

func (h *HTTPHandlers) BuyItem(w http.ResponseWriter, r *http.Request) {
	var item ItemDTO
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		h.writeErr(w, err, http.StatusBadRequest)
		return
	}

	typeItemDTO, err := h.Factory.Items.ItemByName(item.Item)
	if err != nil {
		if errors.Is(err, items.ErrorUnknownItem) {
			h.writeErr(w, err, http.StatusBadRequest)
			return
		} else {
			h.writeErr(w, err, http.StatusInternalServerError)
			return
		}
	}

	if err := h.Factory.BuyItem(typeItemDTO); err != nil {
		if errors.Is(err, factory.ErrorNotEnoughCoal) {
			h.writeErr(w, err, http.StatusConflict)
			return
		} else {
			h.writeErr(w, err, http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	h.Output(w, typeItemDTO)
}

func (h *HTTPHandlers) AllItems(w http.ResponseWriter, r *http.Request) {
	h.Output(w, h.Factory.InfoItems())
}

func (h *HTTPHandlers) AllBoughtItems(w http.ResponseWriter, r *http.Request) {
	h.Output(w, h.Factory.AllBoughtItems())
}
