#include<stdio.h>
#include<stdlib.h>
#include<math.h>
int main(int argc,char **args)
{
    FILE *f_in = fopen(args[1],"r");
    FILE *f_out = fopen(args[2],"r");
    FILE *f_user = fopen(args[3],"r");
    int t;
    double n1,n2;
    while(fscanf(f_out,"%lf",&n1)!=EOF)
    {
        if(fscanf(f_user,"%lf",&n2)!=1) exit(4);
        if(fabs(n1-n2)>=0.0001) exit(4);
    }
    if(fscanf(f_user,"%lf",&n2)!=EOF) exit(4);
	return 0;
}
