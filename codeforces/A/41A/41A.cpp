// Example program
#include <iostream>
#include <string>


using namespace std;

int main()
{
    string s, t;
    
    cin >> s >> t;
    int size = s.length();
    if(size != t.length()){
        cout << "NO" << endl;
        return 0;
    } else{
        for(int  i = 0; i < size; i++){
            if(s[i] != t[size - i - 1]){
                cout << "NO" << endl;
        return 0;
            }
            
        }
    }
    
        cout << "YES" << endl;
    
    return 0;
}