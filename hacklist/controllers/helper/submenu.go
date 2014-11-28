package helper

type SubMenu []map[string]string

func NewSubMenu() *SubMenu {
	return &SubMenu{}
}
func (this *SubMenu) Set(title, href string) *SubMenu {
	*this = append(*this, map[string]string{`href`: href, `title`: title})
	return this
}
