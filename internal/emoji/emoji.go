/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-16 14:06:05
 * @LastEditTime    : 2022-06-17 17:28:38
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
	// Name is the name of the emoji.
	Name []string
	// Code is the code of the emoji.
	Code string
	// Emoji is the description of the emoji.
	Emoji string
} // end of type Emoji

var Emojis = []Emoji{
	{Name: []string{"default", "Introduce", "changes"}, Code: ":boom:", Emoji: "💥"},
	{Name: []string{"首次提交", "first", "commit"}, Code: ":tada:", Emoji: "🎉"},
	{Name: []string{"新功能", "new"}, Code: ":new:", Emoji: "🆕"},
	{Name: []string{"版本", "bookmark", "release"}, Code: ":bookmark:", Emoji: "🔖"},
	{Name: []string{"问题", "bug"}, Code: ":bug:", Emoji: "🐛"},
	{Name: []string{"构建", "build"}, Code: ":build:", Emoji: "🚧"},
	{Name: []string{"检查", "check"}, Code: ":check:", Emoji: "✅"},
	{Name: []string{"紧急", "emergency"}, Code: ":ambulance:", Emoji: "🚑"},
	{Name: []string{"回退", "downgrade"}, Code: ":arrow_down:", Emoji: "⬇️"},
	{Name: []string{"升级", "upgrade"}, Code: ":arrow_up:", Emoji: "⬆️"},
	{Name: []string{"语言", "language"}, Code: ":globe_with_meridians:", Emoji: "🌐"},
	{Name: []string{"界面", "ui"}, Code: ":lipstick:", Emoji: "💄"},
	{Name: []string{"演示", "示例", "example"}, Code: ":clapper:", Emoji: "🎬 "},
	{Name: []string{"警告", "warning"}, Code: ":rotating_light:", Emoji: "🚨"},
	{Name: []string{"配置", "settings"}, Code: ":wrench:", Emoji: "🔧"},
	{Name: []string{"新增", "add"}, Code: ":heavy_plus_sign:", Emoji: "➕"},
	{Name: []string{"移除", "remove"}, Code: ":heavy_minus_sign:", Emoji: "➖"},
	{Name: []string{"无效", "na"}, Code: ":circle_with_x:", Emoji: "❌"},
	{Name: []string{"效率", "performance"}, Code: ":zap:", Emoji: "⚡"},
	{Name: []string{"分析", "analysis"}, Code: ":chart_with_upwards_trend:", Emoji: "📈"},
	{Name: []string{"速度", "性能", "提升", "speed"}, Code: ":rocket:", Emoji: "🚀"},
	{Name: []string{"文档", "doc", "document", "documents", "help"}, Code: ":memo:", Emoji: "📝"},
	{Name: []string{"重构", "rebuild"}, Code: ":hammer:", Emoji: "🔨"},
	{Name: []string{"格式化", "format"}, Code: ":art:", Emoji: "🎨"},
	{Name: []string{"修复", "fix", "repair"}, Code: ":pencil2:", Emoji: "✏️"},
	{Name: []string{"安全", "security", "safety"}, Code: ":lock:", Emoji: "🔒"},
	{Name: []string{"发布", "flag"}, Code: ":checkered_flag:", Emoji: "🏁"},
	{Name: []string{"linux"}, Code: ":penguin:", Emoji: "🐧"},
	{Name: []string{"windows"}, Code: ":computer:", Emoji: "💻"},
	{Name: []string{"mac"}, Code: ":apple:", Emoji: "🍎"},
	{Name: []string{"docker"}, Code: ":whale:", Emoji: "🐳"},
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
