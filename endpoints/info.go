package endpoints

import (
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
	"github.com/tomochain/tomox-sdk/app"
	"github.com/tomochain/tomox-sdk/interfaces"
	"github.com/tomochain/tomox-sdk/utils/httputils"
)

type infoEndpoint struct {
	walletService interfaces.WalletService
	tokenService  interfaces.TokenService
}

func ServeInfoResource(
	r *mux.Router,
	walletService interfaces.WalletService,
	tokenService interfaces.TokenService,
) {

	e := &infoEndpoint{walletService, tokenService}
	r.HandleFunc("/api/info", e.handleGetInfo)
	r.HandleFunc("/api/info/exchange", e.handleGetExchangeInfo)
	r.HandleFunc("/api/info/fees", e.handleGetFeeInfo)
}

func (e *infoEndpoint) handleGetInfo(w http.ResponseWriter, r *http.Request) {
	ex := common.HexToAddress(app.Config.Tomochain["exchange_address"])

	quotes, err := e.tokenService.GetAll()
	if err != nil {
		logger.Error(err)
	}

	var fee string
	if len(quotes) == 0 {
		fee = "0"
	} else {
		f, _, _ := big.ParseFloat(quotes[0].MakeFee.String(), int(10), uint(10), big.ToZero)
		fee = big.NewFloat(0).Mul(f, big.NewFloat(0.0001)).String()
	}

	res := map[string]interface{}{
		"exchangeAddress": ex.Hex(),
		"fee":             fee,
	}

	httputils.WriteJSON(w, http.StatusOK, res)
}

func (e *infoEndpoint) handleGetExchangeInfo(w http.ResponseWriter, r *http.Request) {
	ex := common.HexToAddress(app.Config.Tomochain["exchange_address"])

	res := map[string]string{"exchangeAddress": ex.Hex()}

	httputils.WriteJSON(w, http.StatusOK, res)
}

func (e *infoEndpoint) handleGetFeeInfo(w http.ResponseWriter, r *http.Request) {
	quotes, err := e.tokenService.GetAll()
	if err != nil {
		logger.Error(err)
	}

	var fee string
	if len(quotes) == 0 {
		fee = "0"
	} else {
		f, _, _ := big.ParseFloat(quotes[0].MakeFee.String(), int(10), uint(10), big.ToZero)
		fee = big.NewFloat(0).Mul(f, big.NewFloat(0.0001)).String()
	}

	res := map[string]string{"fee": fee}

	httputils.WriteJSON(w, http.StatusOK, res) // This value will be divided by 1000 on TomoX
}
