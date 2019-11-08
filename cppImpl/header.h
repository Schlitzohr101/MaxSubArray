//William Murray
//John Miner

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


typedef float (*floatFuncPtr)(Tester&);
typedef void (*voidFuncPtr)(Tester&);
//apparently this is a ptr to a function that returns an int and has no arguements

floatFuncPtr whichToMeasure(int);
voidFuncPtr Choices(int);

void choice1(Tester&);
void choice2(Tester&);
void choice3(Tester&);

float Measure1(Tester&);
float Measure2(Tester&);
float Measure3(Tester&);
float Measure4(Tester&);


float calculate1(int, int, float);
float calculate2(int, int, float);
float calculate3(int, int, float);
float calculate4(int, int, float);


void clocksToTime(float);

#endif