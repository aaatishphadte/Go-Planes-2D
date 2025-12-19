package scene

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/m110/airplanes/assets"
)

type Credits struct {
	screenWidth           int
	screenHeight          int
	returnToTitleCallback func()
}

func NewCredits(screenWidth int, screenHeight int, returnToTitleCallback func()) *Credits {
	return &Credits{
		screenWidth:           screenWidth,
		screenHeight:          screenHeight,
		returnToTitleCallback: returnToTitleCallback,
	}
}

func (c *Credits) Update() {
	if ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsKeyPressed(ebiten.KeyEnter) {
		c.returnToTitleCallback()
		return
	}

	touchIDs := inpututil.AppendJustPressedTouchIDs(nil)
	if len(touchIDs) > 0 {
		c.returnToTitleCallback()
		return
	}
}

func (c *Credits) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{10, 10, 30, 255})

	// Title
	text.Draw(screen, "CONGRATULATIONS!", assets.NarrowFont, c.screenWidth/6, 80, color.RGBA{255, 215, 0, 255})
	text.Draw(screen, "Thanks for Playing!", assets.NormalFont, c.screenWidth/6, 140, color.White)

	// Credits
	text.Draw(screen, "CREDITS", assets.NormalFont, c.screenWidth/3, 220, color.RGBA{100, 200, 255, 255})

	text.Draw(screen, "Game Completed By:", assets.SmallFont, c.screenWidth/6, 280, color.RGBA{200, 200, 200, 255})
	text.Draw(screen, "Dattatray Fadte", assets.NormalFont, c.screenWidth/5, 310, color.White)

	text.Draw(screen, "Original Author:", assets.SmallFont, c.screenWidth/6, 370, color.RGBA{200, 200, 200, 255})
	text.Draw(screen, "m110", assets.NormalFont, c.screenWidth/5, 400, color.White)

	text.Draw(screen, "Assets By:", assets.SmallFont, c.screenWidth/6, 460, color.RGBA{200, 200, 200, 255})
	text.Draw(screen, "Kenney", assets.NormalFont, c.screenWidth/5, 490, color.White)

	text.Draw(screen, "Powered By:", assets.SmallFont, c.screenWidth/6, 550, color.RGBA{200, 200, 200, 255})
	text.Draw(screen, "Ebiten Engine", assets.NormalFont, c.screenWidth/5, 580, color.White)

	// Instructions
	text.Draw(screen, "Press Space to return to Title", assets.NormalFont, c.screenWidth/8, 700, color.RGBA{150, 150, 150, 255})
}
