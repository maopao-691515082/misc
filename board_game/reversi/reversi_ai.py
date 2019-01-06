import random, copy
import reversi

def try_set_cell(board, row, col, stat, only_judge = False):
    assert board[row][col] == 0
    score = 0
    for x in (-1, 0, 1):
        for y in (-1, 0, 1):
            if (x, y) != (0, 0):
                r, c = row + x, col + y
                l = []
                while 0 <= r < reversi.BOARD_SIZE and 0 <= c < reversi.BOARD_SIZE:
                    if board[r][c] == stat:
                        break
                    if board[r][c] == 0:
                        l = []
                        break
                    assert board[r][c] == 3 - stat
                    l.append((r, c))
                    r += x
                    c += y
                else:
                    l = []
                if l:
                    if only_judge:
                        return 1
                    for r, c in l:
                        board[r][c] = stat
                        score += 1
    if score:
        board[row][col] = stat
    return score

def _ai_choice(board, stat, deep, ab_score_result):
    if deep <= 0:
        return None, 0

    valid_choice_list = []
    for row in xrange(reversi.BOARD_SIZE):
        for col in xrange(reversi.BOARD_SIZE):
            if board[row][col] == 0 and try_set_cell(board, row, col, 2, only_judge = True):
                valid_choice_list.append((row, col))
    random.shuffle(valid_choice_list)

    pos_result = None, None
    score_result = None
    for row, col in valid_choice_list:
        new_board = copy.deepcopy(board)
        score = try_set_cell(new_board, row, col, stat)
        _, score_add = _ai_choice(new_board, 3 - stat, deep - 1, score_result)
        score += score_add
        if stat == 2:
            if score_result is None or score > score_result:
                pos_result = row, col
                score_result = score
            if ab_score_result is not None and score_result > ab_score_result:
                break
        else:
            assert stat == 1
            if score_result is None or score < score_result:
                pos_result = row, col
                score_result = score
            if ab_score_result is not None and score_result < ab_score_result:
                break
    if score_result is None:
        score_result = 0
    return pos_result, score_result

def ai_choice(board):
    pos, _ = _ai_choice(board, 2, 4, None)
    return pos
