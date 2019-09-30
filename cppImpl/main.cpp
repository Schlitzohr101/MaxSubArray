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
    
}

voidFuncPtr Choices(int x) {
    map< const int,voidFuncPtr> m{
        {1,choice1},
        {2,choice2},
    };
    return m[x];
}

void choice1(Tester &T) {
    while ((getchar()) != '\n'); 
    cout << "Enter your Array as intergers followed by commas i.e.(1,2,3,...)\n:";
    string reply;
    getline(cin, reply);
    vector<string> array = vector<string>();
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
    
}