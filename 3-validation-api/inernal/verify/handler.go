package verify

import (
	"3-validation-api/configs"
	"3-validation-api/pkg/req"
	"3-validation-api/pkg/resp"
	"fmt"
	"net/http"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type VerifyHandlerDeps struct {
	Config           *configs.Config
	VerifyRepository *VerifyRepository
}

type VerifyHandler struct {
	Config           *configs.Config
	VerifyRepository *VerifyRepository
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{
		Config:           deps.Config,
		VerifyRepository: deps.VerifyRepository,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[VerifyCreateRequest](&w, r)
		if err != nil {
			resp.SetJson(w, err, http.StatusBadRequest)
			return
		}
		emailHash := NewEmailHash(body.Email)
		for {
			existedEmailHash, _ := handler.VerifyRepository.GetByHash(emailHash.Hash)
			if existedEmailHash == nil {
				break
			}
			emailHash.GenerateHash()
		}
		_, err = handler.VerifyRepository.Create(emailHash)
		if err != nil {
			resp.SetJson(w, err.Error(), http.StatusBadGateway)
			return
		}
		href := fmt.Sprintf("http://localhost:81/verify/%s", emailHash.Hash)

		e := email.NewEmail()
		e.From = "Yu Ya <yuriy.505@yandex.ru>"
		e.To = []string{"yuriy.505@yandex.ru"}
		// e.Bcc = []string{"yuriy.505_bcc@yandex.ru"}
		// e.Cc = []string{"yuriy.505_cc@yandex.ru"}
		e.Subject = "Go sended"
		// e.Text = []byte("Text Body is, of course, supported!")
		e.HTML = []byte(fmt.Sprintf("<a href=\"%s\">%s</a>", href, href))
		err = e.Send("smtp.yandex.ru:587", smtp.PlainAuth("", "yuriy.505@yandex.ru", handler.Config.Mail.Password, "smtp.yandex.ru"))
		if err != nil {
			resp.SetJson(w, err.Error(), http.StatusCreated)
		}
		resp.SetJson(w, emailHash, http.StatusCreated)
	}
}

func (handler VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		emailHash, err := handler.VerifyRepository.GetByHash(r.PathValue("hash"))
		if err != nil {
			resp.SetJson(w, err, http.StatusNotFound)
			return
		}
		resp.SetJson(w, emailHash, http.StatusCreated)
	}
}
