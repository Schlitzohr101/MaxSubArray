#ifndef TESTER_H
#define TESTER_H
#include<vector>

using namespace std;

class Tester {
    private:
        vector<int> Ar;


    public:
        Tester();
        Tester(vector<int>);
        ~Tester();
        Tester(int);
        void Menu();
        int alg1();
        int alg2();
        int alg3(int, int);
        int alg3middle(int, int, int);
        int alg4();
        vector<int> &getAr();
        void addInt(int);
        void clear();
        int size();

};

#endif