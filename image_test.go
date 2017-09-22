package vips_test

import (
	"testing"

	"github.com/davidbyttow/govips"
)

func TestLoadFromFile(t *testing.T) {
	if testing.Short() {
		return
	}
	vips.LeakTest(func() {
		img, _ := vips.NewImageFromFile("fixtures/canyon.jpg")
		img = img.Resize(0.5)
		// image, err := vips.NewImageFromFile("fixtures/canyon.jpg")
		// require.Nil(t, err)
		// assert.Equal(t, 2560, image.Width())
		// assert.Equal(t, 1600, image.Height())
	})
}

// func TestWriteToFile(t *testing.T) {
// 	vips.LeakTest(func() {
// 		image, err := vips.NewImageFromFile("fixtures/canyon.jpg")
// 		require.Nil(t, err)
//
// 		image = image.Resize(0.25)
//
// 		tempDir, err := ioutil.TempDir("", "TestWriteToFile")
// 		require.Nil(t, err)
// 		defer os.RemoveAll(tempDir)
//
// 		err = image.WriteToFile(tempDir + "/canyon-out.jpg")
// 		require.Nil(t, err)
// 	})
// }
//
// func TestWriteToBytes(t *testing.T) {
// 	vips.LeakTest(func() {
// 		buf, err := ioutil.ReadFile("fixtures/canyon.jpg")
// 		require.Nil(t, err)
//
// 		image, err := vips.NewImageFromBuffer(buf)
// 		require.Nil(t, err)
//
// 		image = image.Resize(0.25)
//
// 		buf, err = image.WriteToBuffer(vips.ImageTypeJPEG)
// 		require.Nil(t, err)
// 		assert.True(t, len(buf) > 0)
//
// 		image, err = vips.NewImageFromBuffer(buf)
// 		require.Nil(t, err)
// 		assert.Equal(t, 640, image.Width())
// 		assert.Equal(t, 400, image.Height())
// 	})
// }
//
// func TestLoadFromMemory(t *testing.T) {
// 	vips.LeakTest(func() {
// 		size := 200
//
// 		bytes := make([]byte, size*size*3)
// 		for i := 0; i < size*size; i++ {
// 			bytes[i*3] = 0xFF
// 			bytes[i*3+1] = 0
// 			bytes[i*3+2] = 0
// 		}
//
// 		image, err := vips.NewImageFromMemory(bytes, size, size, 3, vips.BandFormatUchar)
// 		require.Nil(t, err)
//
// 		tempDir, err := ioutil.TempDir("", "TestLoadFromMemory")
// 		require.Nil(t, err)
// 		defer os.RemoveAll(tempDir)
//
// 		err = image.WriteToFile(tempDir + "red-out.png")
// 		require.Nil(t, err)
// 	})
// }
