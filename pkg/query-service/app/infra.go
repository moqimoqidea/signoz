package app

import (
	"encoding/json"
	"net/http"

	"go.signoz.io/signoz/pkg/query-service/model"
)

func (aH *APIHandler) getHostAttributeKeys(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req, err := parseFilterAttributeKeyRequest(r)

	if err != nil {
		RespondError(w, &model.ApiError{Typ: model.ErrorInternal, Err: err}, nil)
		return
	}

	// get attribute keys
	keys, err := aH.hostsRepo.GetHostAttributeKeys(ctx, *req)
	if err != nil {
		RespondError(w, &model.ApiError{Typ: model.ErrorInternal, Err: err}, nil)
		return
	}

	// write response
	aH.Respond(w, keys)
}

func (aH *APIHandler) getHostAttributeValues(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// parse request
	req, err := parseFilterAttributeValueRequest(r)

	if err != nil {
		RespondError(w, &model.ApiError{Typ: model.ErrorInternal, Err: err}, nil)
		return
	}

	// get attribute values
	values, err := aH.hostsRepo.GetHostAttributeValues(ctx, *req)
	if err != nil {
		RespondError(w, &model.ApiError{Typ: model.ErrorInternal, Err: err}, nil)
		return
	}

	// write response
	aH.Respond(w, values)
}

func (aH *APIHandler) getHostList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := model.HostListRequest{}

	// parse request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		RespondError(w, &model.ApiError{Typ: model.ErrorInternal, Err: err}, nil)
		return
	}

	// get host list
	hostList, err := aH.hostsRepo.GetHostList(ctx, req)
	if err != nil {
		RespondError(w, &model.ApiError{Typ: model.ErrorInternal, Err: err}, nil)
		return
	}

	// write response
	aH.Respond(w, hostList)
}

func (aH *APIHandler) getProcessAttributeKeys(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req, err := parseFilterAttributeKeyRequest(r)
	if err != nil {
		RespondError(w, &model.ApiError{Typ: model.ErrorInternal, Err: err}, nil)
		return
	}

	keys, err := aH.processesRepo.GetProcessAttributeKeys(ctx, *req)
	if err != nil {
		RespondError(w, &model.ApiError{Typ: model.ErrorInternal, Err: err}, nil)
		return
	}

	aH.Respond(w, keys)
}

func (aH *APIHandler) getProcessAttributeValues(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req, err := parseFilterAttributeValueRequest(r)
	if err != nil {
		RespondError(w, &model.ApiError{Typ: model.ErrorInternal, Err: err}, nil)
		return
	}

	values, err := aH.processesRepo.GetProcessAttributeValues(ctx, *req)
	if err != nil {
		RespondError(w, &model.ApiError{Typ: model.ErrorInternal, Err: err}, nil)
		return
	}

	aH.Respond(w, values)
}

func (aH *APIHandler) getProcessList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := model.ProcessListRequest{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		RespondError(w, &model.ApiError{Typ: model.ErrorInternal, Err: err}, nil)
		return
	}

	hostList, err := aH.processesRepo.GetProcessList(ctx, req)
	if err != nil {
		RespondError(w, &model.ApiError{Typ: model.ErrorInternal, Err: err}, nil)
		return
	}

	aH.Respond(w, hostList)
}