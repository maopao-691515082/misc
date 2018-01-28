#include <iostream>
#include <cstdlib>
#include <cstring>

#include <sys/socket.h>
#include <arpa/inet.h>
#include <netinet/tcp.h>
#include <pthread.h>

using namespace std;

void exit_prog(const char *msg)
{
    cout << msg << endl;
    exit(1);
}

void *work(void *arg)
{
    int c = *(int *)arg;

    int no_delay = 1;
    if (setsockopt(c, IPPROTO_TCP, TCP_NODELAY, &no_delay, sizeof(no_delay)) == -1)
    {
        exit_prog("set no delay failed");
    }

    for (;;)
    {
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
    return NULL;
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

    for (;;)
    {
        int c = accept(l, NULL, NULL);
        if (c == -1)
        {
            exit_prog("accept failed");
        }
        pthread_t tid;
        if (pthread_create(&tid, NULL, work, (void *)new int(c)) != 0)
        {
            exit_prog("pthread_create failed");
        }
    }
}
