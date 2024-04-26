package resource

import (
	"errors"
	"github.com/cend-org/duval/graph/model"
	"github.com/cend-org/duval/internal/database"
	"github.com/cend-org/duval/internal/utils/state"
)

const (
	ResTypeMessage = 0
	ResTypeMenu    = 1
)

func GetMessages() (messages []model.Message, err error) {
	query := `SELECT * FROM message WHERE resource_type = ?`
	err = database.Select(&messages, query, ResTypeMessage)
	if err != nil {
		return nil, err
	}
	return messages, err
}

func GetMessagesInLanguage(language int) (messages []model.Message, err error) {
	query := `
				SELECT 
				    COALESCE(target.ID, english.ID) as 'id',
				    COALESCE(target.created_at, english.created_at) as 'created_at',
				    COALESCE(target.updated_at, english.updated_at) as 'updated_at',
				    COALESCE(target.deleted_at, english.deleted_at) as 'deleted_at',
				    COALESCE(target.resource_type, english.resource_type) as 'resource_type',
				    COALESCE(target.resource_number, english.resource_number) as 'resource_number',
				    COALESCE(target.resource_value, english.resource_value) as 'resource_value',
				    COALESCE(target.resource_label, english.resource_label) as 'resource_label',
				    COALESCE(target.resource_language, english.resource_language) as 'resource_language'
				    
					FROM message english
					LEFT JOIN message target ON english.resource_type = target.resource_type
					                                AND english.resource_number = target.resource_number
					                                AND english.resource_value = target.resource_value
					WHERE english.resource_type = ?
					  AND english.resource_language = 0
					  AND target.resource_language = ?
			`
	err = database.Select(&messages, query, ResTypeMessage, language)
	if err != nil {
		return nil, err
	}

	return messages, err
}

func GetMessage(resourceNumber, resourceLanguage int) (message model.Message, err error) {
	query := `
				SELECT 
				    COALESCE(target.ID, english.ID) as 'id',
				    COALESCE(target.created_at, english.created_at) as 'created_at',
				    COALESCE(target.updated_at, english.updated_at) as 'updated_at',
				    COALESCE(target.deleted_at, english.deleted_at) as 'deleted_at',
				    COALESCE(target.resource_type, english.resource_type) as 'resource_type',
				    COALESCE(target.resource_number, english.resource_number) as 'resource_number',
				    COALESCE(target.resource_value, english.resource_value) as 'resource_value',
				    COALESCE(target.resource_label, english.resource_label) as 'resource_label',
				    COALESCE(target.resource_language, english.resource_language) as 'resource_language'
				
					FROM message english
					LEFT JOIN message target ON english.resource_type = target.resource_type
					                                AND english.resource_number = target.resource_number
					                                AND english.resource_value = target.resource_value
													AND target.resource_language = ?
					WHERE english.resource_type = ?
					  AND english.resource_number = ?
					  AND english.resource_value = 0
					  AND english.resource_language = 0
					  
			`
	err = database.Get(&message, query, resourceLanguage, ResTypeMessage, resourceNumber)
	if err != nil {
		return message, err
	}

	return message, err
}

func NewMessage(resourceLabel string, resourceLanguage int) (message model.Message, err error) {
	message.ResourceType = ResTypeMessage
	message.ResourceNumber = getNewMessageNumber()
	message.ResourceValue = 0
	message.ResourceLanguage = resourceLanguage
	message.ResourceLabel = resourceLabel

	message.ID, err = database.InsertOne(message)
	if err != nil {
		return message, err
	}

	if message.ResourceLanguage > 0 {
		englishLanguageMessage := message
		englishLanguageMessage.ResourceLanguage = 0
		_, err = database.InsertOne(englishLanguageMessage)
		if err != nil {
			return message, err
		}
	}

	return message, err
}

func DeleteMessage(resourceNumber, resourceLanguage int) (err error) {
	var message model.Message

	message, err = getMessage(ResTypeMessage, resourceNumber, state.ZERO, resourceLanguage)
	if err != nil {
		return err
	}

	if message.ResourceLanguage == 0 {
		err = database.Exec(`DELETE FROM message WHERE resource_type = ? AND resource_number = ?`, ResTypeMessage, resourceNumber)
		if err != nil {
			return err
		}

		return err
	}

	err = database.Delete(message)
	if err != nil {
		return err
	}

	return err
}

func UpdateMessage(message model.Message) (msg model.Message, err error) {
	msg, err = getMessage(ResTypeMessage, message.ResourceNumber, message.ResourceValue, message.ResourceLanguage)
	if err != nil {
		return message, err
	}

	message.ID = msg.ID
	err = database.Update(message)
	if err != nil {
		return message, err
	}

	return message, err
}

func GetMenuList() (menus []model.Message, err error) {
	query := `SELECT * FROM message WHERE resource_type = ? AND resource_number = 0  AND resource_language = ? ORDER BY resource_number desc`
	err = database.Select(&menus, query, ResTypeMenu, 0)
	if err != nil {
		return nil, err
	}
	return menus, err
}

func GetMenuItems(menuNumber, menuLanguage int) (menus []model.Message, err error) {
	if menuNumber == 0 {
		return nil, errors.New("cannot get menu list")
	}

	query := `
				SELECT 
				    COALESCE(target.ID, english.ID) as 'id',
				    COALESCE(target.created_at, english.created_at) as 'created_at',
				    COALESCE(target.updated_at, english.updated_at) as 'updated_at',
				    COALESCE(target.deleted_at, english.deleted_at) as 'deleted_at',
				    COALESCE(target.resource_type, english.resource_type) as 'resource_type',
				    COALESCE(target.resource_number, english.resource_number) as 'resource_number',
				    COALESCE(target.resource_value, english.resource_value) as 'resource_value',
				    COALESCE(target.resource_label, english.resource_label) as 'resource_label',
				    COALESCE(target.resource_language, english.resource_language) as 'resource_language'
				
					FROM message english
					LEFT JOIN message target ON english.resource_type = target.resource_type
					                                AND english.resource_number = target.resource_number
					                                AND english.resource_value = target.resource_value
													AND target.resource_language = ?
					WHERE english.resource_type = ?
					  AND english.resource_number = ?
					  AND english.resource_language = 0
					ORDER BY resource_type, resource_number,resource_value
			`

	err = database.Select(&menus, query, menuLanguage, ResTypeMenu, menuNumber)
	if err != nil {
		return nil, err
	}

	return menus, err
}

func NewMenu(menuName string, resourceLanguage int) (menu model.Message, err error) {
	menu.ResourceType = ResTypeMenu
	menu.ResourceNumber = 0
	menu.ResourceLabel = menuName
	menu.ResourceValue = getNewMenuNumber()

	if resourceLanguage > 0 {
		menu.ResourceLanguage = 0
		_, err = database.InsertOne(menu)
		if err != nil {
			return menu, err
		}
	}

	menu.ResourceLanguage = resourceLanguage

	menu.ID, err = database.InsertOne(menu)
	if err != nil {
		return menu, err
	}

	return menu, err
}

func DeleteMenu(menuNumber int) (err error) {
	var menu model.Message
	err = database.Get(&menu, `SELECT * FROM message WHERE resource_type = ? AND resource_number = 0 AND resource_value = ?`, ResTypeMenu, menuNumber)
	if err != nil {
		return err
	}

	err = database.Exec(`DELETE FROM message WHERE resource_type = ? AND resource_number = ?`, ResTypeMessage, menuNumber)
	if err != nil {
		return err
	}

	err = database.Delete(menu)
	if err != nil {
		return err
	}

	return err
}

func NewMenuItem(menuLabel string, menuNumber, resourceLanguage int) (menu model.Message, err error) {
	menu.ResourceType = ResTypeMenu
	menu.ResourceNumber = menuNumber
	menu.ResourceLabel = menuLabel
	menu.ResourceValue = getNewMenuItemNumber(menuNumber)

	if resourceLanguage > 0 {
		menu.ResourceLanguage = 0
		_, err = database.InsertOne(menu)
		if err != nil {
			return menu, err
		}
	}

	menu.ResourceLanguage = resourceLanguage

	menu.ID, err = database.InsertOne(menu)
	if err != nil {
		return menu, err
	}

	return menu, err
}

func DeleteMenuItem(menuNumber, menuValue, resourceLanguage int) (err error) {
	var menu model.Message
	err = database.Get(&menu, `SELECT * FROM message WHERE resource_type = ? AND resource_number = ? AND resource_value = ? AND resource_language = ?`, ResTypeMenu, menuNumber, menuValue, resourceLanguage)
	if err != nil {
		return err
	}

	if resourceLanguage == 0 {
		err = database.Exec(`DELETE FROM message WHERE resource_type = ? AND resource_number = ? AND resource_value = ?`, ResTypeMenu, menuNumber, menuValue)
		if err != nil {
			return err
		}
		return err
	}

	err = database.Delete(menu)
	if err != nil {
		return err
	}

	return err
}

func getNewMessageNumber() (number int) {
	err := database.Get(&number, `SELECT resource_number FROM message WHERE resource_type = ? ORDER BY resource_number DESC LIMIT 1`, ResTypeMessage)
	if err != nil {
		return 0
	}
	return number + 1
}

func getNewMenuNumber() (number int) {
	err := database.Get(&number, `SELECT resource_value FROM message WHERE resource_type = ? AND resource_number = 0 ORDER BY resource_value DESC LIMIT 1`, ResTypeMenu)
	if err != nil {
		return 0
	}
	return number + 1
}

func getNewMenuItemNumber(menuNumber int) (number int) {
	err := database.Get(&number, `SELECT resource_value FROM message WHERE resource_type = ? AND resource_number = ? ORDER BY resource_value DESC LIMIT 1`, ResTypeMenu, menuNumber)
	if err != nil {
		return 0
	}
	return number + 1
}

func getMessage(resourceType, resourceNumber, resourceValue, resourceLanguage int) (message model.Message, err error) {
	query := `SELECT * FROM message WHERE resource_type = ? AND resource_number = ? AND resource_value = ? AND resource_language = ?`
	err = database.Get(&message, query, resourceType, resourceNumber, resourceValue, resourceLanguage)
	if err != nil {
		return message, err
	}

	return message, err
}
