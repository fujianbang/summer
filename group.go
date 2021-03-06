package summer

import "path"

// IGroup is the interface for a group
type IGroup interface {
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)
}

type Group struct {
	core   *Core
	prefix string
}

func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		prefix: prefix,
	}
}

func (g *Group) Get(uri string, handler ControllerHandler) {
	uri = path.Join(g.prefix, uri)
	g.core.Get(uri, handler)
}

func (g *Group) Post(uri string, handler ControllerHandler) {
	uri = path.Join(g.prefix, uri)
	g.core.Post(uri, handler)
}

func (g *Group) Put(uri string, handler ControllerHandler) {
	uri = path.Join(g.prefix, uri)
	g.core.Put(uri, handler)
}

func (g *Group) Delete(uri string, handler ControllerHandler) {
	uri = path.Join(g.prefix, uri)
	g.core.Delete(uri, handler)
}

func (g *Group) Group(prefix string) IGroup {
	prefix = path.Join(g.prefix, prefix)
	return NewGroup(g.core, prefix)
}
