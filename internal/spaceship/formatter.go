package spaceship

func formatCreateResponse(res CreateResponseModel) map[string]interface{} {
	return map[string]interface{}{
		"success": res.Success,
	}
}

func formatGetByIDResponse(res GetByIDResponseModel) map[string]interface{} {
	armaments := make([]armamentResponse, len(res.SpaceShip.Armaments))
	for i, armament := range res.SpaceShip.Armaments {
		armaments[i] = armamentResponse{
			Title: armament.Title,
			Qty:   armament.Qty,
		}
	}

	return map[string]interface{}{
		"id":       res.SpaceShip.ID,
		"name":     res.SpaceShip.Name,
		"class":    res.SpaceShip.Class,
		"crew":     res.SpaceShip.Crew,
		"image":    res.SpaceShip.Value,
		"status":   res.SpaceShip.Status,
		"armament": armaments,
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
	spaceships := make([]spaceShipResponse, len(res.SpaceShip))
	for i, spaceship := range res.SpaceShip {
		spaceships[i] = spaceShipResponse{
			ID:     int64(spaceship.ID),
			Name:   spaceship.Name,
			Status: spaceship.Status,
		}
	}

	return map[string]interface{}{
		"data": spaceships,
	}
}
