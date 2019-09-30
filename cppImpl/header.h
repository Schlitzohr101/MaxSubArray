#ifndef HEADER_H
#define HEADER_H
#include<random>
#include<time.h>
// #include<vector>
#include<math.h>
#include<iostream>
#include<map>
#include"Tester.h"

using namespace std;


typedef int (*intFuncPtr)();
typedef void (*voidFuncPtr)(Tester&);
//apparently this is a ptr to a function that returns an int and has no arguements

voidFuncPtr Choices(int);

void choice1(Tester&);
void choice2(Tester&);

#endif