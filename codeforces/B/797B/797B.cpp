#include <iostream>
#include <limits>
#include <cstdlib>

using namespace std;

int main()
{
   int n;
   cin >> n;
   
   int *a = new int[n];
   int i;
   int minOddPos,minOddNeg;
   minOddPos = numeric_limits<int>::max();
   minOddNeg =  -1 * minOddPos;
   int sum = 0;
   
   for(i = 0; i < n; i++){
       cin >> a[i];
   }
   
   for(i = 0; i < n; i++){
       if(a[i] > 0){
           sum += a[i];
           
           if(a[i] & 1){
               if(a[i] < minOddPos){
                   minOddPos = a[i];
               }
           }
           
       } else if(a[i] < 0){
           if(a[i] & 1){
               if(a[i] > minOddNeg){
                   minOddNeg = a[i];
               }
           }
           
       }
       
      
   }
   
   if(!(sum & 1)){
       if(abs(minOddNeg) < minOddPos){
           sum += minOddNeg;
       } else{
           sum -= minOddPos;
       }
   }
   
   
   cout << sum << endl;
   
   return 0;
}