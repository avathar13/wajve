package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"

	"wajve/internal/metrics"
)

var errNotSupportedParameterName = fmt.Errorf("this parameter name is not supported")

func (h *Handler) Get(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	ctx := request.Context()
	supportedParameterNames := map[string]struct{}{
		"text":   {},
		"number": {},
		"found":  {},
		"type":   {},
	}
	filterMap := make(map[string]string)

	for name, value := range request.URL.Query() {
		if _, ok := supportedParameterNames[name]; !ok {
			h.writeRequestError(
				writer,
				http.StatusBadRequest,
				"supportedParameterNames.error",
				fmt.Errorf("%w: %s", errNotSupportedParameterName, name),
			)

			return
		}

		filterMap[name] = value[0]
	}

	samples, err := h.samplesService.Get(ctx, filterMap)
	if err != nil {
		h.writeRequestError(writer, http.StatusInternalServerError, "samplesService.Get.error", err)

		return
	}

	samplesBytes, err := json.Marshal(samples)
	if err != nil {
		h.writeRequestError(writer, http.StatusInternalServerError, "samples.Marshal.error: %w", err)

		return
	}

	if _, err := writer.Write(samplesBytes); err != nil {
		h.writeRequestError(writer, http.StatusInternalServerError, "samples.write.error", err)
	}
}

func (h *Handler) Populate(writer http.ResponseWriter, request *http.Request) {
	rawBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		h.writeRequestError(writer, http.StatusInternalServerError, "cannot read request body", err)

		return
	}

	var body PopulateRequestBody
	if err = json.Unmarshal(rawBody, &body); err != nil {
		h.writeRequestError(writer, http.StatusInternalServerError, "cannot unmarshal request body", err)

		return
	}

	if err = h.samplesService.Populate(request.Context(), body.Path); err != nil {
		h.writeRequestError(writer, http.StatusInternalServerError, "samplesService.Populate.error", err)

		return
	}

	_, _ = writer.Write([]byte("OK"))
}

func (h *Handler) writeRequestError(writer http.ResponseWriter, code int, msg string, err error) {
	h.logger.Error(msg, zap.Error(err))
	metrics.Error(msg)

	writer.WriteHeader(code)

	requestError := RequestError{
		Error: err.Error(),
	}

	requestErrorBytes, err := json.Marshal(requestError)
	if err != nil {
		h.logger.Error("json.Marshal.error: %w", zap.Error(err))

		return
	}

	if _, err = writer.Write(requestErrorBytes); err != nil {
		h.logger.Error("request.write.error: %w", zap.Error(err))
	}
}
