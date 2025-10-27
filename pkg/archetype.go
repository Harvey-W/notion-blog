package notion_blog

import (
	"log"
	"time"

	"github.com/jomei/notionapi"
)

type ArchetypeFields struct {
	Title        string
	Description  string
	Banner       string
	CreationDate time.Time
	LastModified time.Time
	Author       string
	Tags         []notionapi.Option
	Categories   []notionapi.Option
	Content      string
	Properties   notionapi.Properties
}

func MakeArchetypeFields(p notionapi.Page, config BlogConfig) ArchetypeFields {
	a := ArchetypeFields{
		Title:        "",
		Description:  "",
		Banner:       "",
		CreationDate: p.CreatedTime,
		LastModified: p.LastEditedTime,
		Author:       "",
	}

	// --- Title ---
	if nameProp, ok := p.Properties["Name"].(*notionapi.TitleProperty); ok {
		a.Title = ConvertRichText(nameProp.Title)
	} else {
		log.Println("warning: missing or invalid 'Name' property")
	}

	// --- Author ---
	if createdByProp, ok := p.Properties["Created By"].(*notionapi.CreatedByProperty); ok && createdByProp.CreatedBy != nil {
		a.Author = createdByProp.CreatedBy.Name
	} else if authorProp, ok := p.Properties["Author"].(*notionapi.PeopleProperty); ok && len(authorProp.People) > 0 {
		a.Author = authorProp.People[0].Name
	} else {
		log.Println("warning: no 'Created By' or 'Author' property found; leaving Author blank")
	}

	// --- Banner ---
	if p.Cover != nil && p.Cover.GetURL() != "" {
		coverSrc, _ := getImage(p.Cover.GetURL(), config)
		a.Banner = coverSrc
	}

	// --- Description ---
	if v, ok := p.Properties[config.PropertyDescription]; ok {
		if text, ok := v.(*notionapi.RichTextProperty); ok {
			a.Description = ConvertRichText(text.RichText)
		} else {
			log.Println("warning: given property description is not a text property")
		}
	}

	// --- Categories ---
	if v, ok := p.Properties[config.PropertyCategories]; ok {
		if multiSelect, ok := v.(*notionapi.MultiSelectProperty); ok {
			a.Categories = multiSelect.MultiSelect
		} else {
			log.Println("warning: given property categories is not a multi-select property")
		}
	}

	// --- Tags ---
	if v, ok := p.Properties[config.PropertyTags]; ok {
		if multiSelect, ok := v.(*notionapi.MultiSelectProperty); ok {
			a.Tags = multiSelect.MultiSelect
		} else {
			log.Println("warning: given property tags is not a multi-select property")
		}
	}

	a.Properties = p.Properties
	return a
}
