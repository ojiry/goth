package service

import(
	"errors"
)

type AuthorizeRequest struct {
	Scope        string  `json:"scope"`
	ResponseType string  `json:"response_type"`
	ClientID     string  `json:"client_id"`
	RedirectUri  string  `json:"redirect_uri"`
	State        *string `json:"state"`
	ResponseMode *string `json:"response_mode"`
	Nonce        *string `json:"nonce"`
	Display      *string `json:"display"`
	Prompt       *string `json:"prompt"`
	MaxAge       *string `json:"max_age"`
	UILocales    *string `json:"ui_locales"`
	IDTokenHint  *string `json:"id_token_hint"`
	LoginHint    *string `json:"login_hint"`
	AcrValues    *string `json:"acr_values"`
}

type AuthorizeService struct {
	authorizeRequest AuthorizeRequest
}

func NewAuthorizeService(ar AuthorizeRequest) *AuthorizeService {
	return &AuthorizeService{authorizeRequest: ar}
}

func (s *AuthorizeService) Validate() error {
	if s.authorizeRequest.Scope == "" {
		return errors.New("scope parameter is required")
	}
	return nil
}
