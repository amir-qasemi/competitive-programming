#include <iostream>
#include <string>

using namespace std;

int main()
{
    string s;
    int n , t;
    
    cin >> n >> t;
    
    cin >> s;

    for(int j = 0; j < t; j++){
        for(int i = 0; i < n; ){
            if(i != n -1){
                if(s[i] == 'B' && s[i + 1] == 'G'){
                    s[i] = 'G';
                    s[i + 1] = 'B';
                    i+= 2;
                    continue;
                } 
            }
            i++;
        }
    }
    
    cout << s << endl;

    return 0;
}