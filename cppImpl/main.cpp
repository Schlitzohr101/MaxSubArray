#include "header.h"

int main() {
    int choice;
    Tester controller = Tester();
    controller.Menu();
    cin >> choice;
    voidFuncPtr whichChoice = nullptr;
    while (choice != 4)
    {
        whichChoice = Choices(choice);
        whichChoice(controller);
        controller.Menu();
        cin >> choice;
    }
    cout << "Have a good day!\n";
}

voidFuncPtr Choices(int x) {
    map< const int,voidFuncPtr> m{
        {1,choice1},
        {2,choice2},
        {3,choice3},
    };
    return m[x];
}

floatFuncPtr whichToMeasure(int x) {
    map < const int, floatFuncPtr> m {
        {1,Measure1},
        {2,Measure2},
        {3,Measure3},
        {4,Measure4},
    };
    return m[x];
}

void choice1(Tester &T) {
    while ((getchar()) != '\n'); 
    cout << "Enter your Array as intergers followed by commas i.e.(1,2,3,...)\n:";
    string reply;
    getline(cin, reply);
    string delimiter = ",";
    size_t pos;
    while ((pos = reply.find(delimiter)) != string::npos)
    {
        T.addInt(stoi(reply.substr(0,pos)));
        reply.erase(0,pos + delimiter.length());
    }
    T.addInt(stoi(reply));
    cout << "\nComputed Max Sub-Array Sum -\n";
    cout << "Algorithm #1: " << T.alg1() << endl;
    cout << "Algorithm #2: " << T.alg2() << endl;
    int l = 0;
    int r = T.size()-1;
    cout << "Algorithm #3: " << T.alg3(l,r) << endl;
    cout << "Algorithm #4: " << T.alg4() << endl << endl;
    T.clear();
}

void choice2(Tester& T) {
    cout << "Enter a array size\n:";
    int size;
    cin >> size;
    getchar();
    T = Tester(size);

    cout << "Enter which Algorithms to test i.e.(1234)\n:";
    string reply;
    getline(cin,reply);
    if (reply.find("1") != string::npos) {
        T.measure1();
    }
    if (reply.find("2") != string::npos) {
        T.measure2();
    }
    if (reply.find("3") != string::npos) {
        T.measure3();
    }
    if (reply.find("4") != string::npos) {
        T.measure4();
    }
    
}

void choice3(Tester& T) {
    cout << "Which Algorithm would you like to test\n:";
    int alg;
    cin >> alg;
    getchar();
    // floatFuncPtr algorithm = whichToMeasure(T,alg);
    cout << "Enter the baseline array size\n:";
    int base;
    cin >> base;
    getchar();
    cout << "Enter a array size to estimate\n:";
    int est;
    cin >> est;
    getchar();
    T = Tester(base);
    floatFuncPtr algorithm = whichToMeasure(alg);
    float baseRun = algorithm(T);

    float estimated;
    switch (alg)
    {
    case 1: estimated = calculate1(base,est,baseRun);
        break;
    case 2: estimated = calculate2(base,est,baseRun);
        break;
    case 3: estimated = calculate3(base,est,baseRun);
        break;
    case 4: estimated = calculate4(base,est,baseRun);
        break;
    default: estimated = 0;
        break;
    }

    cout << "The expected runtime for Algorithm #" << alg << " is ";
    clocksToTime(estimated);
    cout << endl << endl;

    T = Tester(est);
    algorithm = whichToMeasure(alg);
    algorithm(T);
}

float Measure1(Tester& T) {
    return T.measure1();
}
float Measure2(Tester& T) {
    return T.measure2();
}
float Measure3(Tester& T) {
    return T.measure3();
}
float Measure4(Tester& T) {
    return T.measure4();
}



float calculate1(int m, int n, float runtime) {
    return runtime * pow(float(n/m),3.0);
}

float calculate2(int m, int n, float runtime) {
    return runtime * pow(float(n/m),2.0);
}

float calculate3(int m, int n, float runtime) {
    float C = runtime/(m * log(m));
    return C * n * log(n);
}

float calculate4(int m, int n, float runtime) {
    return runtime * float(n/m);
}

void clocksToTime(float diff) {
    if (diff/pow(10.0,9.0) > 1) {
        printf("%.3f (s)",diff/pow(10.0,9.0));
    } else if (diff/pow(10.0,6.0) > 1) {
        printf("%.3f (ms)",diff/pow(10.0,6.0));
    } else if (diff/pow(10.0,3.0) > 1) {
        printf("%.3f (µs)",diff/pow(10.0,3.0));
    } else {
        printf("%d (ns)",int(diff));
    }
}