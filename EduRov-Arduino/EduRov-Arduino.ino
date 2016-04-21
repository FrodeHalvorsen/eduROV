// string to keep data received over serial
String readString;

// motor control outputs
const int motorRear1 = 2; // 12 oclock
const int motorRear2 = 3; // 4 oclock
const int motorRear3 = 4; // 8 oclock
const int motorFront1 = 5; // // 2oclock
const int motorFront2 = 6; // 6 oclock
const int motorFront3 = 7; // 10 oclock

// led outputs
const int ledPin = 13;

void setup() {
  Serial.begin(9600);
  // outputs
  pinMode(motorRear1, OUTPUT);
  pinMode(motorRear2, OUTPUT);
  pinMode(motorRear3, OUTPUT);
  pinMode(motorFront1, OUTPUT);
  pinMode(motorFront2, OUTPUT);
  pinMode(motorFront3, OUTPUT);
  pinMode(ledPin, OUTPUT); 
}

unsigned int i=0;

void loop() {

  while (Serial.available()) {
    delay(3);  
    char c = Serial.read();
    readString += c; 
  }
  readString.trim();
  if (readString.length() >0) {
      if(readString == "left"){
        motorLeft();
      }
      if(readString == "right"){
        motorRight();
      }
      if(readString == "bck"){
        motorBackward();
      }
      if(readString == "fwd"){
        motorForward();
      }
      if(readString == "stop"){
        motorStop();
      }
        readString = "";
  }

  if(i>=64000){
    Serial.println("TestData");
    i=0;
  }
  i++;

}

void motorLeft(){
  digitalWrite(motorRear1, LOW);
  digitalWrite(motorRear2,HIGH);
  digitalWrite(motorRear3,LOW);
  digitalWrite(motorFront1,HIGH);
  digitalWrite(motorFront2,LOW); 
  digitalWrite(motorFront3,LOW);
}

void motorRight(){
  digitalWrite(motorRear1, LOW);
  digitalWrite(motorRear2,LOW);
  digitalWrite(motorRear3,HIGH);
  digitalWrite(motorFront1,LOW);
  digitalWrite(motorFront2,LOW); 
  digitalWrite(motorFront3,HIGH);  
}

void motorForward(){
  digitalWrite(motorRear1,HIGH);
  digitalWrite(motorRear2,HIGH);
  digitalWrite(motorRear3,HIGH);
  digitalWrite(motorFront1,LOW);
  digitalWrite(motorFront2,LOW); 
  digitalWrite(motorFront3,LOW);
}

void motorBackward(){
  digitalWrite(motorRear1, LOW);
  digitalWrite(motorRear2,LOW);
  digitalWrite(motorRear3,LOW);
  digitalWrite(motorFront1,HIGH);
  digitalWrite(motorFront2,HIGH); 
  digitalWrite(motorFront3,HIGH);
}

void motorStop(){
  digitalWrite(motorRear1, LOW);
  digitalWrite(motorRear2,LOW);
  digitalWrite(motorRear3,LOW);
  digitalWrite(motorFront1,LOW);
  digitalWrite(motorFront2,LOW); 
  digitalWrite(motorFront3,LOW);
}
