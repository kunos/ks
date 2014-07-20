package khtml

type Paragraph struct {
	Item
	Text string
}

func NewParagraph(text string) *Paragraph {
	p := Paragraph{Text: text}

	return &p
}

func (this *Paragraph) Render(code *string) {
	*code = *code + `<p>` + this.Text

	this.Item.Render(code)
	*code = *code + `</p>`

}
