package khtml

type IElement interface {
	Render(*string)
	AddElement(IElement) IElement
}

type Item struct {
	Elements []IElement
}

type Page struct {
	Item
	Name string
}

func NewPage(name string) *Page {
	p := Page{Name: name}

	return &p
}

func (this *Page) Render(code *string) {

	s1 := `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" >
<head>
    <title>` + this.Name + `</title><body>`

	*code = *code + s1

	// RENDER SUBS
	this.Item.Render(code)

	*code = *code + `</body>
</html>`
}

func (this *Item) AddElement(i IElement) IElement {
	this.Elements = append(this.Elements, i)
	return i
}

func (this *Item) Render(code *string) {

	for i := range this.Elements {
		this.Elements[i].Render(code)
	}

}
