package view

///////////////////////////////////////////////////////////////////////////////
// Select

type Select struct {
	ViewBaseWithId
	Model    SelectModel
	Name     string
	Size     int // 0 shows all items, 1 shows a dropdownbox, other values show size items
	Class    string
	Disabled bool
}

func (self *Select) IterateChildren(callback IterateChildrenCallback) {
	self.Model.IterateChildren(self, callback)
}

func (self *Select) Render(ctx *Context) (err error) {
	ctx.Response.XML.OpenTag("select")
	ctx.Response.XML.AttribIfNotDefault("id", self.id)
	ctx.Response.XML.AttribIfNotDefault("class", self.Class)
	ctx.Response.XML.Attrib("name", self.Name)
	if self.Disabled {
		ctx.Response.XML.Attrib("disabled", "disabled")
	}

	size := self.Size

	if self.Model != nil {
		numOptions := self.Model.NumOptions()
		if size == 0 {
			size = numOptions
		}
		ctx.Response.XML.Attrib("size", size)

		for i := 0; i < numOptions; i++ {
			ctx.Response.XML.OpenTag("option")
			ctx.Response.XML.AttribIfNotDefault("value", self.Model.Value(i))
			if self.Model.Selected(i) {
				ctx.Response.XML.Attrib("selected", "selected")
			}
			if self.Model.Disabled(i) {
				ctx.Response.XML.Attrib("disabled", "disabled")
			}
			err = self.Model.RenderItem(i, ctx)
			if err != nil {
				return err
			}
			ctx.Response.XML.CloseTag() // option
		}
	} else {
		ctx.Response.XML.Attrib("size", size)
	}

	ctx.Response.XML.CloseTag() // select
	return nil
}

///////////////////////////////////////////////////////////////////////////////
// Model

type SelectModel interface {
	NumOptions() int
	Value(index int) string
	Selected(index int) bool
	Disabled(index int) bool
	RenderItem(index int, ctx *Context) (err error)
	IterateChildren(parent *Select, callback func(parent View, child View) (next bool))
}

///////////////////////////////////////////////////////////////////////////////
// StringsSelectModel

type StringsSelectModel struct {
	Options        []string
	SelectedOption string
}

func (self *StringsSelectModel) NumOptions() int {
	return len(self.Options)
}

func (self *StringsSelectModel) Value(index int) string {
	return self.Options[index]
}

func (self *StringsSelectModel) Selected(index int) bool {
	return self.Options[index] == self.SelectedOption
}

func (self *StringsSelectModel) Disabled(index int) bool {
	return false
}

func (self *StringsSelectModel) RenderItem(index int, ctx *Context) (err error) {
	ctx.Response.WriteString(self.Options[index])
	return nil
}

func (self *StringsSelectModel) IterateChildren(parent *Select, callback func(parent View, child View) (next bool)) {
}

///////////////////////////////////////////////////////////////////////////////
// IndexedStringsSelectModel

type IndexedStringsSelectModel struct {
	Options []string
	Index   int
}

func (self *IndexedStringsSelectModel) NumOptions() int {
	return len(self.Options)
}

func (self *IndexedStringsSelectModel) Value(index int) string {
	return self.Options[index]
}

func (self *IndexedStringsSelectModel) Selected(index int) bool {
	return index == self.Index
}

func (self *IndexedStringsSelectModel) Disabled(index int) bool {
	return false
}

func (self *IndexedStringsSelectModel) RenderItem(index int, ctx *Context) (err error) {
	ctx.Response.WriteString(self.Options[index])
	return nil
}

func (self *IndexedStringsSelectModel) IterateChildren(parent *Select, callback func(parent View, child View) (next bool)) {
}
