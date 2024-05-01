package contracts

import (
	"ddd/internal/domain/cars"
	"ddd/internal/domain/contracts"

	"github.com/google/uuid"
)

func (cs *ContractService) Create(
	managerId uuid.UUID,
	clientId uuid.UUID,
	carId uuid.UUID,
) (*contracts.Contract, error) {

	manager, err := cs.managerRepo.GetManager(managerId)
	if err != nil {
		return nil, err
	}

	client, err := cs.clientRepo.GetClient(clientId)
	if err != nil {
		return nil, err
	}

	car, err := cs.carRepo.GetCar(carId)
	if err != nil {
		return nil, err
	}

	// рассчитываем бонус менеджера
	managerBonus := calculateManagerBonus(managerId)

	// узнаем скидку клиента
	clientDiscount := client.Discount()

	// рассчитываем сумму контракта
	contractSumma := calculateContractSumma(car, clientDiscount)

	contract, err := contracts.CreateContract(
		managerId,
		clientId,
		carId,
		contractSumma,
	)

	if err != nil {
		return nil, err
	}

	// Сохраняем кнтракт
	if err = cs.contractRepo.SaveContract(*contract); err != nil {
		return nil, err
	}

	// тут чтобы сохранить в БД нужно передавать всю структуру manager и вызвать метод SaveManager() репозитория?
	// или нужно сделать в репозитории отдельный метод SaveBonus()?
	// сохраняем бонус менеджера в структуре Manager
	manager.SetBonus(managerBonus)
	// cs.managerRepo.SaveManager(manager) - так правильно?
	// cs.managerRepo.SaveManagerBonus(managerId, managerBonus) - или так правильно?

	// формируем текст письма клиенту
	emailText := contract.MakeTextForEmail()

	// отправляем письмо клиенту
	client.SendEmail(emailText)

	return contract, nil
}

func calculateManagerBonus(_ uuid.UUID) int {

	// Рассчитываем бонус менеджера

	return 10
}

func calculateContractSumma(_ *cars.Car, clientDiscount int) int {

	// рассчиываем сумму контракта
	contractSumma := clientDiscount

	return contractSumma
}
