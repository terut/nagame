package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"
	"os"
)

func main() {
	var imageModel = NewImageModel(nil)
	var presenter = NewImagePresenter(nil)
	presenter.SetImageModel(imageModel)

	//var ir = new(ImageReader)
	//var images, _ = ir.ReadDir("")
	//imageModel.SetImageFile(images)

	// Create application
	app := gui.NewQGuiApplication(len(os.Args), os.Args)

	// Enable high DPI scaling
	app.SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	// Use the material style for qml
	quickcontrols2.QQuickStyle_SetStyle("material")

	// Create a QML application engine
	engine := qml.NewQQmlApplicationEngine(nil)

	//engine.RootContext().SetContextProperty("imageModel", imageModel)
	engine.RootContext().SetContextProperty("presenter", presenter)

	// Load the main qml file
	engine.Load(core.NewQUrl3("qrc:/qml/main.qml", 0))

	// Execute app
	gui.QGuiApplication_Exec()
}
