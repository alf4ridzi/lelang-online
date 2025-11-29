package config

import "lelang-online-api/models"

var ModelMigration = []any{
	models.Role{},
	models.User{},
	models.Item{},
	models.Auction{},
	models.AuctionHistory{},
}
