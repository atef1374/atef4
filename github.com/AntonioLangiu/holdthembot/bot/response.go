package bot

import (
    "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendText(update tgbotapi.Update, ctx *BotContext, text string) {
	response := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	response.ParseMode = "HTML"
	ctx.Bot.Send(response)
}

func StartMessage(update tgbotapi.Update, ctx *BotContext) {
	out := "من ربات   @videomassagexbot   هستم "
	out += "ویدیو مسیج های شما رو به ویدیو و فایل قابل دانلود تبدیل میکنم"
	
	    out += " ویدیو مسیج های خود را فقط در ربات فوروارد کنید"

    out += "در صورت کار نکردن ربات به آیدی زیر اطلاع دهید تا مشکل را برطرف کنیم "
	   
	
	
   out += "@vmxadmin "
	
	
	
	
	out += "    کانال تلگرام ویدیو مسیج ما         "
	
		  out += "@videomassagex "



	
	
    SendText(update, ctx, out)
}

func HelpMessage(update tgbotapi.Update, ctx *BotContext) {
	out := "من ربات   @videomassagexbot   هستم "
	out += "ویدیو مسیج های شما رو به ویدیو و فایل قابل دانلود تبدیل میکنم"
	
	    out += " ویدیو مسیج های خود را فقط در ربات فوروارد کنید"

    out += "در صورت کار نکردن ربات به آیدی زیر اطلاع دهید تا مشکل را برطرف کنیم "
	   
	  out += "@vmxadmin "
	
	out += "کانال تلگرام ویدیو مسیج ما "
	
		  out += "@videomassagex "
	SendText(update, ctx, out)
}
