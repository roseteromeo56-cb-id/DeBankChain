package httputil

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWrappedResponseWriterWriteHeaderIsIdempotent(t *testing.T) {
	recorder := httptest.NewRecorder()
	writer := NewWrappedResponseWriter(recorder)

	writer.WriteHeader(http.StatusNotFound)
	writer.WriteHeader(http.StatusInternalServerError)

	require.Equal(t, http.StatusNotFound, writer.StatusCode)
	require.Equal(t, http.StatusNotFound, recorder.Code)
}
