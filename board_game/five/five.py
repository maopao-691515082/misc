import sys, os, socket, json, copy
import board_game

AI_DEBUG = False

BOARD_SIZE = 15

GAME_STAT_HUMAN = "human thinking"
GAME_STAT_AI    = "ai thinking"
GAME_STAT_OVER  = "game over"

class _Game:
    def __init__(self):
        img_map = {stat: os.path.join("image", "%d.png" % stat) for stat in (0, 1, 2)}
        self.board_game = board_game.Game(BOARD_SIZE, img_map, self.init, self.on_touch_down)
        self.board_game.run()

    def init(self):
        for row in xrange(BOARD_SIZE):
            for col in xrange(BOARD_SIZE):
                self.set_cell_stat(row, col, 0)
        self.set_stat(GAME_STAT_HUMAN)

    def set_cell_stat(self, row, col, stat):
        if stat != 0:
            print "%-10s%s" % ("human" if stat == 1 else "ai", board_game.fmt_pos(row, col))
        self.board_game.set_cell_stat(row, col, stat, set_play_seq = stat != 0)

    def notify(self, text):
        self.board_game.notify(text)

    def set_stat(self, stat):
        self.stat = stat
        self.notify(stat)
        if AI_DEBUG and stat == GAME_STAT_HUMAN:
            self.board_game.schedule_once(self._ai_choice_sched_cb)

    def on_touch_down(self, row, col):
        board = self.board_game.get_board()
        if not AI_DEBUG and self.stat == GAME_STAT_HUMAN and board[row][col] == 0:
            self.set_cell_stat(row, col, 1)
            self.check_over()
            if self.stat != GAME_STAT_OVER:
                self.set_stat(GAME_STAT_AI)
                self.board_game.schedule_once(self._ai_choice_sched_cb)

    def _ai_choice_sched_cb(self):
        if self.stat == GAME_STAT_HUMAN:
            assert AI_DEBUG
            board = self.board_game.get_board()
            rev_board = copy.deepcopy(board)
            for row in xrange(BOARD_SIZE):
                for col in xrange(BOARD_SIZE):
                    if rev_board[row][col] != 0:
                        rev_board[row][col] = 3 - rev_board[row][col]
            row, col = self.ai_choice(rev_board)
            assert board[row][col] == 0
            self.set_cell_stat(row, col, 1)
            self.check_over()
            if self.stat != GAME_STAT_OVER:
                self.set_stat(GAME_STAT_AI)
                self.board_game.schedule_once(self._ai_choice_sched_cb)
        else:
            board = self.board_game.get_board()
            row, col = self.ai_choice(board)
            assert board[row][col] == 0
            self.set_cell_stat(row, col, 2)
            self.check_over()
            if self.stat != GAME_STAT_OVER:
                self.set_stat(GAME_STAT_HUMAN)

    def ai_choice(self, board):
        s = socket.socket()
        s.settimeout(10)
        s.connect(("localhost", 9999))
        msg = json.dumps(board)
        s.sendall(msg)
        s.shutdown(socket.SHUT_WR)
        rsp = ""
        while True:
            data = s.recv(10000)
            if data == "":
                break
            rsp += data
        row, col = json.loads(rsp)
        if row >= BOARD_SIZE or col >= BOARD_SIZE:
            raise Exception("invalid row or col")
        return row, col

    def check_over(self):
        board = self.board_game.get_board()
        for stat_row in board:
            for stat in stat_row:
                if stat == 0:
                    break
            else:
                continue
            break
        else:
            self.set_stat(GAME_STAT_OVER)
            return
        def get_winner(line):
            assert len(line) == 5
            s = set([stat for stat in line])
            if len(s) == 1:
                return list(s)[0]
        for row in xrange(BOARD_SIZE):
            for col in xrange(BOARD_SIZE):
                line_list = []
                if col + 5 <= BOARD_SIZE:
                    line_list.append(board[row][col : col + 5])
                if row + 5 <= BOARD_SIZE:
                    line_list.append([board[row + i][col] for i in xrange(5)])
                if row + 5 <= BOARD_SIZE and col + 5 <= BOARD_SIZE:
                    line_list.append([board[row + i][col + i] for i in xrange(5)])
                if row + 5 <= BOARD_SIZE and col >= 4:
                    line_list.append([board[row + i][col - i] for i in xrange(5)])
                for line in line_list:
                    winner = get_winner(line)
                    if winner == 1 or winner == 2:
                        self.set_stat(GAME_STAT_OVER)
                        return

def main():
    prog_dir = os.path.dirname(sys.argv[0])
    if prog_dir:
        os.chdir(prog_dir)

    _Game()

if __name__ == '__main__':
    main()
