
#include <SFE_BMP180.h>
#include <Wire.h>


// BMP180
SFE_BMP180 pressure;
#define ALTITUDE 1655.0


// motor control outputs
const int motorRear1 = 2;  // 12 o clock
const int motorRear2 = 3; // 4 o clock
const int motorRear3 = 4; // 7 o clock
const int motorFront1 = 5; // 2 o clock
const int motorFront2 = 6; // 6 o clock
const int motorFront3 = 7; // 10 o clock

// motor control inputs
const int forward = 10;
const int backward = 9;
const int left = 8;
const int right = 11;

int buttonStateForward = 0;
int buttonStateBackward = 0;
int buttonStateLeft = 0;
int buttonStateRight = 0;
//int buttonStateUp = 0;
//int buttonStateDown = 0;


// lights
const int leds = 7; // leds on-off
const int ledToggle = 6;


void setup() {

// Serial

 Serial.begin(9600);

 // BMP180
 if (pressure.begin()){
 Serial.println("BMP180 OK");
 digitalWrite(leds, HIGH);
 delay(500);
 digitalWrite(leds, LOW);
 }
 else
 {
  Serial.println("BMP180 NOT OK");
  while(1);
 }

  
// outputs
pinMode(motorRear1, OUTPUT);
pinMode(motorRear2, OUTPUT);
pinMode(motorRear3, OUTPUT);
pinMode(motorFront1, OUTPUT);
pinMode(motorFront2, OUTPUT);
pinMode(motorFront3, OUTPUT);

pinMode (leds, OUTPUT);

// inputs
pinMode(forward, INPUT);
digitalWrite(forward, LOW);
pinMode(backward, INPUT);
//digitalWrite(backward, LOW);
pinMode(left, INPUT);
//digitalWrite(left, LOW);
pinMode(right, INPUT);
//digitalWrite(right, LOW);

pinMode(ledToggle, INPUT);
//digitalWrite(leds, LOW);



}

void loop() {

// BMP180

char status;
double T, P, p0, a;

/*
Serial.println();
Serial.print(ALTITUDE,0);
Serial.print(" meter");
*/

status = pressure.getPressure(P,T);
        if (status != 0)
        {
          // Print out the measurement:
          Serial.println("absolute pressure: ");
          Serial.print(P,2);
          Serial.print(" mb, ");
        }

// Motor control

buttonStateForward = digitalRead(forward);
buttonStateBackward = digitalRead(backward);
buttonStateLeft = digitalRead(left);
buttonStateRight = digitalRead(right);
//buttonStateUP = digitalRead(up);
//buttonStateDown = digitalRead(down);

if ((buttonStateLeft == HIGH) && (buttonStateRight == LOW) && (buttonStateBackward == LOW) && (buttonStateForward == LOW )) {
  digitalWrite(motorRear1, LOW);
  digitalWrite(motorRear2,HIGH);
  digitalWrite(motorRear3,LOW);
  digitalWrite(motorFront1,LOW);
  digitalWrite(motorFront2,LOW);
  digitalWrite(motorFront3,HIGH);
  Serial.println("LEFT");
}
else if ((buttonStateLeft == LOW) && (buttonStateRight == HIGH) && (buttonStateBackward == LOW) && (buttonStateForward == LOW )) {
  digitalWrite(motorRear1,LOW);
  digitalWrite(motorRear2,LOW);
  digitalWrite(motorRear3,HIGH);
  digitalWrite(motorFront1,HIGH);
  digitalWrite(motorFront2,LOW);
  digitalWrite(motorFront3,LOW);
  Serial.println("RIGHT");
}
else if ((buttonStateLeft == LOW) && (buttonStateRight == LOW) && (buttonStateBackward == HIGH) && (buttonStateForward == LOW )) {
  digitalWrite(motorRear1, LOW);
  digitalWrite(motorRear2,LOW);
  digitalWrite(motorRear3,LOW);
  digitalWrite(motorFront1,HIGH);
  digitalWrite(motorFront2,HIGH);
  digitalWrite(motorFront3,HIGH);
  Serial.println("RWD");
}
else if ((buttonStateLeft == LOW) && (buttonStateRight == LOW) && (buttonStateBackward == LOW) && (buttonStateForward == HIGH )) {
digitalWrite(motorRear1, HIGH);
  digitalWrite(motorRear2,HIGH);
  digitalWrite(motorRear3,HIGH);
  digitalWrite(motorFront1,LOW);
  digitalWrite(motorFront2,LOW);
  digitalWrite(motorFront3,LOW);
  Serial.println("FWD");
}
else {
  digitalWrite(motorRear1,LOW);
  digitalWrite(motorRear2,LOW);
  digitalWrite(motorRear3,LOW);
  digitalWrite(motorFront1,LOW);
  digitalWrite(motorFront2,LOW);
  digitalWrite(motorFront3,LOW);
}

// leds
if (digitalRead(ledToggle == HIGH)){
  digitalWrite(leds, HIGH);
}
else{
  digitalWrite(leds, LOW);
}


} // end bracket


