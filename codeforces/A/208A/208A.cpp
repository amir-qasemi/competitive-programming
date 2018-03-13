// Example program
#include <iostream>
#include <string>


using namespace std;

int main()
{
    string s, t, b = " "; 
     
     cin >> s;
     
    for(int i = 0 ; i < s.length(); i++){
        if(s[i] == 'W' && s[i + 1] == 'U' && s[i + 2] == 'B'){
            i += 2;
            t.append(b);
            continue;
        } else{
              t += s[i];
        }
    }
    int begin = 0, end = t.length();
    for(int i = 0; i < t.length(); i++){
        if(t[i] != ' '){
            begin = i;
            break;
        }
    }
    
    
    for(int i = t.length() - 1; i >= 0; i--){
        if(t[i] != ' '){
            end = i;
            break;
        }
    }
    
    
    cout << t.substr(begin, end + 1) << endl;
    return 0;
}