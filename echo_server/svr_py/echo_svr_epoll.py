import socket
import select

def main():
    l = socket.socket()
    l.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    l.bind(("localhost", 9999))
    l.listen(10)
    ep = select.epoll(1024)
    ep.register(l.fileno(), select.EPOLLIN)
    m = {}
    while True:
        events = ep.poll()
        for fd, _ in events:
            if fd == l.fileno():
                c, _ = l.accept()
                m[c.fileno()] = c
                ep.register(c.fileno(), select.EPOLLIN)
            else:
                c = m[fd]
                s = c.recv(10000)
                c.sendall(s)

if __name__ == "__main__":
    main()
