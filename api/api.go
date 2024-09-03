// Copyright (c) 2018 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package api

import (
	"net/http"
	"net/http/pprof"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/luckypickle/vet-thor/api/accounts"
	"github.com/luckypickle/vet-thor/api/blocks"
	"github.com/luckypickle/vet-thor/api/debug"
	"github.com/luckypickle/vet-thor/api/doc"
	"github.com/luckypickle/vet-thor/api/events"
	"github.com/luckypickle/vet-thor/api/node"
	"github.com/luckypickle/vet-thor/api/subscriptions"
	"github.com/luckypickle/vet-thor/api/transactions"
	"github.com/luckypickle/vet-thor/api/transfers"
	"github.com/luckypickle/vet-thor/bft"
	"github.com/luckypickle/vet-thor/chain"
	"github.com/luckypickle/vet-thor/log"
	"github.com/luckypickle/vet-thor/logdb"
	"github.com/luckypickle/vet-thor/state"
	"github.com/luckypickle/vet-thor/thor"
	"github.com/luckypickle/vet-thor/txpool"
)

var logger = log.WithContext("pkg", "api")

// New return api router
func New(
	repo *chain.Repository,
	stater *state.Stater,
	txPool *txpool.TxPool,
	logDB *logdb.LogDB,
	bft bft.Finalizer,
	nw node.Network,
	forkConfig thor.ForkConfig,
	allowedOrigins string,
	backtraceLimit uint32,
	callGasLimit uint64,
	pprofOn bool,
	skipLogs bool,
	allowCustomTracer bool,
	enableReqLogger bool,
	enableMetrics bool,
	logsLimit uint64,
	allowedTracers map[string]interface{},
	soloMode bool,
) (http.HandlerFunc, func()) {
	origins := strings.Split(strings.TrimSpace(allowedOrigins), ",")
	for i, o := range origins {
		origins[i] = strings.ToLower(strings.TrimSpace(o))
	}

	router := mux.NewRouter()

	// to serve stoplight, swagger and api docs
	router.PathPrefix("/doc").Handler(
		http.StripPrefix("/doc/", http.FileServer(http.FS(doc.FS))),
	)

	// redirect stoplight-ui
	router.Path("/").HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			http.Redirect(w, req, "doc/stoplight-ui/", http.StatusTemporaryRedirect)
		})

	accounts.New(repo, stater, callGasLimit, forkConfig, bft).
		Mount(router, "/accounts")

	if !skipLogs {
		events.New(repo, logDB, logsLimit).
			Mount(router, "/logs/event")
		transfers.New(repo, logDB, logsLimit).
			Mount(router, "/logs/transfer")
	}
	blocks.New(repo, bft).
		Mount(router, "/blocks")
	transactions.New(repo, txPool).
		Mount(router, "/transactions")
	debug.New(repo, stater, forkConfig, callGasLimit, allowCustomTracer, bft, allowedTracers, soloMode).
		Mount(router, "/debug")
	node.New(nw).
		Mount(router, "/node")
	subs := subscriptions.New(repo, origins, backtraceLimit, txPool)
	subs.Mount(router, "/subscriptions")

	if pprofOn {
		router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		router.HandleFunc("/debug/pprof/profile", pprof.Profile)
		router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		router.HandleFunc("/debug/pprof/trace", pprof.Trace)
		router.PathPrefix("/debug/pprof/").HandlerFunc(pprof.Index)
	}

	if enableMetrics {
		router.Use(metricsMiddleware)
	}

	handler := handlers.CompressHandler(router)
	handler = handlers.CORS(
		handlers.AllowedOrigins(origins),
		handlers.AllowedHeaders([]string{"content-type", "x-genesis-id"}),
		handlers.ExposedHeaders([]string{"x-genesis-id", "x-thorest-ver"}),
	)(handler)

	if enableReqLogger {
		handler = RequestLoggerHandler(handler, logger)
	}

	return handler.ServeHTTP, subs.Close // subscriptions handles hijacked conns, which need to be closed
}
