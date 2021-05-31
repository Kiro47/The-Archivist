package models

import "time"

type GuildMember struct {
	// https://discord.com/developers/docs/resources/guild#guild-member-object
	// UserID Gorm Reference ID
	UserID int `gorm:"not null"`
	// User Foreign key reference to User, if it is null the message is from System
	User User `gorm:"not null"`
	// Nick nickname of the user
	Nick string
	// NickChanged 	true if Nick has been changed before since Created
	NickChanged bool `gorm:"not null"`
	// Roles array of snowflakes
	// TODO: we're going to have to figure out some magic here.
	// TODO: Int's with delimiter? Guild role tables with references? both?
	Roles string
	JoinedAt time.Time `gorm:"not null"`
	// PremiumSince When the user began to boost the server
	//TODO: may try keeping track of in boost events instead of here,
	//TODO: this is finicky due to lapsing
	PremiumSince time.Time
	// Deaf is the user deafened in voice channels
	Deaf bool `gorm:"not null"`
	// Mute is the user muted in voice channels
	Mute bool `gorm:"not null"`
	// Pending If the user has passed screening
	Pending bool `gorm:"not null"`

	// GuildID Associated Guild reference
	GuildID int `gorm:"not null"`
	// Guild Associated Guild
	Guild Guild `gorm:"not null"`
	// Created Timestamp of when this entry was created
	Created int64 `gorm:"autoCreateTime"`
/*
	Some fields are always using default values when accessed as a bot, this is due
	to them not being supported outside of OAuth2 access.
		permissions
 */
}

// TODO: NickUpdates, DeafUpdates,MuteUpdates,
// TODO: JoinedAtUpdates Better off being handled at join/leave time,
// TODO: but we can keep this update to track "current length"
// TODO: RoleUpdates may be better from the Audit log?
// TODO: PremiumSinceUpdates, better in boost events, should we even bother here?