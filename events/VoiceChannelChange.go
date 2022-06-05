package events

import (
	"context"
	"fmt"
	"time"

	"github.com/Constani/discordbot/reference"
	"github.com/Constani/discordbot/scripts"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var col *mongo.Collection = scripts.GetCollection(scripts.DB, "bot")

func VoiceStateUpdate(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
	result, err := IsData(v.UserID)
	if err != nil {
		fmt.Println(result != nil)
		if v.VoiceState.ChannelID == "979681528252604437" {

			if result == nil {
				user, _ := s.User(v.UserID)
				channel, err := s.GuildChannelCreateComplex(v.GuildID, discordgo.GuildChannelCreateData{
					Name:     "🌒・" + user.Username,
					Type:     discordgo.ChannelTypeGuildVoice,
					ParentID: "979681434711232543",
				})
				if err != nil {
					s.ChannelMessageSend(v.UserID, "Kanal Oluşturulken hata oldu")
				}
				err = s.GuildMemberMove(v.GuildID, v.UserID, &channel.ID)
				if err != nil {
					s.ChannelMessageSend(v.UserID, "Kanal Oluşturulken hata oldu")
				}
				err = CreateData(v.UserID, channel.ID)
				if err != nil {
					s.ChannelMessageSend(v.UserID, "Kanal Oluşturulken hata oldu")
				}
			}
			fmt.Println(result == nil)
			if result != nil && result.ChannelID == "yok" {
				if v.VoiceState.ChannelID == "979681528252604437" {
					user, _ := s.User(v.UserID)
					channel, err := s.GuildChannelCreateComplex(v.GuildID, discordgo.GuildChannelCreateData{
						Name:     "🌒・" + user.Username,
						Type:     discordgo.ChannelTypeGuildVoice,
						ParentID: "979681434711232543",
					})
					if err != nil {
						s.ChannelMessageSend(v.UserID, "Kanal Oluşturulken hata oldu")
					}
					err = s.GuildMemberMove(v.GuildID, v.UserID, &channel.ID)
					if err != nil {
						s.ChannelMessageSend(v.UserID, "Kanal Oluşturulken hata oldu")
					}
					err = CreateData(v.UserID, channel.ID)
					if err != nil {
						s.ChannelMessageSend(v.UserID, "Kanal Oluşturulken hata oldu")
					}
				}
			}
		}
	} else {
		if v.BeforeUpdate.ChannelID == result.ChannelID {
			channel, err := s.Channel(result.ChannelID)
			if err != nil {
				s.ChannelMessageSend(v.UserID, "Kanal Silinirken hata oldu")
			}
			err = updateData(v.UserID)
			if err != nil {
				s.ChannelMessageSend(v.UserID, "Kanal Silinirken hata oldu")
			}
			s.ChannelDelete(channel.ID)
		}
	}
}

func IsData(id string) (*reference.User, error) {
	var result reference.User
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	err := col.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return &reference.User{}, err
	}
	return &result, nil
}
func CreateData(id string, channelID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newData := reference.User{
		Id:        id,
		ChannelID: channelID,
	}
	_, err := col.InsertOne(ctx, newData)
	if err != nil {
		return err
	}
	return nil
}

func updateData(id string) error {
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "channelid", Value: "yok"}}}}
	_, err := col.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
