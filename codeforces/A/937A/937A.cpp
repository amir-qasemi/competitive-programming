#include <iostream>

using namespace std;

int main()
{
    int n;
    cin >> n;
    int a[601];
    int temp;
    for(int i = 1; i < 601; i++){
        a[i]  = 0;
    }
    for(int i = 0; i <  n; i++){
        cin >> temp;
        a[temp] += 1;
    }
    
    int num = 0;
    for(int i = 1; i < 601; i++){
        if(a[i] > 0){
            num++;
        }
    }
    
    cout << num << endl;
    
   return 0;
}