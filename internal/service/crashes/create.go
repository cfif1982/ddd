package crashes

import (
	"ddd/internal/domain/crashes"

	"github.com/google/uuid"
)

func (crs *CrashService) Create(contractID uuid.UUID, description string) (*crashes.Crash, error) {

	// создаем объект ДТП
	crash, err := crashes.CreateCrash(contractID, description)
	if err != nil {
		return nil, err
	}

	// сохраняем ДТП
	if err = crs.crashRepo.SaveCrash(*crash); err != nil {
		return nil, err
	}

	// нужно изменить скидку клиента
	// для этого достаточно просто в БД изменить поле discount
	// для этого по ContractID находим конракт, у него находим ClientID и у него меняем поле discount
	// как это лучше сделать?
	// есть 2 варианта:
	// 1. Создать несколько функций запросов к БД:
	//	- запросить из репозитория id клиента по contractID
	//	- запросить из репозитория поле discount по ClientID
	//	- сохранить через репзиторий измененное поле discount
	// 2. Создать нужные объекты и через них изменить это поле
	//	- создаем объект contract через метод репозитория GetContract()
	//	- создаем объект клиент через метод репозитория GetClient()
	//	- узнаем поде discount через client.Discount()
	//	- изменяем поле discount у клиента на новое значение
	//	-
	// находим контракт
	contract, err := crs.contractRepo.GetContract(contractID)
	if err != nil {
		return nil, err
	}

	// находим клиента
	client, err := crs.clientRepo.GetClient(contract.ClientID())

	// узнаем действующую скидку клиента
	discount := client.Discount()

	// рассчитываем новую скидку клиенту
	newDiscount := calculateDiscount(discount)

	// сораняем скидку

	// тут чтобы сохранить в БД нужно передавать всю структуру client и вызвать метод SaveClient() репозитория?
	// или нужно сделать в репозитории отдельный метод SaveDiscount()?
	// сохраняем скидку клиента в структуре Client
	client.SetDiscount(newDiscount)
	// cs.clientRepo.SaveClient(client) - так правильно?
	// cs.clientRepo.SaveClientBonus(clientId, newDiscount) - или так правильно?

	return crash, nil
}

func calculateDiscount(d int) int {

	// рассчитываем скидук клиенту
	newDiscount := d

	return newDiscount
}
