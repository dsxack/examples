import socket
import concurrent.futures
import multiprocessing

BUF_SIZE = 512


def handler(conn, address):
    conn.send(b"HTTP/1.1 200 OK\n")
    conn.send(b"Transfer-Encoding: chunked\n")
    conn.send(b"\r\n")

    while True:
        data = conn.recv(BUF_SIZE)
        if not data:
            break

        n = len(data)

        conn.send("{0:02x}\r\n".format(n).encode())
        conn.send(data)
        conn.send(b"\r\n")

        if n < BUF_SIZE:
            break

    conn.send(b"0\r\n\r\n")
    conn.close()


def main():
    sock = socket.socket()
    sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    sock.bind(('', 8000))
    sock.listen(100)

    print("Start listening 8000 port for connections")
    with concurrent.futures.ThreadPoolExecutor(max_workers=multiprocessing.cpu_count()) as executor:
        while True:
            conn, address = sock.accept()
            executor.submit(handler, conn, address)


if __name__ == '__main__':
    main()
