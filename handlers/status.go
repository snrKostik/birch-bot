package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func ChangeStatus(dg *discordgo.Session, appID string) {

	dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		err := dg.UpdateStatusComplex(discordgo.UpdateStatusData{
			Activities: []*discordgo.Activity{
				{
					Name:          "Зона Березовых Домов",
					Type:          discordgo.ActivityTypeGame,
					ApplicationID: appID,
					Assets: discordgo.Assets{
						LargeImageID: "zbd",
						LargeText:    "Зона Березовых Домов",
					},
				},
			},
			Status: "dnd",
		})
		if err != nil {
			fmt.Println("cant change status:", err)
			return
		} else {
			fmt.Println("now bot have cool status")
		}
	})
}
