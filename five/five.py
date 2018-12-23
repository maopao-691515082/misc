#qpy:kivy

import sys, os, random

from kivy.app import App
from kivy.uix.widget import Widget
from kivy.uix.image import Image
from kivy.graphics import Color, Rectangle

def _assert(e):
    if not e:
        raise Exception("Bug")

CELL_IMG_SIZE = 51
CELL_SIZE = CELL_IMG_SIZE * 4 / 3

class Cell(Image):

    def __init__(self, row, col, **kwarg):
        Image.__init__(self, **kwarg)

        self.size = CELL_SIZE, CELL_SIZE
        self.allow_stretch = True
        self.keep_ratio = False

        self.row, self.col = row, col

        self.pos = 100 + row * CELL_SIZE, 100 + col * CELL_SIZE

    def on_touch_down(self, touch):
        if self.collide_point(*touch.pos) and touch.button == "left":
            self.source = os.path.join("image", "1.png")
            return True

BOARD_SIZE = 15

class FiveGame(Widget):
    def init(self):
        self.width = self.height = 1000

        with self.canvas:
            Color(rgb = [0.0, 0.3, 0.0])
            Rectangle(size = self.size)

        board = [[0] * BOARD_SIZE for _ in xrange(BOARD_SIZE)]

        for row in xrange(BOARD_SIZE):
            for col in xrange(BOARD_SIZE):
                cell = Cell(row, col, source = os.path.join("image", "0.png"))
                self.add_widget(cell)

game = None

class FiveApp(App):
    def build(self):
        global game
        game = FiveGame()
        return game

    def on_start(self):
        game.init()

    def on_pause(self):
        return True

if __name__ == '__main__':
    prog_dir = os.path.dirname(sys.argv[0])
    if prog_dir:
        os.chdir(prog_dir)
    app = FiveApp()
    app.run()
