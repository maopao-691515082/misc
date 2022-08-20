import sys, os, copy
import reversi_ai

AI_DEBUG = False

BOARD_SIZE = 8

GAME_STAT_HUMAN = "human thinking"
GAME_STAT_AI    = "ai thinking"
GAME_STAT_OVER  = "game over"

class _Game:
    def __init__(self):
        img_map = {stat: os.path.join("image", "%d.png" % stat) for stat in (0, 1, 2)}
        self.board_game = board_game.Game(BOARD_SIZE, img_map, self.init, self.on_touch_down)
        self.board_game.run()

    def init(self):
        assert BOARD_SIZE > 2 and BOARD_SIZE % 2 == 0
        for row in range(BOARD_SIZE):
            for col in range(BOARD_SIZE):
                self.set_cell_stat(row, col, 0)
        for row in (BOARD_SIZE // 2 - 1, BOARD_SIZE // 2):
            for col in (BOARD_SIZE // 2 - 1, BOARD_SIZE // 2):
                self.set_cell_stat(row, col, 1 if row == col else 2)
        assert self.can_go(1)
        self.set_stat(GAME_STAT_HUMAN)

    def set_cell_stat(self, row, col, stat, is_play = False):
        if stat != 0 and is_play:
            print("%-10s%s" % ("human" if stat == 1 else "ai", board_game.fmt_pos(row, col)))
        self.board_game.set_cell_stat(row, col, stat, set_play_seq = is_play)

    def refresh_board(self, board, played_row, played_col):
        for row in range(BOARD_SIZE):
            for col in range(BOARD_SIZE):
                self.set_cell_stat(row, col, board[row][col], (row, col) == (played_row, played_col))

    def notify(self, text):
        stat_list = sum(self.board_game.get_board(), [])
        text += "\nblack:%d\nwhite:%d" % (stat_list.count(1), stat_list.count(2))
        self.board_game.notify(text)

    def set_stat(self, stat):
        self.stat = stat
        self.notify(stat)
        if AI_DEBUG and stat == GAME_STAT_HUMAN:
            self.board_game.schedule_once(self._ai_choice_sched_cb)

    def on_touch_down(self, row, col):
        board = self.board_game.get_board()
        if not AI_DEBUG and self.stat == GAME_STAT_HUMAN and board[row][col] == 0 and reversi_ai.try_set_cell(board, row, col, 1):
            self.refresh_board(board, row, col)
            self.check_over()
            if self.stat != GAME_STAT_OVER:
                if self.can_go(2):
                    self.set_stat(GAME_STAT_AI)
                    self.board_game.schedule_once(self._ai_choice_sched_cb)

    def _ai_choice_sched_cb(self):
        if self.stat == GAME_STAT_HUMAN:
            assert AI_DEBUG
            board = self.board_game.get_board()
            rev_board = copy.deepcopy(board)
            for row in range(BOARD_SIZE):
                for col in range(BOARD_SIZE):
                    if rev_board[row][col] != 0:
                        rev_board[row][col] = 3 - rev_board[row][col]
            row, col = reversi_ai.ai_choice(rev_board)
            assert board[row][col] == 0 and reversi_ai.try_set_cell(board, row, col, 1)
            self.refresh_board(board, row, col)
            self.check_over()
            if self.stat != GAME_STAT_OVER:
                if self.can_go(2):
                    self.set_stat(GAME_STAT_AI)
                    self.board_game.schedule_once(self._ai_choice_sched_cb)
                else:
                    assert self.can_go(1)
                    self.set_stat(GAME_STAT_HUMAN)
        else:
            board = self.board_game.get_board()
            row, col = reversi_ai.ai_choice(copy.deepcopy(board))
            assert board[row][col] == 0 and reversi_ai.try_set_cell(board, row, col, 2)
            self.refresh_board(board, row, col)
            self.check_over()
            if self.stat != GAME_STAT_OVER:
                if self.can_go(1):
                    self.set_stat(GAME_STAT_HUMAN)
                else:
                    assert self.can_go(2) and self.stat == GAME_STAT_AI
                    self.board_game.schedule_once(self._ai_choice_sched_cb)

    def can_go(self, next_stat):
        board = self.board_game.get_board()
        for row in range(BOARD_SIZE):
            for col in range(BOARD_SIZE):
                if board[row][col] == 0 and reversi_ai.try_set_cell(board, row, col, next_stat, only_judge = True):
                    return True
        return False

    def check_over(self):
        if self.can_go(1) or self.can_go(2):
            return
        self.set_stat(GAME_STAT_OVER)

board_game = None

def main():
    global board_game

    prog_dir = os.path.dirname(sys.argv[0])
    if prog_dir:
        os.chdir(prog_dir)
    sys.path.append("..")
    import board_game

    _Game()

if __name__ == '__main__':
    main()
