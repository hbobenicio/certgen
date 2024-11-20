package certgen

import (
	"certgen/internal/application"
	"certgen/internal/reqstate"
	"crypto/x509/pkix"
	"errors"
	"fmt"
	"log/slog"
	"math/big"
	"net/http"
)

func PostCaHandler(app *application.App) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		slog.InfoContext(req.Context(), "generating new ca...")

		//TODO refactor this to get this from the request payload instead
		caName := pkix.Name{
			SerialNumber:       big.NewInt(2019).String(), //NOTE any problem creating it here? ca also has a field...
			Country:            []string{"BR"},
			Organization:       []string{"My Organization"},
			OrganizationalUnit: []string{"My Organization Unit"},
			CommonName:         "my-root-ca",
		}

		//TODO refactor this to a service
		ca, err := NewCaWithSubject(caName)
		if err != nil {
			httpInternalServerError(rw, req, err)
			return
		}

		//TODO refactor this to a repository
		sql := "INSERT INTO ca (name, pkey_pem, cert_pem) VALUES ($1, $2, $3)"
		dbRes, err := app.DB.ExecContext(req.Context(), sql, caName.CommonName, ca.privateKeyPem.String(), ca.certPem.String())
		if err != nil {
			httpInternalServerError(rw, req, err)
			return
		}
		rowsAffected, err := dbRes.RowsAffected()
		if err != nil {
			httpInternalServerError(rw, req, err)
			return
		}
		if rowsAffected != 1 {
			httpInternalServerError(rw, req, fmt.Errorf("bad rows affected. expected 1, affected %d", rowsAffected))
			return
		}

		recordId, err := dbRes.LastInsertId()
		if err != nil {
			httpInternalServerError(rw, req, fmt.Errorf("failed to parse LastInsertId from db response: %w", err))
			return
		}

		reqState := reqstate.Get(req)
		reqState.Logger.InfoContext(req.Context(), "ca created successfully", "cn", caName.CommonName, "id", recordId)

		newResourceLocation := fmt.Sprintf("%s/%d", req.URL.Path, recordId)

		rw.Header().Add("Location", newResourceLocation)
		rw.WriteHeader(http.StatusCreated)
	}
}

func GetCaPkeyHandler(app *application.App) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		caId := req.PathValue("caId")
		if len(caId) == 0 {
			httpBadRequestError(rw, req, errors.New("missing required path param caId"))
			return
		}

		sql := "SELECT pkey_pem FROM ca WHERE id = $1"
		row := app.DB.QueryRowContext(req.Context(), sql, caId)
		_ = row
	}
}
