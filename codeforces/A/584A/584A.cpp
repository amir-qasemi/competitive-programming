#include <iostream>
#include <string>

using namespace std;

int main(){

	long long n, t;
	
	cin >> n >> t;
	
	if(n == 1 && t == 10){
		cout << -1 << endl;
		return 0;
	}
	
	
	string num = to_string(t);
	if(t == 10){
		for(int i = 0 ; i < n - 2; i++){
				num += "0";
		}
	} else{
		
		for(int i = 0 ; i < n - 1; i++){
				num  += "0";
		}
	}
	cout << num << endl;


	return 0;
}