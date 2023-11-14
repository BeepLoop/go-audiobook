package store

import "log"

func InitStorage() error {
	err := InitDirectories()
	if err != nil {
		log.Println(err)
		return err
	}

	err = InitStoryStore()
	if err != nil {
		log.Println(err)
		return err
	}

	err = InitAdminStore()
	if err != nil {
		log.Println(err)
		return err
	}

	err = LoadCredential()
	if err != nil {
		log.Println(err)
		return err
	}

	err = LoadStories()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
