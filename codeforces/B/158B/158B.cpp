#include <iostream>

using namespace std;

int main()
{
    int n;
    int *g;
    
    cin >> n;
    g = new int[n];
    
    for(int i = 0; i < n; i++){
        cin >> g[i];
    }
    
    int o, tw, th , f;
    
    o = tw = th = f = 0;
    
    
    for(int i = 0; i < n; i++){
        if(g[i] == 1){
            o++;
        } else if(g[i] == 2){
            tw++;
        } else if(g[i] == 3){
            th++;
        } else if(g[i] == 4){
            f++;
        }
    
    }
    
    int num = 0;
    num += tw / 2;

    if(tw % 2 == 0 ){
        tw = 0;
    } else {
        tw = 1;
    }
        
    num += f;
    
    if(th < o){
        num += th;
        o -= th;
        th = 0;
            
            
        int temp = o / 4;
        num += temp;
        o = o - (temp * 4);
        
        
        if(o % 4 == 0){
            
        } else if(o % 4 == 1){
            num++;
			tw = 0;
        }  else if(o % 4 == 2){
            tw = 0;
			num++;
        }  else if(o % 4 == 3){
            num++;
        }
        
    } else if (o < th){
        num += th;
        th = o = 0;
        
    } else{
        num += o;
    }
    
    if(tw == 1){
        num++;
    }
    

    cout << num << endl;
    
   return 0;
}