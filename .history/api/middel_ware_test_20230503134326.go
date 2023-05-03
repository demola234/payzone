package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/demola234/payzone/db/mock"
	"github.com/demola234/payzone/token"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestAuthMiddleWare(t *testing.T) {
	testCases := []struct {
		name          string
		buildStubs    func(store *mockdb.MockStore)
		setupAuth     func(t *testing.T, request *http.Request, token token.Maker)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			server := newTestServer(t, nil)
			

			authPath := "/auth"

			require.NoError(t, err)

			server.router.GET(
				authPath,
				authMiddleWare(server.tokenMaker),
				func(ctx *gin.Context) {
					ctx.JSON(http.StatusOK, gin.H{})
				},
			)
			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, authPath, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			



		})
	}
}
