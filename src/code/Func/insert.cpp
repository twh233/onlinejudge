#include <cstdio>
#include <algorithm>
using namespace std;

__CODE__

int main(){
    int a,b;
    while(scanf("%d%d",&a,&b)!=EOF){
        int c = add(a,b);
        printf("%d\n",c);
    }
    return 0;
}
