package main

import (
	"embed"
	"io/ioutil"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
    "github.com/wailsapp/wails/v2/pkg/options/linux"
    "github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Read the contents of the Slimev1.png file into a byte slice
    iconBytes, icon_err := ioutil.ReadFile("./build/appicon.png");
    if icon_err != nil {
        panic(icon_err)
    };

	// Create an instance of the app structure
	app := NewApp();
	

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Dungeoneer",
		Width:  409,
		Height: 627,
		DisableResize: true,
		//WindowIcon:  "./frontend/src/assets/Slimev1.png",
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},

		Windows: &windows.Options{
		WebviewIsTransparent:              false,
		WindowIsTranslucent:               false,
		BackdropType:                      windows.Mica,
		DisableWindowIcon:                 false,
		DisableFramelessWindowDecorations: false,
		WebviewUserDataPath:               "",
		WebviewBrowserPath:                "",
		Theme:                             windows.SystemDefault,
		CustomTheme: &windows.ThemeSettings{
			DarkModeTitleBar:   windows.RGB(20, 20, 20),
			DarkModeTitleText:  windows.RGB(200, 200, 200),
			DarkModeBorder:     windows.RGB(20, 0, 20),
			LightModeTitleBar:  windows.RGB(200, 200, 200),
			LightModeTitleText: windows.RGB(20, 20, 20),
			LightModeBorder:    windows.RGB(200, 200, 200),
			},
		},
        Linux: &linux.Options{
            Icon: iconBytes,
            WindowIsTranslucent: false,
            WebviewGpuPolicy: linux.WebviewGpuPolicyAlways,
			ProgramName: "Dungeoneer",
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
