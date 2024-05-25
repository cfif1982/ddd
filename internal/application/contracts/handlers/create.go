package handlers

import (
	"ddd/internal/service/contracts"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (h *Handler) CreateContract(rw http.ResponseWriter, req *http.Request) {

	// создаем сервис для работы с контрактом
	cs := contracts.NewContractService(
		h.contractRepo,
		h.managerRepo,
		h.clientRepo,
	)

	// узнаем переменные из запроса
	managerId, _ := uuid.Parse(req.URL.Query().Get("manager_id"))
	clientId, _ := uuid.Parse(req.URL.Query().Get("client_id"))

	// создаем контракт
	contract, err := cs.Create(
		managerId,
		clientId,
	)

	if err != nil {
		log.Fatal("internal server error")
	}

	// выводим ответ сервера
	// Устанавливаем в заголовке тип передаваемых данных
	rw.Header().Set("Content-Type", "text/plain")

	// устанавливаем код 201
	rw.WriteHeader(http.StatusCreated)

	// формируем текст ответа сервера
	answerText := "contract created! ID:" + contract.ID().String()

	// выводим ответ сервера
	_, err = rw.Write([]byte(answerText))
	if err != nil {
		log.Fatal(err.Error())
	}

}
