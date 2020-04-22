#include <cstdio>
#include <algorithm>
#include <cstring>
#include <time.h>
using namespace std;

int main(){
    int cnt = 0;
    srand((unsigned)time(NULL));
    int ans = rand()%100 + 1;
    printf("%d %d\n",1,100);
    fflush(stdout);
    int num;
    while(scanf("%d",&num)!=EOF){
        ++ cnt;
        if(cnt > 100){
            exit(128);
        }
        if(num == ans){
            printf("%d\n",0);
            fflush(stdout);
            if(cnt <= 7){
                exit(228);
            }else if(cnt <= 20){
                exit(188);
            }else{
                exit(158);
            }
        }else if (num > ans){
            printf("%d\n",1);
            fflush(stdout);
        }else{
            printf("%d\n",-1);
            fflush(stdout);
        }
    }
    return 0;
}
