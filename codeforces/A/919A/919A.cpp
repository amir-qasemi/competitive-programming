// Example program
#include <iostream>
#include <limits>

using namespace std;
int main()
{
    double n, m;
    cin >> n >> m;
    double min = numeric_limits<double>::max();
    double a, b;
    for(int i =0; i< n; i++){
        cin >> a >> b;
        if((double)((double)a / (double)b) < (double)min){
            min = (double)((double)a / (double)b);
        }
    }
    
    cout.precision(10);
    cout << (double)((double)min * (double)m) << endl;
    
    return 0;
}
