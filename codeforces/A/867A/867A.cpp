#include <iostream>
#include <string>

using namespace std;

int main(){
	string a;
	
	int s = 0 , f = 0;
	int num = 0;
	cin >> num;
	cin >> a;
	
	for(int i = 1; i < a.length() ; i++){
		if( a[i - 1] != a[i]){
			if(a[i] == 'S'){
				s++;
			} else {
				f++;
			}
		}
		
	}

	
	if(f > s){
	    cout << "YES" << endl; 
	} else {
	    cout << "NO" << endl;
	}
	
	return 0;
}