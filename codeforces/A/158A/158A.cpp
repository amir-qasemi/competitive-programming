#include <iostream>


using namespace std;


int main(){
	int n , k;
	
	cin >> n >> k;
	
	
	int last;
	int num = 0;
	
	bool isIs = true;
	cin >> last;
	
	if(last > 0){
		num++;
	}
	
	
	int temp;
	for(int i = 1 ; i < n; i++){
		cin >> temp;
		 
		 
		if(isIs){
			if(num < k && temp != 0){
				num++;
			} else if(num >= k && temp == last && temp != 0 ){
				num++;
			} else if(num >= k && temp != last && temp != 0){
				isIs = false;
			}
		}
		
		
		last = temp;
	}

    cout << num << endl;
    
    
    
	return 0;
}