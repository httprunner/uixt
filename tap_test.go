package uixt

import (
	"testing"

	"github.com/electricbubble/gwda"
)

func TestDriverExt_TapWithNumber(t *testing.T) {
	driver, err := gwda.NewUSBDriver(nil)
	checkErr(t, err)

	driverExt, err := Extend(driver, 0.95)
	checkErr(t, err)

	pathSearch := "/Users/hero/Documents/temp/2020-05/opencv/flag7.png"

	// gwda.SetDebug(true)

	err = driverExt.TapWithNumber(pathSearch, 3)
	checkErr(t, err)

	err = driverExt.TapWithNumberOffset(pathSearch, 3, 0.5, 0.75)
	checkErr(t, err)
}

func TestDriverExt_TapWithCoordinate(t *testing.T) {
	driver, err := gwda.NewUSBDriver(nil)
	checkErr(t, err)

	driverExt, err := Extend(driver, 0.95)
	checkErr(t, err)

	err = driverExt.Tap([]float64{0.4, 0.5})
	checkErr(t, err)
}

func TestDriverExt_TapWithOCR(t *testing.T) {
	driver, err := gwda.NewUSBDriver(nil)
	checkErr(t, err)

	driverExt, err := Extend(driver, 0.95)
	checkErr(t, err)

	// 需要点击文字上方的图标
	err = driverExt.TapOffset("抖音", 0.5, -1)
	checkErr(t, err)
}
