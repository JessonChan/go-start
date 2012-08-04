package view

import (
	"fmt"
	"reflect"
)

type LinkModel interface {
	URL
	LinkContent(response *Response) View
	LinkTitle(response *Response) string
	LinkRel(response *Response) string
}

func NewLinkModel(url interface{}, content ...interface{}) LinkModel {
	getContent := func() View {
		if len(content) == 0 {
			return NewView(url)
		}
		return NewViews(content...)
	}
	getContentString := func() string {
		if len(content) == 0 {
			return ""
		}
		s, _ := content[0].(string)
		return s
	}
	switch s := url.(type) {
	case **Page:
		return &PageLink{Page: s, Content: getContent(), Title: getContentString()}
	case *ViewWithURL:
		return &URLLink{Url: IndirectViewWithURL(s), Content: NewViews(content...), Title: getContentString()}
	case LinkModel:
		if len(content) > 0 {
			return &URLLink{Url: s, Content: NewViews(content...), Title: getContentString()}
		}
		return s
	case URL:
		return &URLLink{Url: s, Content: getContent(), Title: getContentString()}
	case fmt.Stringer:
		return &StringLink{Url: s.String(), Content: getContent(), Title: getContentString()}
	}
	v := reflect.ValueOf(url)
	if v.Kind() != reflect.String {
		panic(fmt.Errorf("Invalid type for url: %T", url))
	}
	return &StringLink{Url: v.String(), Content: getContent(), Title: getContentString()}
}

func NewLinkModelRel(url interface{}, rel string, content ...interface{}) LinkModel {
	getContent := func() View {
		if len(content) == 0 {
			return NewView(url)
		}
		return NewViews(content...)
	}
	getContentString := func() string {
		if len(content) == 0 {
			return ""
		}
		s, _ := content[0].(string)
		return s
	}
	switch s := url.(type) {
	case **Page:
		return &PageLink{Page: s, Content: getContent(), Rel: rel, Title: getContentString()}
	case *ViewWithURL:
		return &URLLink{Url: IndirectViewWithURL(s), Content: NewViews(content...), Rel: rel, Title: getContentString()}
	case LinkModel:
		if len(content) > 0 {
			return &URLLink{Url: s, Content: NewViews(content...), Rel: rel, Title: getContentString()}
		}
		return s
	case URL:
		return &URLLink{Url: s, Content: getContent(), Rel: rel, Title: getContentString()}
	case fmt.Stringer:
		return &StringLink{Url: s.String(), Content: getContent(), Rel: rel, Title: getContentString()}
	}
	v := reflect.ValueOf(url)
	if v.Kind() != reflect.String {
		panic(fmt.Errorf("Invalid type for url: %T", url))
	}
	return &StringLink{Url: v.String(), Content: getContent(), Rel: rel}
}
