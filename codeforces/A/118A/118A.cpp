#include <iostream>
#include <string>

using namespace std;

int main()
{
    string a;
    
    cin >> a;
    for(int i = 0; i < a.size(); i++){
       if(!(a[i] == 'a' || a[i] == 'A' || a[i] == 'o' || a[i] == 'O' || a[i] == 'y' || a[i] == 'Y' || a[i] == 'e' || a[i] == 'E' || a[i] == 'u' || a[i] == 'U' ||  a[i] == 'i' || a[i] == 'I')){
           cout << ".";
           if(a[i] >= 65 && a[i] <= 90){
               cout << ((char)(a[i] + 32));
           } else{
               cout << a[i];
           }
       }
    }
   cout << endl;
   return 0;
}