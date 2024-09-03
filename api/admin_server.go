// Copyright (c) 2024 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package api

import (
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/luckypickle/vet-thor/co"
	"github.com/pkg/errors"
)

func HTTPHandler(logLevel *slog.LevelVar) http.Handler {
	router := mux.NewRouter()
	sub := router.PathPrefix("/admin").Subrouter()
	sub.Path("/loglevel").
		Methods(http.MethodGet).
		Name("get-log-level").
		HandlerFunc(getLogLevelHandler(logLevel))

	sub.Path("/loglevel").
		Methods(http.MethodPost).
		Name("post-log-level").
		HandlerFunc(postLogLevelHandler(logLevel))

	return handlers.CompressHandler(router)
}

func StartAdminServer(addr string, logLevel *slog.LevelVar) (string, func(), error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return "", nil, errors.Wrapf(err, "listen admin API addr [%v]", addr)
	}

	router := mux.NewRouter()
	router.PathPrefix("/admin").Handler(HTTPHandler(logLevel))
	handler := handlers.CompressHandler(router)

	srv := &http.Server{Handler: handler, ReadHeaderTimeout: time.Second, ReadTimeout: 5 * time.Second}
	var goes co.Goes
	goes.Go(func() {
		srv.Serve(listener)
	})
	return "http://" + listener.Addr().String() + "/admin", func() {
		srv.Close()
		goes.Wait()
	}, nil
}
