#include <iostream>

using namespace std;

int main()
{
    int n;
    cin >> n;
    int * a = new int[n];
    
    for(int i = 0 ; i < n; i++){
        cin >> a[i];
    }
    int sum = 0;
    int index = n;
    int best = 0;
    for(int i = n - 1; i >=0; i--){
        sum += a[i];
        if(sum < best){
            best = sum;
            index = i;
        }
    }
    
    
    int b = 0, c = 0;
    for(int i = 0; i < index; i++){
        b += a[i];
    }
    
    for(int i = index; i < n; i++){
        c += a[i];
    }
    
    int first = b - c;



	 sum = 0;
    index = -1;
    best = 0;
    for(int i = 0; i < n; i++){
        sum += a[i];
        if(sum < best){
            best = sum;
            index = i;
        }
    }
    
    
    b = 0, c = 0;
    for(int i = 0; i <= index; i++){
        b += a[i];
    }
    
    for(int i = index + 1; i < n; i++){
        c += a[i];
    }
	
	
	if(c - b > first){
		cout << c - b << endl;
	} else{
		cout << first << endl;
	}
   return 0;
}