
#include <SFE_BMP180.h>
#include <Wire.h>


// BMP180
SFE_BMP180 pressure;
#define ALTITUDE 1655.0


// motor control outputs
const int motorRearLeft = 2;
const int motorRearRight = 3;
const int motorFrontLeft = 4;
const int motorFrontRight = 5;

// motor control inputs
const int forward = 10;
const int backward = 9;
const int left = 8;
const int right = 11;

int buttonStateForward = 0;
int buttonStateBackward = 0;
int buttonStateLeft = 0;
int buttonStateRight = 0;


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
pinMode(motorRearLeft, OUTPUT);
pinMode(motorRearRight, OUTPUT);
pinMode(motorFrontLeft, OUTPUT);
pinMode(motorFrontRight, OUTPUT);

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

if ((buttonStateLeft == HIGH) && (buttonStateRight == LOW) && (buttonStateBackward == LOW) && (buttonStateForward == LOW )) {
  digitalWrite(motorRearLeft, LOW);
  digitalWrite(motorRearRight,HIGH);
  digitalWrite(motorFrontLeft,HIGH);
  digitalWrite(motorFrontRight,LOW);
  Serial.println("LEFT");
}
else if ((buttonStateLeft == LOW) && (buttonStateRight == HIGH) && (buttonStateBackward == LOW) && (buttonStateForward == LOW )) {
  digitalWrite(motorRearLeft, HIGH);
  digitalWrite(motorRearRight,LOW);
  digitalWrite(motorFrontLeft,LOW);
  digitalWrite(motorFrontRight,HIGH);
  Serial.println("RIGHT");
}
else if ((buttonStateLeft == LOW) && (buttonStateRight == LOW) && (buttonStateBackward == HIGH) && (buttonStateForward == LOW )) {
  digitalWrite(motorRearLeft, LOW);
  digitalWrite(motorRearRight,LOW);
  digitalWrite(motorFrontLeft,HIGH);
  digitalWrite(motorFrontRight,HIGH);
  Serial.println("RWD");
}
else if ((buttonStateLeft == LOW) && (buttonStateRight == LOW) && (buttonStateBackward == LOW) && (buttonStateForward == HIGH )) {
  digitalWrite(motorRearLeft, HIGH);
  digitalWrite(motorRearRight,HIGH);
  digitalWrite(motorFrontLeft, LOW);
  digitalWrite(motorFrontRight, LOW);
  Serial.println("FWD");
}else {
   digitalWrite(motorRearLeft, LOW);
  digitalWrite(motorRearRight,LOW);
  digitalWrite(motorFrontLeft, LOW);
  digitalWrite(motorFrontRight, LOW); 
}

// leds
if (digitalRead(ledToggle == HIGH)){
  digitalWrite(leds, HIGH);
}
else{
  digitalWrite(leds, LOW);
}


} // end bracket


