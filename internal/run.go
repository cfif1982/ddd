package internal

func Run() error {

	// загружаем конфигурацию

	// настраиваем логгер

	srv := NewServer()

	if err := srv.Run(); err != nil {
		return err
	}

	return nil
}
