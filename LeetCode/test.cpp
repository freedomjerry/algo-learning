#include<iostream>
using namespace std;
int main(){
    unsigned char *p1; 
    unsigned long *p2; 
    p1 = (unsigned char *)0x801000; 
    p2 = (unsigned long *)0x810000; 
    cout << p1 << endl;
    return 0;
}