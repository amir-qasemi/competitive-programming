#include <iostream>
#include <string>

using namespace std;

int main()
{
    string a, b;
    short aN, bN;
    cin >> a >> b;
    bool equal = true, fGrt = false; 
    
    for(int i = 0 ; i < a.size(); i++){
        aN = ((int)a[i]) % 32;
        bN = ((int)b[i]) % 32;
        if( aN > bN ){
            equal = false;
            fGrt = true;
            break;
        } else if(aN < bN){
            equal = false;
            break;
        }
    }
    
    if(equal == true){
        cout << 0 << endl;
    } else if(fGrt == true){
        cout << 1 << endl;
    } else{
        cout << -1 << endl;
    }
    
    
    return 0;
    
}