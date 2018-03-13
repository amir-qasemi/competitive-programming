#include <iostream>
#include <string>

using namespace std;

int main()
{
    string a;
    bool dang = false;
    bool z;
    short num = 1;
	
	cin >> a;
	
	if(a[0] == '0'){
		z = true;
    } else{
		z = false;
	}
		
    for(int i = 1; i < a.size(); i++){
        if(z){
            if(a[i] == '0'){
                num++;
            } else{
                num = 1;
                z = false;
            }
        } else{
            if(a[i] == '1'){
                num++;
            } else{
                num = 1;
                z = true;
            }
        }
        
        if(num == 7){
            cout << "YES" << endl;
            return 0;
        }
    }
    
    cout << "NO" << endl;
    return 0;
    
}