package org.dsxack.http_echo_over_tcp;

import java.io.EOFException;
import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.net.ServerSocket;
import java.net.Socket;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class Main {

    public static void main(String[] args) throws IOException {
        int cores = Runtime.getRuntime().availableProcessors();

        ExecutorService executor = Executors.newFixedThreadPool(cores);

        ServerSocket listener = new ServerSocket(8000);
        System.out.print("Listen :8000 port for new connections\n");


        while (true) {
            Socket socket;
            try {
                socket = listener.accept();
            } catch (IOException e) {
                break;
            }

            executor.execute(new Handler(socket));
        }
    }

    static class Handler implements Runnable {
        static private int BUF_SIZE = 512;

        Socket socket;

        Handler(Socket socket) {
            this.socket = socket;
        }

        @Override
        public void run() {
            try {
                handleSocket();
            } catch (IOException e) {
                e.printStackTrace();
            }
            try {
                socket.close();
            } catch (IOException e) {
                e.printStackTrace();
            }
        }

        private void handleSocket() throws IOException {
            InputStream input = socket.getInputStream();
            OutputStream output = socket.getOutputStream();

            output.write("HTTP/1.1 200 OK\n".getBytes());
            output.write("Transfer-Encoding: chunked\n".getBytes());
            output.write("\r\n".getBytes());

            byte[] buf = new byte[Handler.BUF_SIZE];

            while (true) {
                int n;

                try {
                    n = input.read(buf);
                } catch (EOFException e) {
                    break;
                }

                if (n == -1) {
                    break;
                }

                output.write(String.format("%X\r\n", n).getBytes());
                output.write(buf, 0, n);
                output.write("\r\n".getBytes());

                if (n < Handler.BUF_SIZE) {
                    break;
                }
            }

            output.write("0\r\n\r\n".getBytes());

            input.close();
            output.close();
        }
    }
}
