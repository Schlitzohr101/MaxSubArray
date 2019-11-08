//William Murray
//John Miner
#include "header.h"

//Builds blank array
Tester::Tester() {Ar = vector<int>();}

//Builds with argument array
Tester::Tester(vector<int> ar ) {
    Ar = ar;
}

//Builds array with random values between -50-50
//of size n
Tester::Tester(int n) {
    srand(time(0));
    for (int i = 0; i < n; i++)
    {
        Ar.push_back(50 - rand()%101);
    }
}

//clears memory
Tester::~Tester() {
    Ar.clear();
}

void Tester::Menu() {
    cout << "MENU\n" 
         << "1) Enter your own array to test the algorithms\n" 
         << "2) Enter an a array size to be used to test select algorithms\n"
         << "3) Select a algorithmn to estimate runtimes based on different array sizes\n"
         << "4) Quit\n:";
}


void Tester::clear() {
    Ar.clear();
}

int Tester::size() {
    return Ar.size();
}

int Tester::at(int x) {
    return Ar.at(x);
}

void Tester::addInt(int x) {
    Ar.push_back(x);
}

/*
Algorithm 1
thriple nested loops
runtime of O(n^3)
*/
int Tester::alg1() {
    int maxSum = 0;

    // cout << "here" << endl;
    //sets the primary index
    for (int i = 0; i < Ar.size(); i++) {
        //each subsequent loop is then starts at
        //whatever the primary index is at 
        //loop is the end index
        for (int j = i; j < Ar.size(); j++)
        {
            //intialize the temperary sum
            int thisSum = 0;
            //loop is the adding index
            for (int k = i; k <= j; k++)
            {
                //each value is added to the temporary sum
                thisSum += Ar[k];
            }
            //if its bigger than the max sum then max = temp
            if (thisSum > maxSum) {
                maxSum = thisSum;
            }
        }
    }
    return maxSum;
}

/*
Algorithm #2
double nested
runtime of O(n^2)
*/
int Tester::alg2() {
    int maxSum = 0;
    //sets primary index
    for (int i = 0; i < Ar.size(); i++)
    {
        int thisSum = 0;
        //sets the adding index
        for (int j = i; j < Ar.size();j++) {
            thisSum += Ar[j];
            if (thisSum > maxSum) {
                maxSum = thisSum;
            }
        }
    }
    return maxSum;
}

/*
Algorithm #3
recursive summation
runtime of O(nlog(n))
*/
int Tester::alg3(int l, int r) {
    //recursive checks
    //if the argument indexes are equal to each other
    if (l == r) {
        return Ar[l];
    }
    //if the two indexs are side by side
    if (l+1 == r) {
        return max (Ar[l],max(Ar[r],Ar[l]+Ar[r]));
    }

    //calculate the middle index (integer math)
    int m = (l + r)/2;

    //recursive calls
    int LSS = alg3(l, m);
    int RSS = alg3(m+1,r);
    int MSS = alg3middle(l,m,r);

    return max(LSS,max(RSS,MSS));
}

int Tester::alg3middle(int l, int m, int r) {
    int maxLeft = -(2000000000);
    int sum = 0;

    for (int i = m; i >= l; i--)
    {
        sum += Ar.at(i);
        if (sum > maxLeft) {
            maxLeft = sum;
        }
    }

    int maxRight = -(2000000000);
    sum = 0;
    
    for (int i = m+1; i <= r; i++)
    {
        sum += Ar.at(i);
        if (sum > maxRight) {
            maxRight = sum;
        }
    }
    return maxRight + maxLeft;
    
}

int Tester::alg4() {
    int maxSum = 0;
    int thisSum = 0;
    for (int i = 0; i < Ar.size(); i++)
    {
        thisSum += Ar[i];
        if (thisSum > maxSum) {
            maxSum = thisSum;
        } else if ( thisSum < 0) {
            thisSum = 0;
        }
    }
    return maxSum;
}

float Tester::measure1() {
    cout << "Algorithm #1 is being tested" << endl;
    // cout << "Array to be tested:\n";
    // for (int i = 0; i < Ar.size(); i++)
    // {
    //     cout << Ar.at(i) << (i == Ar.size()-1 ? "\n": ", ");
    // }

    float totaldiff = 0;
    clock_t start = clock();
    cout << "MSS found: " << alg1() << endl;
    totaldiff += float(clock() - start);
    for (int i = 0; i < 9; i++)
    {
        clock_t start = clock();
        alg1();
        totaldiff += float(clock() - start);
    }
    

    cout << "Elapsed time: ";
    clocksToTime(totaldiff/10.0);
    cout << endl << endl;

    return totaldiff/10.0; 
}

float Tester::measure2() {
    cout << "Algorithm #2 is being tested" << endl;
    // cout << "Array to be tested:\n";
    // for (int i = 0; i < Ar.size(); i++)
    // {
    //     cout << Ar.at(i) << (i == Ar.size()-1 ? "\n": ", ");
    // }
    float totaldiff = 0;
    for (int i = 0; i < 9; i++) {
        clock_t start = clock();
         alg2();
        totaldiff += float(clock() - start);
    }
    clock_t start = clock();
    cout << "MSS found: " << alg2() << endl;
    totaldiff += float(clock() - start);
    cout << "Elapsed time: ";
    clocksToTime(totaldiff/10);
    cout << endl << endl;

    return totaldiff/10; 
}

float Tester::measure3() {
    cout << "Algorithm #3 is being tested" << endl;
    // cout << "Array to be tested:\n";
    // for (int i = 0; i < Ar.size(); i++)
    // {
    //     cout << Ar.at(i) << (i == Ar.size()-1 ? "\n": ", ");
    // }
    int l = 0;
    int r = Ar.size()-1;

    float totaldiff = 0;
    for (int i = 0; i < 9; i++) {
        clock_t start = clock();
         alg3(l,r);
        totaldiff += float(clock() - start);
    }

    clock_t start = clock();
    cout << "MSS found: " << alg3(l,r) << endl;
    totaldiff += float(clock() - start);
    cout << "Elapsed time: ";
    clocksToTime(totaldiff/10.0);
    cout << endl << endl;

    return totaldiff/10.0; 
}

float Tester::measure4() {
    cout << "Algorithm #4 is being tested" << endl;
    // cout << "Array to be tested:\n";
    // for (int i = 0; i < Ar.size(); i++)
    // {
    //     cout << Ar.at(i) << (i == Ar.size()-1 ? "\n": ", ");
    // }

    float totaldiff = 0;
    for (int i = 0; i < 9; i++) {
        clock_t start = clock();
        alg4();
        totaldiff += float(clock() - start);
    }
    clock_t start = clock();
    cout << "MSS found: " << alg4() << endl;
    totaldiff += float(clock() - start);
    cout << "Elapsed time: ";
    clocksToTime(totaldiff/10.0);
    cout << endl << endl;

    return totaldiff/10.0; 
}