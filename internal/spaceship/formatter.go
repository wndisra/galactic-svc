package spaceship

func formatCreateResponse(res CreateResponseModel) map[string]interface{} {
	return map[string]interface{}{
		"success": res.Success,
	}
}

func formatGetByIDResponse(res GetByIDResponseModel) map[string]interface{} {
	return map[string]interface{}{
		"id":       res.SpaceShip.ID,
		"name":     res.SpaceShip.Name,
		"class":    res.SpaceShip.Class,
		"crew":     res.SpaceShip.Crew,
		"image":    res.SpaceShip.Value,
		"status":   res.SpaceShip.Status,
		"armament": res.SpaceShip.Armaments,
	}
}

func formatUpdateResponse(res UpdateResponseModel) map[string]interface{} {
	return map[string]interface{}{
		"success": res.Success,
	}
}

func formatDeleteByIDResponse(res DeleteByIDResponseModel) map[string]interface{} {
	return map[string]interface{}{
		"success": res.Success,
	}
}

func formatGetAllResponse(res GetAllResponseModel) map[string]interface{} {
	return map[string]interface{}{
		"data": map[string]interface{}{},
	}
}
