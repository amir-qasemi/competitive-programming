// Example program
#include <iostream>



using namespace std;

int main()
{
    int n, m;
    cin >> n >> m;
    int *a = new int[m];
    long long time = 0;
    int currentPos = 1;
    for(int i = 0; i < m; i++){
        cin >>  a[i];
    }
    
    
    for(int i = 0; i < m; i++){
     
        if( a[i] < currentPos){
            time += n - currentPos + a[i];
            currentPos = a[i];
        } else{
            time += a[i] - currentPos;
            currentPos = a[i];
        }
    }

	
    cout  << fixed <<time << endl;
    return 0;
}