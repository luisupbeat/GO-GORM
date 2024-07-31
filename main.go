package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/luisupbeat/go-gorm/controllers"
	db "github.com/luisupbeat/go-gorm/db"
	"github.com/luisupbeat/go-gorm/models"
)

func main() {
	db.DBConnection()
	fmt.Println("Hello men")
	controllers.GetHaciendas()
	myApp := app.New()
	w := myApp.NewWindow("box Layout")

	haciendas, err := controllers.GetHaciendas()
	if err != nil {
		fmt.Println("Error al obtener las haciendas:", err)
		return
	}

	nuevaHacienda := models.Hacienda{
		CodHacienda:   7300,
		NomHacienda:   "Hacienda de Prueba",
		NomProductor:  "Juan Pérez",
		ProvHacienda:  "Córdoba",
		DirHacienda:   "Calle Falsa 123",
		TelHacienda:   "123456789",
		Tel1Hacienda:  "987654321",
		EstHacAct:     1,
		PRODUCTOR_COD: "P001",
		TipoHac:       2,
	}

	controllers.SaveHacienda(nuevaHacienda)

	list := widget.NewList(
		func() int {
			return len(haciendas)
		},
		func() fyne.CanvasObject {
			label := widget.NewLabel("")
			return container.NewHBox(label)
		},
		func(i int, obj fyne.CanvasObject) {
			ct := obj.(*fyne.Container)
			label := ct.Objects[0].(*widget.Label)
			hacienda := haciendas[i]
			label.SetText(hacienda.NomHacienda)
		})

	w.SetContent(container.NewScroll(list))
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
