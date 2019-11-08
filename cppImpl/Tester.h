//William Murray
//John Miner
#ifndef TESTER_H
#define TESTER_H
#include<vector>

using namespace std;
/*
*Class 
*Used to hold the array and all methods relating to
*Testing the algorithms
*/

class Tester {
    private:
        vector<int> Ar;

    public:
        //Constructor and Destructors
        Tester();
        Tester(vector<int>);
        Tester(int);
        ~Tester();
        //Vector wrappers
        int at(int);
        void addInt(int);
        void clear();
        int size();
        //Algorithms
        int alg1();
        int alg2();
        int alg3(int, int);
        int alg3middle(int, int, int);
        int alg4();
        //Time measures for each method
        float measure1();
        float measure2();
        float measure3();
        float measure4();
        //Print menu
        void Menu();

};

#endif