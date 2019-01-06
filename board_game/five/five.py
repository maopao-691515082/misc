import sys, os, socket, json
import board_game

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
        s = socket.socket()
        s.settimeout(10)
        s.connect(("localhost", 9999))
        msg = json.dumps(self.board_game.get_board())
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
        for row in xrange(BOARD_SIZE - 5):
            for col in xrange(BOARD_SIZE - 5):
                for line in (board[row][col : col + 5], [board[row + i][col] for i in xrange(5)], [board[row + i][col + i] for i in xrange(5)]):
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
