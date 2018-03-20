package DOM

import (
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"github.com/brokenbydefault/Nanollet/Numbers"
	"github.com/brokenbydefault/Nanollet/GUI/Storage"
	"html"
)

func UpdateAmount(w *window.Window) error {
	hamm := Numbers.NewHumanFromRaw(Storage.Amount)

	for el, scale := range map[string]int{
		".ammount": 6,
	} {
		balance, err := hamm.ConvertToBase(Numbers.MegaXRB, scale)
		if err != nil {
			return err
		}

		display, err := SelectFirstElement(w, el)
		if err != nil {
			return err
		}

		err = display.SetValue(sciter.NewValue(balance))
		if err != nil {
			return err
		}
	}

	return nil
}

func UpdateNotification(w *window.Window, msg string) {
	box, _ := SelectFirstElement(w, "section.notification")

	nt, err := sciter.CreateElement("button", html.EscapeString(msg))
	if err != nil {
		return
	}

	nt.SetAttr("class", "notification")
	box.Append(nt)
	nt.OnClick(func() {
		nt.SetHtml(" ", sciter.SOH_REPLACE)
		nt.Clear()
	})

}
