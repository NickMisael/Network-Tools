#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/socket.h>
#include <netdb.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <time.h>

#define PORT 53

typedef struct
{
    char nome[100];
    char paddr[100];
    char saddr[100];
    double time;
} dns;

dns servers[] = { {"OpenDNS","208.67.222.222","208.67.220.220"}, {"Cloudflare","1.1.1.1","1.0.0.1"}, {"Google","8.8.8.8","8.8.4.4"}, 
                  {"Norton","199.85.126.10","199.85.127.10"}, {"Verisign", "64.6.64.6","64.6.65.6"}, {"NuSEC","8.26.56.26","8.20.247.20"},
                  {"Quad9","9.9.9.9","149.112.112.112"}, {"Neustar","156.154.70.5","156.154.71.5"}, {"SafeDNS","195.46.39.39","195.46.39.40"}, 
                  {"Yandex","77.88.8.8","77.88.8.1"}, {"AdGuard DNS","94.140.14.14","94.140.15.15"}, {"Alternate DNS","76.76.19.19","76.223.122.150"},
                  {"CleanBrowsing","185.228.168.9","185.228.169.9"}, {"Ultra DNS","64.6.64.6","64.6.65.6"}, {"Oracle Dyn DNS","216.146.35.35","216.146.36.36"},
                  {"Level3 DNS", "209.244.0.3","209.244.0.4"}, {"Comodo Secure DNS","8.26.56.26","8.20.247.20"}, {"Freenom World DNS", "80.80.80.80", "80.80.80.81"}};

void help(char * arg);
void verify();
void change();
void bubblesort(dns * server);
void showDnsList();

int main(int argc, char **argv) {
    
    if(argc != 2){
        help(argv[0]);
        exit(-1);
    }

    if((strcmp(argv[1],"-C")) == 0){
        verify();
    } else if((strcmp(argv[1],"-S")) == 0){
        change();
    }


    return 0;
}

void help(char * arg){
    printf("Example of use:\n"
           "%s [option of use]\n"
           "-C - Check the best DNS option\n"
           "-S - Switch to best DNS option\n", arg);
}
void verify(){
    for(int j = 0; j < 10; j++){
        for(int i = 0; i < 18; i++){
            clock_t t;
            int fd;
            struct in_addr addr;
            struct sockaddr_in sock;
            struct protoent *proto = getprotobyname("tcp");

            memset(&addr, '\0', sizeof(addr));
            memset(&sock, '\0', sizeof(sock));

            if((fd = socket(AF_INET, SOCK_STREAM, proto->p_proto)) == -1){
                perror("socket");
                exit(-1);        
            }

            if(!inet_aton(servers[i].paddr, &addr)){
                fprintf(stderr, "Invalid Address!\n");
                close(fd);
                exit(-1);
            }

            sock.sin_family = AF_INET;
            sock.sin_addr   = addr;
            sock.sin_port   = htons(PORT);
            
            t = clock();
            if(connect(fd, (struct sockaddr *) &sock, sizeof(sock)) == -1){
                t = clock() - t;
                servers[i].time = (((double)t)/((CLOCKS_PER_SEC)/1000));   
            } else {
                printf("Server %s - esta inacessivel\n", servers[i].nome);
            }
            close(fd);
        }
    }
    bubblesort(servers);
    showDnsList();
}

void change(){
    printf("teste"); 
}

void bubblesort(dns * server) {
    int k,j;
    dns aux;
    for(k = 18 - 1; k > 0; k--){

        for(j = 0; j < k; j++){
            if(server[j].time > server[j+1].time){
                aux = server[j];
                server[j] = server[j+1];
                server[j+1] = aux;
            }
        }
    }
}

void showDnsList(){
    printf("\a");
	printf("\t _________________________________________\n");
	printf("\t|                                         |\n");
	printf("\t|       Host | DNS Server | Time ms       |\n");
	printf("\t|_________________________________________|\n");
	printf("\t|                                         |\n");
	for(int i = 0; i < 18; i++){
		printf("\t|-----------------------------------------|\n");
        printf("\t| %s \n", servers[i].nome);
		printf("\t| P: %s \n", servers[i].paddr);
		printf("\t| S: %s \n", servers[i].saddr);
		printf("\t| %.2f milisegundos \n", servers[i].time);
		if(i == 18-1){
			printf("\t|_________________________________________|\n");
		}
	}
}
