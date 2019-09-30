#include "header.h"

Tester::Tester() {Ar = vector<int>();}

Tester::Tester(vector<int> ar ) {
    Ar = ar;
}
 
Tester::Tester(int n) {
    srand(time(0));
    for (int i = 0; i < n; i++)
    {
        Ar.push_back(50 - rand()%101);
    }
}

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
// vector<int>& Tester::getAr() {
//     return Ar;
// }

void Tester::addInt(int x) {
    Ar.push_back(x);
}

int Tester::alg1() {
    int maxSum = 0;

    // cout << "here" << endl;
    for (int i = 0; i < Ar.size(); i++) {
        for (int j = i; j < Ar.size(); j++)
        {
            int thisSum = 0;
            for (int k = i; k <= j; k++)
            {
                thisSum += Ar[k];
            }
            if (thisSum > maxSum) {
                maxSum = thisSum;
            }
        }
    }
    return maxSum;
}

int Tester::alg2() {
    int maxSum = 0;
    for (int i = 0; i < Ar.size(); i++)
    {
        int thisSum = 0;
        for (int j = i; j < Ar.size();j++) {
            thisSum += Ar[j];
            if (thisSum > maxSum) {
                maxSum = thisSum;
            }
        }
    }
    return maxSum;
}

int Tester::alg3(int l, int r) {
    if (l == r) {
        return max(Ar[l],Ar[r]);
    }
    if (l+1 == r) {
        return max (Ar[l],max(Ar[r],Ar[l]+Ar[r]));
    }

    int m = (l + r)/2;

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
    cout << "Array to be tested:\n";
    for (int i = 0; i < Ar.size(); i++)
    {
        cout << Ar.at(i) << (i == Ar.size()-1 ? "\n": ", ");
    }
    clock_t start = clock();
    cout << "MSS found: " << alg1() << endl;
    float diff = float(clock() - start);
    cout << "Elapsed time: ";
    clocksToTime(diff);
    cout << endl << endl;

    return diff; 
}

float Tester::measure2() {
    cout << "Algorithm #2 is being tested" << endl;
    cout << "Array to be tested:\n";
    for (int i = 0; i < Ar.size(); i++)
    {
        cout << Ar.at(i) << (i == Ar.size()-1 ? "\n": ", ");
    }
    clock_t start = clock();
    cout << "MSS found: " << alg2() << endl;
    float diff = float(clock() - start);
    cout << "Elapsed time: ";
    clocksToTime(diff);
    cout << endl << endl;

    return diff; 
}

float Tester::measure3() {
    cout << "Algorithm #3 is being tested" << endl;
    cout << "Array to be tested:\n";
    for (int i = 0; i < Ar.size(); i++)
    {
        cout << Ar.at(i) << (i == Ar.size()-1 ? "\n": ", ");
    }
    int l = 0;
    int r = Ar.size()-1;
    clock_t start = clock();
    cout << "MSS found: " << alg3(l,r) << endl;
    float diff = float(clock() - start);
    cout << "Elapsed time: ";
    clocksToTime(diff);
    cout << endl << endl;

    return diff; 
}

float Tester::measure4() {
    cout << "Algorithm #4 is being tested" << endl;
    cout << "Array to be tested:\n";
    for (int i = 0; i < Ar.size(); i++)
    {
        cout << Ar.at(i) << (i == Ar.size()-1 ? "\n": ", ");
    }
    clock_t start = clock();
    cout << "MSS found: " << alg4() << endl;
    float diff = float(clock() - start);
    cout << "Elapsed time: ";
    clocksToTime(diff);
    cout << endl << endl;

    return diff; 
}