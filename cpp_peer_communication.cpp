#include <iostream>
#include <string>
#include <sys/socket.h>
#include <netinet/in.h>

class PeerCommunication {
private:
    int sockfd;
    int port;
public:
    PeerCommunication(int p) : port(p) {}

    bool init() {
        sockfd = socket(AF_INET, SOCK_STREAM, 0);
        return sockfd >= 0;
    }

    bool sendMessage(int fd, std::string msg) {
        send(fd, msg.c_str(), msg.size(), 0);
        return true;
    }

    std::string recvMessage(int fd) {
        char buf[1024] = {0};
        recv(fd, buf, 1024, 0);
        return std::string(buf);
    }
};
