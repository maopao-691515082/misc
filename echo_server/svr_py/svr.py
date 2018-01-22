import socket
import thread

def work(c):
    #c.setsockopt(socket.IPPROTO_TCP, socket.TCP_NODELAY, 1)
    while True:
        s = c.recv(10000)
        c.sendall(s)

def main():
    l = socket.socket()
    l.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    l.bind(("localhost", 9999))
    l.listen(10)
    while True:
        c, _ = l.accept()
        thread.start_new_thread(work, (c,))
        c = None

if __name__ == "__main__":
    main()
