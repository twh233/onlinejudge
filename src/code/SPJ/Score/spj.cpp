/*
spj���ܣ��ж��û���������ֺͱ�׼�������Ƿ񳬹�0.0001
*/
#include<stdio.h>
#include<stdlib.h>
#include<math.h>
#include<string.h>
#include<list>
#include<algorithm>
#include<random>
#include<string>
#include<time.h>
using namespace std;
#define AC 0
#define WA 4
void DONE(int st)
{
    exit(st);
}
/**
 * ���ص÷�
 * ����Ϊ0~100������
 */
void DONE_SC(int score)
{
    exit(0x80 + score);
}
string s1,s2;
char s[1000105];
void fiter(char s[])
{
    int len=strlen(s);
    while(len)
    {
        if(s[len-1]<=' ')
        {
            len--;
            s[len]=0;
        }
        else
        {
            return ;
        }
    }
}
int main(int argc,char **args)
{
    //  FILE *f_in = fopen(args[1],"r");//��׼�����ļ�
    FILE *f_out = fopen(args[2],"r");//��׼����ļ�
    FILE *f_user = fopen(args[3],"r");//�û�����ļ�
//   FILE *f_code = fopen(args[4],"r");//�û������ļ�
//    FILE *f_out = fopen("wine.out","r");//��׼����ļ�
//    FILE *f_user = fopen("2.out","r");//�û�����ļ�
    s1="";
    s2="";
    int len;
    while(fgets(s,1000000,f_out))
    {
        fiter(s);
        s1+=s;
    }
    while(fgets(s,1000000,f_user))
    {
        fiter(s);
        s2+=s;
    }
    if(s1==s2)
        DONE_SC(100);
    else
        DONE_SC(0);
}
