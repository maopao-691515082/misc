import copy
import sys

rand_seed = 1
def rand():
    global rand_seed
    rand_seed = (214013 * rand_seed + 2531011) & 0xFFFFFFFF
    return int((rand_seed >> 16) & 0x7FFF)

class CFCStat:
    def __init__(self, card, switch, wait, prev = None, op = None):
        self.card = card
        self.card.sort()
        self.switch = switch
        self.switch.sort()
        self.wait = wait
        self.prev = prev
        self.op = op
        self._calc_hash()
        self.steped = False

        if sum(self.wait) == 198:
            self._print_solution()
            sys.exit(0)

    def _print_solution(self):
        op_list = []
        st = self
        while True:
            if st.op is None:
                break
            op_list.append(st.op)
            st = st.prev
        for i in xrange(len(op_list) - 1, -1, -1):
            print op_list[i]
        print success

    def __hash__(self):
        return self.hash

    def __eq__(self, other):
        return (self.hash == other.hash and self.card == other.card and
                self.switch == other.switch and self.wait == other.wait)

    def __ne__(self, other):
        return not self == other

    def _calc_hash(self):
        h = 0
        i = 0
        for l in self.card:
            for c in l:
                i += 1
                h += i * c
        for c in self.switch:
            i += 1
            h += i * c
        self.hash = h

    def _c_2_str(self, c):
        hl = "CDHS"
        nl = "A23456789TJQK"
        return hl[c % 4] + nl[c / 4]

    def output(self):
        for l in self.card:
            for c in l:
                print self._c_2_str(c),
            print
        for c in self.switch:
            print self._c_2_str(c),
        print
        for c in self.wait:
            print self._c_2_str(c),
        print

    def _can_put_below(self, c, a):
        nc = c / 4
        na = a / 4
        if nc + 1 != na:
            return False
        hc = c % 4
        ha = a % 4
        return hc != ha and hc + ha != 3

    def step(self):
        if self.steped:
            return
        self.steped = True
        #check A and 2
        for c in self.switch:
            if c < 8 and self.wait[c & 3] == c:
                card = copy.deepcopy(self.card)
                switch = self.switch[:]
                wait = self.wait[:]
                switch.remove(c)
                wait[c & 3] += 4
                new_st = CFCStat(card, switch, wait, self,
                                 "%s ok" % self._c_2_str(c))
                if new_st not in d:
                    d.add(new_st)
                    return
                del card
                del switch
                del wait
        for i in xrange(8):
            if len(self.card[i]) == 0:
                continue
            c = self.card[i][-1]
            if c < 8 and self.wait[c & 3] == c:
                card = copy.deepcopy(self.card)
                switch = self.switch[:]
                wait = self.wait[:]
                card[i].pop()
                wait[c & 3] += 4
                new_st = CFCStat(card, switch, wait, self,
                                 "%s ok" % self._c_2_str(c))
                if new_st not in d:
                    d.add(new_st)
                    return
                del card
                del switch
                del wait
        #check other
        for c in self.switch:
            if self.wait[c & 3] == c:
                card = copy.deepcopy(self.card)
                switch = self.switch[:]
                wait = self.wait[:]
                switch.remove(c)
                wait[c & 3] += 4
                new_st = CFCStat(card, switch, wait, self,
                                 "%s ok" % self._c_2_str(c))
                if new_st not in d:
                    d.add(new_st)
                del card
                del switch
                del wait
        for i in xrange(8):
            if len(self.card[i]) == 0:
                continue
            c = self.card[i][-1]
            if self.wait[c & 3] == c:
                card = copy.deepcopy(self.card)
                switch = self.switch[:]
                wait = self.wait[:]
                card[i].pop()
                wait[c & 3] += 4
                new_st = CFCStat(card, switch, wait, self,
                                 "%s ok" % self._c_2_str(c))
                if new_st not in d:
                    d.add(new_st)
                del card
                del switch
                del wait
        #card to switch
        if len(self.switch) < 4:
            for i in xrange(8):
                if len(self.card[i]) == 0:
                    continue
                card = copy.deepcopy(self.card)
                switch = self.switch[:]
                wait = self.wait[:]
                c = card[i].pop()
                switch.append(c)
                new_st = CFCStat(card, switch, wait, self,
                                 "%s to switch" % self._c_2_str(c))
                if new_st not in d:
                    d.add(new_st)
                del card
                del switch
                del wait
        #switch to card
        for c in self.switch:
            for i in xrange(8):
                if (len(self.card[i]) == 0 or
                    self._can_put_below(c, self.card[i][-1])):
                    card = copy.deepcopy(self.card)
                    switch = self.switch[:]
                    wait = self.wait[:]
                    switch.remove(c)
                    card[i].append(c)
                    if len(card[i]) == 1:
                        op = "%s to blank" % self._c_2_str(c)
                    else:
                        op = ("%s below %s" %
                              (self._c_2_str(c), self._c_2_str(card[i][-2])))
                    new_st = CFCStat(card, switch, wait, self, op)
                    if new_st not in d:
                        d.add(new_st)
                    del card
                    del switch
                    del wait
        #card to card
        for i in xrange(8):
            if len(self.card[i]) == 0:
                continue
            c = self.card[i][-1]
            for j in xrange(8):
                if i == j:
                    continue
                if (len(self.card[j]) == 0 or
                    self._can_put_below(c, self.card[j][-1])):
                    card = copy.deepcopy(self.card)
                    switch = self.switch[:]
                    wait = self.wait[:]
                    card[i].pop()
                    card[j].append(c)
                    if len(card[j]) == 1:
                        op = "%s to blank" % self._c_2_str(c)
                    else:
                        op = ("%s below %s" %
                              (self._c_2_str(c), self._c_2_str(card[j][-2])))
                    new_st = CFCStat(card, switch, wait, self, op)
                    if new_st not in d:
                        d.add(new_st)
                    del card
                    del switch
                    del wait

d = set()

def main():
    global d
    order = range(52)
    card = [[] for i in xrange(8)]
    for i in xrange(52):
        si = rand() % (52 - i)
        di = i % 8
        card[di].append(order[si])
        order[si] = order[51 - i]
    d.add(CFCStat(card, [], [0, 1, 2, 3]))
    st_num = len(d)
    while True:
        print st_num
        for st in list(d):
            st.step()
        new_st_num = len(d)
        if st_num == new_st_num:
            print "no solution"
            break
        st_num = new_st_num

if __name__ == "__main__":
    main()
