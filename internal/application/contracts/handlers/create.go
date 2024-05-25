package handlers

import (
	dContracts "ddd/internal/domain/contracts"
	"ddd/internal/service/contracts"
	"encoding/json"
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
	carItems := req.URL.Query().Get("cars") // получили json car

	var cars []dContracts.Car

	if err := json.Unmarshal([]byte(carItems), &cars); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	// создаем контракт
	contract, err := cs.Create(
		managerId,
		clientId,
		cars,
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
