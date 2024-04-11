package views

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func FrontPage() g.Node {
	return Page(
		"Canvas",
		"/",
		H1(g.Text(`Solution to problem`)),
		P(g.Text("Do you Have problems? we also had problems")),
		P(g.Raw(`Then we created the <em>canvas</em> app, and now we don't'`)),
	)

}
