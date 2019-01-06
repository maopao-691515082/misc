#qpy:kivy

import sys, os
import board_game

BOARD_SIZE = 8

GAME_STAT_HUMAN = "human thinking"
GAME_STAT_AI    = "ai thinking"
GAME_STAT_OVER  = "game over"

class _Game:
    def __init__(self):
        img_map = {stat: os.path.join("image", "%d.png" % (0 if stat is None else stat)) for stat in (None, 0, 1, 2)}
        self.board_game = board_game.Game(BOARD_SIZE, img_map, self.init, self.on_touch_down)
        self.board_game.run()

    def init(self):
        assert BOARD_SIZE % 2 == 0
        for row in xrange(BOARD_SIZE):
            for col in xrange(BOARD_SIZE):
                self.set_cell_stat(row, col, 0)
        self.set_cell_stat(BOARD_SIZE / 2 - 1, BOARD_SIZE / 2 - 1, 1)
        self.set_cell_stat(BOARD_SIZE / 2, BOARD_SIZE / 2, 1)
        self.set_cell_stat(BOARD_SIZE / 2 - 1, BOARD_SIZE / 2, 2)
        self.set_cell_stat(BOARD_SIZE / 2, BOARD_SIZE / 2 - 1, 2)
        self.set_stat(GAME_STAT_HUMAN)

    def set_cell_stat(self, row, col, stat):
        if stat != 0:
            print "%-10s%s" % ("human" if stat == 1 else "ai", board_game.fmt_pos(row, col))
        self.board_game.set_cell_stat(row, col, stat)

    def notify(self, text):
        self.board_game.notify(text)

    def set_stat(self, stat):
        self.stat = stat
        self.notify(stat)

    def on_touch_down(self, row, col):
        board = self.board_game.get_board()
        if self.stat == GAME_STAT_HUMAN and board[row][col] == 0:
            self.set_cell_stat(row, col, 1)
            self.check_over()
            if self.stat != GAME_STAT_OVER:
                self.set_stat(GAME_STAT_AI)
                self.board_game.schedule_once(self._ai_choice_sched_cb)

    def _ai_choice_sched_cb(self):
        row, col = self.ai_choice()
        assert self.board_game.get_board()[row][col] == 0
        self.set_cell_stat(row, col, 2)
        self.check_over()
        if self.stat != GAME_STAT_OVER:
            self.set_stat(GAME_STAT_HUMAN)

    def ai_choice(self):
        raise

    def check_over(self):
        raise

def main():
    prog_dir = os.path.dirname(sys.argv[0])
    if prog_dir:
        os.chdir(prog_dir)

    _Game()

if __name__ == '__main__':
    main()
