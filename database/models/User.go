package models

// User Table container Discord users (does not store non-accessible bot fields)
type User struct {
	// https://discord.com/developers/docs/resources/user#user-object-user-structure
	ID int64 `gorm:"primaryKey;unique;not null"`// SnowflakeID https://en.wikipedia.org/wiki/Snowflake_ID
	// Username Name of the user
	Username string `gorm:"not null"`
	// UsernameChanged true if username has been changed before since Created
	UsernameChanged bool `gorm:"not null"`
	Discriminator string `gorm:"not null"`// Technically a string according to docs
	// Avatar Hash of avatar, can be null if never set
	Avatar string
	// AvatarChanged true if avatar has been changed before since Created
	AvatarChanged  bool `gorm:"not null"`
	// Bot is bot?
	Bot bool `gorm:"not null"`
	// Flags I think this is bot accessible?  I don't have anyone to test with to confirm
	// https://discord.com/developers/docs/resources/user#user-object-user-flags
	Flags int `gorm:"not null"`
	// Created Timestamp of when this entry was created
	Created int64 `gorm:"autoCreateTime"`
	/*
	Some fields are always using default values when accessed as a bot, this is due
	to them not being supported outside of OAuth2 access.
	         "email" : "",
	         "flags" : 0,
	         "locale" : "",
	         "mfa_enabled" : false,
	         "premium_type" : 0,
	         "token" : "",
	         "verified" : false
	*/
}

// UsernameUpdate Table containing username updates
type UsernameUpdate struct {
	ID uint64 `gorm:"autoIncrement;unique"`
	UserID int // GORM Reference ID
	User User `gorm:"not null"`
	Username string `gorm:"not null"`
	Created int64 `gorm:"autoCreateTime"`
}

// UserAvatarUpdate Table containing user avatar updates
type UserAvatarUpdate struct {
	ID uint64 `gorm:"autoIncrement;unique"`
	UserID int // GORM Reference ID
	User User `gorm:"not null"`
	Avatar string
	Created int64 `gorm:"autoCreateTime"`
}