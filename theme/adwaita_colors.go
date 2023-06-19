package theme

// This file is generated by adwaita_colors_generator.go
// Please do not edit manually, use:
// go generate ./theme/...
//
// The colors are taken from: https://gnome.pages.gitlab.gnome.org/libadwaita/doc/1.0/named-colors.html

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var adwaitaDarkScheme = map[fyne.ThemeColorName]color.Color{
	theme.ColorBlue:                  color.RGBA{0x78, 0xae, 0xed, 0xff}, // Adwaita color name @accent_color
	theme.ColorGreen:                 color.RGBA{0x8f, 0xf0, 0xa4, 0xff}, // Adwaita color name @success_color
	theme.ColorNameBackground:        color.RGBA{0x24, 0x24, 0x24, 0xff}, // Adwaita color name @window_bg_color
	theme.ColorNameButton:            color.RGBA{0x30, 0x30, 0x30, 0xff}, // Adwaita color name @headerbar_bg_color
	theme.ColorNameError:             color.RGBA{0xff, 0x7b, 0x63, 0xff}, // Adwaita color name @destructive_color
	theme.ColorNameForeground:        color.RGBA{0xff, 0xff, 0xff, 0xff}, // Adwaita color name @window_fg_color
	theme.ColorNameInputBackground:   color.RGBA{0x1e, 0x1e, 0x1e, 0xff}, // Adwaita color name @view_bg_color
	theme.ColorNameOverlayBackground: color.RGBA{0x38, 0x38, 0x38, 0xff}, // Adwaita color name @popover_bg_color
	theme.ColorNamePrimary:           color.RGBA{0x35, 0x84, 0xe4, 0xff}, // Adwaita color name @accent_bg_color
	theme.ColorNameShadow:            color.RGBA{0x00, 0x00, 0x00, 0x5b}, // Adwaita color name @shade_color
	theme.ColorRed:                   color.RGBA{0xff, 0x7b, 0x63, 0xff}, // Adwaita color name @destructive_color
	theme.ColorYellow:                color.RGBA{0xf8, 0xe4, 0x5c, 0xff}, // Adwaita color name @warning_color
}

var adwaitaLightScheme = map[fyne.ThemeColorName]color.Color{
	theme.ColorBlue:                  color.RGBA{0x1c, 0x71, 0xd8, 0xff}, // Adwaita color name @accent_color
	theme.ColorGreen:                 color.RGBA{0x26, 0xa2, 0x69, 0xff}, // Adwaita color name @success_color
	theme.ColorNameBackground:        color.RGBA{0xfa, 0xfa, 0xfa, 0xff}, // Adwaita color name @window_bg_color
	theme.ColorNameButton:            color.RGBA{0xeb, 0xeb, 0xeb, 0xff}, // Adwaita color name @headerbar_bg_color
	theme.ColorNameError:             color.RGBA{0xc0, 0x1c, 0x28, 0xff}, // Adwaita color name @destructive_color
	theme.ColorNameForeground:        color.RGBA{0x00, 0x00, 0x00, 0xcc}, // Adwaita color name @window_fg_color
	theme.ColorNameInputBackground:   color.RGBA{0xff, 0xff, 0xff, 0xff}, // Adwaita color name @view_bg_color
	theme.ColorNameOverlayBackground: color.RGBA{0xff, 0xff, 0xff, 0xff}, // Adwaita color name @popover_bg_color
	theme.ColorNamePrimary:           color.RGBA{0x35, 0x84, 0xe4, 0xff}, // Adwaita color name @accent_bg_color
	theme.ColorNameShadow:            color.RGBA{0x00, 0x00, 0x00, 0x11}, // Adwaita color name @shade_color
	theme.ColorRed:                   color.RGBA{0xc0, 0x1c, 0x28, 0xff}, // Adwaita color name @destructive_color
	theme.ColorYellow:                color.RGBA{0xae, 0x7b, 0x03, 0xff}, // Adwaita color name @warning_color
}
