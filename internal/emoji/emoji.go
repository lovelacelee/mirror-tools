/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-16 14:06:05
 * @LastEditTime    : 2022-06-17 18:09:22
 * @LastEditors     : Lovelace
 * @Description     : git emoji
 * @FilePath        : /internal/emoji/emoji.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 * Only some emojis are supported. https://gitmoji.dev/
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
	{Emoji: "💥", Code: ":boom:", Name: []string{"default", "Introduce", "changes"}},
	{Emoji: "🎉", Code: ":tada:", Name: []string{"首次提交", "first", "commit"}},
	{Emoji: "🆕", Code: ":new:", Name: []string{"新功能", "new"}},
	{Emoji: "🔖", Code: ":bookmark:", Name: []string{"版本", "bookmark", "release"}},
	{Emoji: "🐛", Code: ":bug:", Name: []string{"问题", "bug"}},
	{Emoji: "🚧", Code: ":build:", Name: []string{"构建", "build"}},
	{Emoji: "✅", Code: ":check:", Name: []string{"检查", "check", "test", "pass"}},
	{Emoji: "🚑", Code: ":ambulance:", Name: []string{"紧急", "emergency"}},
	{Emoji: "⬇️", Code: ":arrow_down:", Name: []string{"回退", "downgrade"}},
	{Emoji: "⬆️", Code: ":arrow_up:", Name: []string{"升级", "upgrade"}},
	{Emoji: "🌐", Code: ":globe_with_meridians:", Name: []string{"语言", "language"}},
	{Emoji: "💄", Code: ":lipstick:", Name: []string{"界面", "ui"}},
	{Emoji: "🎬", Code: ":clapper:", Name: []string{"演示", "示例", "example"}},
	{Emoji: "🚨", Code: ":rotating_light:", Name: []string{"警告", "warning"}},
	{Emoji: "🔧", Code: ":wrench:", Name: []string{"配置", "settings", "cfg", "config"}},
	{Emoji: "➕", Code: ":heavy_plus_sign:", Name: []string{"新增", "add"}},
	{Emoji: "➖", Code: ":heavy_minus_sign:", Name: []string{"移除", "remove"}},
	{Emoji: "❌", Code: ":circle_with_x:", Name: []string{"无效", "na"}},
	{Emoji: "⚡️", Code: ":zap:", Name: []string{"效率", "performance"}},
	{Emoji: "📈", Code: ":chart_with_upwards_trend:", Name: []string{"分析", "analysis"}},
	{Emoji: "🚀", Code: ":rocket:", Name: []string{"速度", "性能", "提升", "speed"}},
	{Emoji: "📝", Code: ":memo:", Name: []string{"文档", "doc", "document", "documents", "help"}},
	{Emoji: "🔨", Code: ":hammer:", Name: []string{"重构", "rebuild"}},
	{Emoji: "🎨", Code: ":art:", Name: []string{"格式化", "format"}},
	{Emoji: "✏️", Code: ":pencil2:", Name: []string{"修复", "fix", "repair"}},
	{Emoji: "🔒", Code: ":lock:", Name: []string{"安全", "security", "safety"}},
	{Emoji: "🏁", Code: ":checkered_flag:", Name: []string{"发布", "flag"}},
	{Emoji: "🐧", Code: ":penguin:", Name: []string{"linux"}},
	{Emoji: "💻", Code: ":computer:", Name: []string{"windows"}},
	{Emoji: "🍎", Code: ":apple:", Name: []string{"mac"}},
	{Emoji: "🐳", Code: ":whale:", Name: []string{"docker"}},
	{Emoji: "🔥", Code: ":fire:", Name: []string{"hot"}},
	{Emoji: "✨", Code: ":sparkles:", Name: []string{"新特性", "feature"}},
	{Emoji: "🔐", Code: ":closed_lock_with_key:", Name: []string{"secrets"}},
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
