package main

import "net/http"

func addRoutes(
	mux *http.ServeMux,
	logger *logging.Logger,
	config Config,
	tenantsStore *TenantsStore,
	commentsStore *CommentsStore,
	conversationService *ConversationService,
	chatGPTService *ChatGPTService,
	authProxy *authProxy,
) {
	mux.Handle("/api/v1/", handleTenantsGet(logger, tenantsStore))
	mux.Handle("/oauth2/", handleOAuth2Proxy(logger, authProxy))
	mux.HandleFunc("/healthz", handleHealthzPlease(logger))
	mux.Handle("/", http.NotFoundHandler())
}
