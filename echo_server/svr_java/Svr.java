import java.net.*;
import java.io.*;

class Worker extends Thread
{
    Worker(Socket c) throws Exception
    {
        this.c = c;
	this.c.setTcpNoDelay(true);
        this.i = c.getInputStream();
        this.o = c.getOutputStream();
        this.buf = new byte[10000];
    }

    public void run()
    {
        try
        {
            for (;;)
            {
                int recvLen = i.read(this.buf);
                if (recvLen <= 0)
                {
                    throw new Exception("read failed");
                }
                o.write(this.buf, 0, recvLen);
            }
        }
        catch (Exception e)
        {
            return;
        }
    }

    Socket c;
    InputStream i;
    OutputStream o;
    byte[] buf;
}

public class Svr
{
    public static void main(String[] args) throws Exception
    {
        ServerSocket s = new ServerSocket(9999);
        s.setReuseAddress​(true);
        for (;;)
        {
            Socket c = s.accept();
            new Worker(c).start();
        }
    }
}
