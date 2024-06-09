package styles

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props"
	"github.com/AccentDesign/gcss/variables"
)

// Resets returns the styles for resets for the base stylesheet.
// From: https://github.com/tailwindlabs/tailwindcss/blob/4aefd26f44cbb5ede9bc02c00085e97550d0953e/packages/tailwindcss/preflight.css
func (ss *StyleSheet) Resets() Styles {
	return Styles{
		{
			Selector: "*,::after,::before,::backdrop,::file-selector-button",
			Props: gcss.Props{
				Border: props.Border{
					Width: variables.Size0,
					Style: props.BorderStyleSolid,
				},
				BoxSizing: props.BoxSizingBorderBox,
				Margin:    variables.Size0,
				Padding:   variables.Size0,
			},
		},
		{
			Selector: "html,:host",
			Props: gcss.Props{
				FontFamily: props.FontFamilySans,
				LineHeight: props.UnitRaw(1.5),
			},
			CustomProps: []gcss.CustomProp{
				{"-webkit-tap-highlight-color", "transparent"},
				{"-webkit-text-size-adjust", "100%"},
				{"tab-size", "4"},
			},
		},
		{
			Selector: "body",
			Props: gcss.Props{
				LineHeight: props.UnitInherit(),
			},
		},
		{
			Selector: "hr",
			Props: gcss.Props{
				Color:  props.ColorInherit(),
				Height: variables.Size0,
			},
			CustomProps: []gcss.CustomProp{
				{"border-top-width", "1px"},
			},
		},
		{
			Selector: "abbr:where([title])",
			Props: gcss.Props{
				TextDecorationLine:  props.TextDecorationLineUnderline,
				TextDecorationStyle: props.TextDecorationStyleDotted,
			},
		},
		{
			Selector: "h1,h2,h3,h4,h5,h6",
			Props: gcss.Props{
				FontSize:   props.UnitInherit(),
				FontWeight: "inherit",
			},
		},
		{
			Selector: "a",
			Props: gcss.Props{
				Color:               props.ColorInherit(),
				TextDecorationLine:  "inherit",
				TextDecorationStyle: "inherit",
			},
		},
		{
			Selector: "b,strong",
			Props: gcss.Props{
				FontWeight: "bolder",
			},
		},
		{
			Selector: "code,kbd,samp,pre",
			Props: gcss.Props{
				FontFamily: props.FontFamilyMono,
				FontSize:   props.UnitEm(1),
			},
		},
		{
			Selector: "small",
			Props: gcss.Props{
				FontSize: props.UnitPercent(80),
			},
		},
		{
			Selector: "sub,sup",
			Props: gcss.Props{
				FontSize:      props.UnitPercent(75),
				LineHeight:    variables.Size0,
				Position:      props.PositionRelative,
				VerticalAlign: props.VerticalAlignBaseline,
			},
		},
		{
			Selector: "sub",
			Props: gcss.Props{
				Bottom: props.UnitEm(-0.25),
			},
		},
		{
			Selector: "sup",
			Props: gcss.Props{
				Top: props.UnitEm(-0.5),
			},
		},
		{
			Selector: "table",
			Props: gcss.Props{
				BorderCollapse: props.BorderCollapseCollapse,
				BorderColor:    props.ColorInherit(),
				TextIndent:     variables.Size0,
			},
		},
		{
			Selector: "button,input,optgroup,select,textarea,::file-selector-button",
			Props: gcss.Props{
				BackgroundColor: props.ColorTransparent(),
				Color:           props.ColorInherit(),
				FontFamily:      "inherit",
				FontSize:        props.UnitInherit(),
			},
			CustomProps: []gcss.CustomProp{
				{"letter-spacing", "inherit"},
			},
		},
		{
			Selector: "input:where(:not([type='button'],[type='reset'],[type='submit'])),select,textarea",
			Props: gcss.Props{
				Border: props.Border{
					Width: props.UnitPx(1),
					Style: props.BorderStyleSolid,
				},
			},
		},
		{
			Selector: "button,input:where([type='button'],[type='reset'],[type='submit']),::file-selector-button",
			Props: gcss.Props{
				Appearance: props.AppearanceButton,
			},
		},
		{
			Selector: ":-moz-focusring",
			CustomProps: []gcss.CustomProp{
				{"outline", "auto"},
			},
		},
		{
			Selector: ":-moz-ui-invalid",
			CustomProps: []gcss.CustomProp{
				{"box-shadow", "none"},
			},
		},
		{
			Selector: "progress",
			Props: gcss.Props{
				VerticalAlign: props.VerticalAlignBaseline,
			},
		},
		{
			Selector: "::-webkit-inner-spin-button,::-webkit-outer-spin-button",
			Props: gcss.Props{
				Height: props.UnitAuto(),
			},
		},
		{
			Selector: "::-webkit-search-decoration",
			CustomProps: []gcss.CustomProp{
				{"-webkit-appearance", "none"},
			},
		},
		{
			Selector: "summary",
			Props: gcss.Props{
				Display: props.DisplayListItem,
			},
		},
		{
			Selector: "ol,ul,menu",
			Props: gcss.Props{
				ListStyleType:     props.ListStyleTypeNone,
				ListStylePosition: props.ListStylePositionInside,
			},
		},
		{
			Selector: "textarea",
			CustomProps: []gcss.CustomProp{
				{"resize", "vertical"},
			},
		},
		{
			Selector: "::placeholder",
			Props: gcss.Props{
				Color:   props.Color{Keyword: "color-mix(in srgb, currentColor 50%, transparent)"},
				Opacity: props.UnitRaw(1),
			},
		},
		{
			Selector: ":disabled",
			Props: gcss.Props{
				Cursor: props.CursorDefault,
			},
		},
		{
			Selector: "img,svg,video,canvas,audio,iframe,embed,object",
			Props: gcss.Props{
				Display:       props.DisplayBlock,
				VerticalAlign: props.VerticalAlignMiddle,
			},
		},
		{
			Selector: "img,video",
			Props: gcss.Props{
				Height:   props.UnitAuto(),
				MaxWidth: variables.Full,
			},
		},
		{
			Selector: "[hidden]",
			Props: gcss.Props{
				Display: "none !important",
			},
		},
	}
}
