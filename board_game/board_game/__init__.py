#qpy:kivy

import string

import kivy.app
import kivy.uix.widget
import kivy.uix.image
import kivy.uix.label
import kivy.graphics
import kivy.clock

class _Cell(kivy.uix.image.Image):
    def __init__(self, intf, row, col, **kwarg):
        kivy.uix.image.Image.__init__(self, **kwarg)

        self.intf = intf

        self.size = intf.cell_size, intf.cell_size
        self.allow_stretch = True
        self.keep_ratio = False

        self.row, self.col = row, col

        self.pos = 100 + row * intf.cell_size, 100 + col * intf.cell_size

        self.set_stat(None)

    def set_stat(self, stat):
        self.stat = stat
        self.source = self.intf.img_map[stat]

    def on_touch_down(self, touch):
        if self.collide_point(*touch.pos) and touch.button == "left":
            self.intf.touch_down_cb(self.row, self.col)
            return True

class _Intf(kivy.uix.widget.Widget):
    def __init__(self, board_size, img_map, init_game, touch_down_cb):
        kivy.uix.widget.Widget.__init__(self)

        self.board_size = board_size
        self.img_map = img_map
        self.init_game = init_game
        self.touch_down_cb = touch_down_cb

        self.cell_size = 900 / board_size
        if self.cell_size > 100:
            self.cell_size = 100

    def init(self):
        with self.canvas:
            kivy.graphics.Color(rgb = [0.5, 0.5, 0.5])
            kivy.graphics.Rectangle(size = self.size)

        self.board = [[None] * self.board_size for _ in xrange(self.board_size)]

        for row in xrange(self.board_size):
            for col in xrange(self.board_size):
                cell = _Cell(self, row, col)
                self.board[row][col] = cell
                self.add_widget(cell)

        for i in xrange(self.board_size):
            x, y = 100 + i * self.cell_size, 100 - self.cell_size
            for pos in (x, y), (y, x):
                p = kivy.uix.label.Label(text = _fmt_pos_num(i))
                p.size = self.cell_size, self.cell_size
                p.bold = True
                p.font_size = self.cell_size
                p.color = [1, 1, 1, 1]
                p.pos = pos
                self.add_widget(p)

        self.notify_pad = kivy.uix.label.Label(text = "Notify")
        self.notify_pad.size = 300, 50
        self.notify_pad.bold = True
        self.notify_pad.font_size = 50
        self.notify_pad.color = [1, 1, 1, 1]
        self.notify_pad.pos = 1200, 550
        self.add_widget(self.notify_pad)

        self.init_game()

    def notify(self, text):
        self.notify_pad.text = text

    def get_board(self):
        return [[cell.stat for cell in cell_row] for cell_row in self.board]

    def set_cell_stat(self, row, col, stat):
        self.board[row][col].set_stat(stat)

class _App(kivy.app.App):
    def __init__(self, game, *args):
        kivy.app.App.__init__(self)

        self.game = game
        self.args = args

    def build(self):
        self.game.intf = _Intf(*self.args)
        return self.game.intf

    def on_start(self):
        self.game.intf.init()

    def on_pause(self):
        return True

class Game:
    def __init__(self, *args):
        self.app = _App(self, *args)
        self.intf = None

    def run(self):
        self.app.run()

    def notify(self, text):
        self.intf.notify(text)

    def schedule_once(self, cb):
        kivy.clock.Clock.schedule_once(lambda dt: cb())

    def get_board(self):
        return self.intf.get_board()

    def set_cell_stat(self, row, col, stat):
        self.intf.set_cell_stat(row, col, stat)

def _fmt_pos_num(n):
    return string.letters[n]

def fmt_pos(row, col):
    return "(%s,%s)" % (_fmt_pos_num(row), _fmt_pos_num(col))
