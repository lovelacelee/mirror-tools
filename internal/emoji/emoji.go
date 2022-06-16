/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-16 14:06:05
 * @LastEditTime    : 2022-06-16 14:43:56
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /internal/emoji.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package emoji

import (
	"strings"
)

type Emoji struct {
	// Name is the name of the emoji.
	Name string
	// Code is the code of the emoji.
	Code string
	// Emoji is the description of the emoji.
	Emoji string
} // end of type Emoji

var Emojis = []Emoji{
	{Name: "default", Code: ":circle_arrow_right:", Emoji: "➡️"},
	{Name: "first commit", Code: ":tada:", Emoji: "🎉"},
	{Name: "new", Code: ":new:", Emoji: "🆕"},
	{Name: "bookmark", Code: ":bookmark:", Emoji: "🔖"},
	{Name: "bug", Code: ":bug:", Emoji: "🐛"},
	{Name: "build", Code: ":build:", Emoji: "🚧"},
	{Name: "check", Code: ":check:", Emoji: "✅"},
	{Name: "Emergency repair", Code: ":ambulance:", Emoji: "🚑"},
	{Name: "downgrade", Code: ":circle_arrow_down:", Emoji: "⬇️"},
	{Name: "upgrade", Code: ":circle_arrow_up:", Emoji: "⬆️"},
	{Name: "language globalize", Code: ":globe_with_meridians:", Emoji: "🌐"},
	{Name: "ui", Code: ":lipstick:", Emoji: "💄"},
	{Name: "example display show", Code: ":clapper:", Emoji: "🎬 "},
	{Name: "warning remove", Code: ":rotating_light:", Emoji: "🚨"},
	{Name: "settings", Code: ":wrench:", Emoji: "🔧"},
	{Name: "Add", Code: ":heavy_plus_sign:", Emoji: "➕"},
	{Name: "Remove", Code: ":heavy_minus_sign:", Emoji: "➖"},
	{Name: "NA", Code: ":circle_with_x:", Emoji: "❌"},
	{Name: "performance", Code: ":zap:", Emoji: "⚡"},
	{Name: "analysis", Code: ":chart_with_upwards_trend:", Emoji: "📈"},
	{Name: "speed up", Code: ":rocket:", Emoji: "🚀"},
	{Name: "doc documents document", Code: ":memo:", Emoji: "📝"},
	{Name: "rebuild", Code: ":hammer:", Emoji: "🔨"},
	{Name: "format", Code: ":art:", Emoji: "🎨"},
	{Name: "repair fix", Code: ":pencil2:", Emoji: "✏️"},
	{Name: "security safty", Code: ":lock:", Emoji: "🔒"},
	{Name: "Checked flag", Code: ":checkered_flag:", Emoji: "🏁"},
	{Name: "linux", Code: ":penguin:", Emoji: "🐧"},
	{Name: "windows", Code: ":computer:", Emoji: "💻"},
	{Name: "mac", Code: ":apple:", Emoji: "🍎"},
	{Name: "docker", Code: ":whale:", Emoji: "🐳"},
} // end of var Emojis

func (e *Emoji) String() string {
	return e.Code
}

func (e *Emoji) GetName() string {
	return e.Name
}

func Message(message string) string {
	result := Emojis[0].Code + message
	for _, emoji := range Emojis {
		if strings.Contains(message, emoji.Name) {
			result = emoji.Code + message
			break
		}
	}
	return result
} // end of func Message
