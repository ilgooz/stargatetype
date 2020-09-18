package rest

import (
	"net/http"

    "github.com/gorilla/mux"

    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/j/j8/x/j8/types"
)

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
    r.HandleFunc("/j8/user", createUserHandler(clientCtx)).Methods("POST")
}

type createUserRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	Name string `json:"name"`
	Email string `json:"email"`
	
}

func createUserHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createUserRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgUser(creator,  req.Name,  req.Email, )
		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
