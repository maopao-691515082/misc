#include <iostream>
#include <cstdlib>
#include <cstring>

#include <sys/socket.h>
#include <arpa/inet.h>
#include <netinet/tcp.h>
#include <sys/epoll.h>

using namespace std;

void exit_prog(const char *msg)
{
    cout << msg << endl;
    exit(1);
}

int main()
{
    struct sockaddr_in addr;
    memset(&addr, 0, sizeof(addr));
    addr.sin_family = AF_INET;
    addr.sin_port = htons(9999);
    addr.sin_addr.s_addr = INADDR_ANY;

    int l = socket(AF_INET, SOCK_STREAM, 0);
    if (l == -1)
    {
        exit_prog("create listen sock failed");
    }

    int reuse_addr = 1;
    if (setsockopt(l, SOL_SOCKET, SO_REUSEADDR, &reuse_addr, sizeof(reuse_addr)) == -1)
    {
        exit_prog("set reuse addr failed");
    }

    if (bind(l, (struct sockaddr *)&addr, sizeof(addr)) == -1)
    {
        exit_prog("bind failed");
    }

    if (listen(l, 10) == -1)
    {
        exit_prog("listen failed");
    }

    int ep = epoll_create(1024);
    if (ep == -1)
    {
        exit_prog("epoll create failed");
    }

#define MAX_EVENTS 100
    struct epoll_event ev, events[MAX_EVENTS];

    ev.events = EPOLLIN;
    ev.data.fd = l;
    if (epoll_ctl(ep, EPOLL_CTL_ADD, l, &ev) == -1) {
        exit_prog("epoll_ctl failed");
    }

    for (;;)
    {
        int nfds = epoll_wait(ep, events, MAX_EVENTS, -1);
        if (nfds == -1) {
            exit_prog("epoll_wait fail");
        }

        for (int n = 0; n < nfds; ++ n)
        {
            if (events[n].data.fd == l)
            {
                int c = accept(l, NULL, NULL);
                if (c == -1)
                {
                    exit_prog("accept failed");
                }
                ev.events = EPOLLIN;
                ev.data.fd = c;
                if (epoll_ctl(ep, EPOLL_CTL_ADD, c, &ev) == -1)
                {
                    exit_prog("epoll_ctl failed");
                }
            }
            else
            {
                int c = events[n].data.fd;
                char buf[10000];
                ssize_t recv_len = recv(c, buf, sizeof(buf), 0);
                if (recv_len <= 0)
                {
                    exit_prog("recv error");
                }
                ssize_t sent_len = send(c, buf, (size_t)recv_len, 0);
                if (sent_len != recv_len)
                {
                    exit_prog("send error");
                }
            }
        }
    }
}
