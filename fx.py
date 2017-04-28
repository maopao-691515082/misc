import struct
import math

def write_bmp(pm):
    x = len(pm)
    y = len(pm[0])
    pm.reverse()
    dib_line_pad_len = (y * 3 + 3) / 4 * 4 - y * 3
    dib = []
    for l in pm:
        assert len(l) == y
        dib_line = []
        for p in l:
            dib_line.append(struct.pack("=I", p)[: 3])
        dib_line.append("\x00" * dib_line_pad_len)
        dib += dib_line
    dib = "".join(dib)
    cont = struct.pack("=2sIHHIIIIHHIIIIII", "BM", 54 + len(dib), 0, 0, 54,
                       40, y, x, 1, 24, 0, len(dib), 0, 0, 0, 0) + dib
    open("test.bmp", "wb").write(cont)

def draw_line(pm, (x1, y1), (x2, y2)):
    x1 = int(x1)
    y1 = int(y1)
    x2 = int(x2)
    y2 = int(y2)
    if x1 == x2:
        for i in xrange(min(y1, y2), max(y1, y2) + 1):
            pm[x1][i] = 0
        return
    if y1 == y2:
        for i in xrange(min(x1, x2), max(x1, x2) + 1):
            pm[i][y1] = 0
        return
    k = float(y1 - y2) / (x1 - x2)
    a = y1 - k * x1
    for x in xrange(min(x1, x2), max(x1, x2) + 1):
        y = int(round(k * x + a))
        pm[x][y] = 0
    k = float(x1 - x2) / (y1 - y2)
    a = x1 - k * y1
    for y in xrange(min(y1, y2), max(y1, y2) + 1):
        x = int(round(k * y + a))
        pm[x][y] = 0

"""
pm = []
for i in xrange(256):
    l = []
    pm.append(l)
    for j in xrange(256):
        l.append(i * 65536 + j * 256)
write_bmp(pm)
raise
"""

x = 400
half_x = x / 2
y = 600
half_y = y / 2
pm = []
for i in xrange(x):
    l = []
    pm.append(l)
    for j in xrange(y):
        l.append(0xFFFFFF)

def f(z, n):
    return z ** n - complex(1, 1)

def df(z, n):
    return n * z ** (n - 1)

def color(i, j):
    t = 2.0
    x = (i - half_x) / t
    y = (j - half_y) / t / 1.2
    n = 3
    """
    length = math.sqrt(x * x + y * y)
    if length != 0:
        degree = math.asin(y / length)
        if x < 0:
            if degree > 0:
                degree = math.pi - degree
            else:
                degree = - math.pi - degree
        try:
            degree += 1 / length ** 0.4
        except:
            pass
        if length < 1:
            length = 2 - length
        x = length * math.cos(degree)
        y = length * math.sin(degree)
    """
    z = complex(x, y)
    iter_num = 32
    for k in xrange(iter_num):
        if abs(f(z, n)) < 0.00001:
            break
        try:
            z = z - f(z, n) / df(z, n)
        except:
            break
    if k < iter_num / 4:
        r = 0xFF
        g = 0xFF - int(k / float(iter_num / 4 - 1) * 0xFF)
        return r * 65536 + g * 256
    else:
        r = 0xFF - int((iter_num - 1 - k) / float(iter_num - 1 - iter_num / 4) * 0xFF)
        g = 0xFF
        return r * 65536 + g * 256

for i in xrange(x):
    print i
    for j in xrange(y):
        pm[i][j] = color(i, j)

write_bmp(pm)
