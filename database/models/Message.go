package models

/*
TODO: This will work for testing, but messages will need to be moved to a
TODO: dynamically generated model based on guild ID.  When you begin to scale
TODO: to many servers, one supermassive message table is just going to be
TODO: falling apart from a performance perspective.
 */

type Message struct {
	// https://discord.com/developers/docs/resources/channel#message-object
	/*
	For now we use a triple composite key which Discord also uses as refs
	once we moved to a message table per guild ChannelID and GuildID will just
	become foreign key references.
	 */
	// MessageID message Snowflake  https://en.wikipedia.org/wiki/Snowflake_ID
	MessageID int64 `gorm:"primaryKey;not null"`
	// ChannelID channel snowflake https://en.wikipedia.org/wiki/Snowflake_ID
	ChannelID int64 `gorm:"primaryKey;not null"`
	// GuildID guild snowflake https://en.wikipedia.org/wiki/Snowflake_ID
	GuildID int64 `gorm:"primaryKey;not null"`
	// AuthorID Gorm Reference ID
	AuthorID int `gorm:"not null"`
	// Author Foreign key reference to User, if it is null the message is from System
	Author User `gorm:"not null"`
	//TODO: Finish
}
