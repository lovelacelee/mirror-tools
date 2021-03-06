/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-16 14:06:05
 * @LastEditTime    : 2022-06-17 18:27:25
 * @LastEditors     : Lovelace
 * @Description     : git emoji
 * @FilePath        : /internal/emoji/emoji.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 * Only some emojis are supported. For more info: https://gitmoji.dev/
 */
package emoji

import (
	"strings"
)

type Emoji struct {
	// Emoji is the description of the emoji.
	Emoji string
	// Code is the code of the emoji.
	Code string
	// Name is the name of the emoji.
	Name []string
} // end of type Emoji

var Emojis = []Emoji{
	{Emoji: "๐ฅ", Code: ":boom:", Name: []string{"default", "Introduce", "changes"}},
	{Emoji: "๐", Code: ":tada:", Name: []string{"้ฆๆฌกๆไบค", "first", "commit"}},
	{Emoji: "๐", Code: ":new:", Name: []string{"ๆฐๅ่ฝ", "new"}},
	{Emoji: "๐", Code: ":bookmark:", Name: []string{"็ๆฌ", "bookmark", "release"}},
	{Emoji: "๐", Code: ":bug:", Name: []string{"้ฎ้ข", "bug"}},
	{Emoji: "๐ง", Code: ":construction:", Name: []string{"ๆๅปบ", "build"}},
	{Emoji: "โ", Code: ":white_check_mark:", Name: []string{"ๆฃๆฅ", "check", "test", "pass"}},
	{Emoji: "๐", Code: ":ambulance:", Name: []string{"็ดงๆฅ", "emergency"}},
	{Emoji: "โฌ๏ธ", Code: ":arrow_down:", Name: []string{"ๅ้", "downgrade"}},
	{Emoji: "โฌ๏ธ", Code: ":arrow_up:", Name: []string{"ๅ็บง", "upgrade"}},
	{Emoji: "๐", Code: ":globe_with_meridians:", Name: []string{"่ฏญ่จ", "language"}},
	{Emoji: "๐", Code: ":lipstick:", Name: []string{"็้ข", "ui"}},
	{Emoji: "๐ฌ", Code: ":clapper:", Name: []string{"ๆผ็คบ", "็คบไพ", "example"}},
	{Emoji: "๐จ", Code: ":rotating_light:", Name: []string{"่ญฆๅ", "warning"}},
	{Emoji: "๐ง", Code: ":wrench:", Name: []string{"้็ฝฎ", "settings", "cfg", "config"}},
	{Emoji: "โ", Code: ":heavy_plus_sign:", Name: []string{"ๆฐๅข", "add"}},
	{Emoji: "โ", Code: ":heavy_minus_sign:", Name: []string{"็งป้ค", "remove"}},
	{Emoji: "โก๏ธ", Code: ":zap:", Name: []string{"ๆ็", "performance"}},
	{Emoji: "๐", Code: ":chart_with_upwards_trend:", Name: []string{"ๅๆ", "analysis"}},
	{Emoji: "๐", Code: ":rocket:", Name: []string{"้ๅบฆ", "ๆง่ฝ", "ๆๅ", "speed"}},
	{Emoji: "๐", Code: ":memo:", Name: []string{"ๆๆกฃ", "doc", "document", "documents", "help"}},
	{Emoji: "๐จ", Code: ":hammer:", Name: []string{"้ๆ", "rebuild"}},
	{Emoji: "๐จ", Code: ":art:", Name: []string{"ๆ?ผๅผๅ", "format"}},
	{Emoji: "โ๏ธ", Code: ":pencil2:", Name: []string{"ไฟฎๅค", "fix", "repair"}},
	{Emoji: "๐", Code: ":lock:", Name: []string{"ๅฎๅจ", "security", "safety"}},
	{Emoji: "๐", Code: ":checkered_flag:", Name: []string{"ๅๅธ", "flag"}},
	{Emoji: "๐ง", Code: ":penguin:", Name: []string{"linux"}},
	{Emoji: "๐ป", Code: ":computer:", Name: []string{"windows"}},
	{Emoji: "๐", Code: ":apple:", Name: []string{"mac"}},
	{Emoji: "๐ณ", Code: ":whale:", Name: []string{"docker"}},
	{Emoji: "๐ฅ", Code: ":fire:", Name: []string{"hot"}},
	{Emoji: "โจ", Code: ":sparkles:", Name: []string{"ๆฐ็นๆง", "feature"}},
	{Emoji: "๐", Code: ":closed_lock_with_key:", Name: []string{"secrets"}},
	{Emoji: "๐จ", Code: ":rotating_light:", Name: []string{"compiler", "compiling"}},
	{Emoji: "๐", Code: ":green_heart:", Name: []string{"ci", "continuous", "integration"}},
	{Emoji: "๐", Code: ":pushpin:", Name: []string{"specific", "version"}},
	{Emoji: "โป๏ธ", Code: ":recycle:", Name: []string{"refactor", "recycle"}},
	{Emoji: "๐ฉ", Code: ":poop:", Name: []string{"damn", "bad"}},
	{Emoji: "โช๏ธ", Code: ":rewind:", Name: []string{"revert", "restore"}},
	{Emoji: "๐", Code: ":twisted_rightwards_arrows:", Name: []string{"merge", "branch"}},
	{Emoji: "๐ฆ๏ธ", Code: ":package:", Name: []string{"pack", "dist"}},
	{Emoji: "๐", Code: ":page_facing_up:", Name: []string{"license"}},
	{Emoji: "๐ก", Code: ":bulb:", Name: []string{"comment", "note"}},
	{Emoji: "๐ฌ", Code: ":speech_balloon:", Name: []string{"chat"}},
	{Emoji: "๐ฉ", Code: ":triangular_flag_on_post:", Name: []string{"flags", "้ๅคง", "important"}},
	{Emoji: "๐งโ๐ป", Code: ":technologist:", Name: []string{"tech", "improve", "experience"}},
	{Emoji: "๐ท๏ธ", Code: "Add or update types.", Name: []string{"tag"}},
} // end of var Emojis

func (e *Emoji) String() string {
	return e.Code
}

func (e *Emoji) GetName() []string {
	return e.Name
}

func Message(message string) string {
	result := Emojis[0].Code + message
	for _, emoji := range Emojis {
		for _, name := range emoji.Name {
			if strings.Contains(message, name) {
				result = emoji.Code + message
				return result
			}
			if strings.Contains(strings.ToLower(message), name) {
				result = emoji.Code + message
				return result
			}
		}
	}
	return result
} // end of func Message
