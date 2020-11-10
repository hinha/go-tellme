package button

import tb "gopkg.in/tucnak/telebot.v2"

var (
	// Universal markup builders.
	Menu     = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
	Selector = &tb.ReplyMarkup{}

	// Reply buttons.
	BtnInfo     = Menu.Text("ℹ Info")
	BtnSettings = Menu.Text("⚙ Settings")
	BtnToken    = Menu.Text("🔑 Token")
	BtnHelp     = Menu.Text("🆘 Help")

	BtnPrev = Selector.Data("⬅", "prev", "a")
	BtnNext = Selector.Data("➡", "next", "b")
)
