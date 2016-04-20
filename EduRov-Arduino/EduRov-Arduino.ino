// string to keep data received over serial
String readString;

// motor control outputs
const int motorRearLeft = 2;
const int motorRearRight = 3;
const int motorFrontLeft = 4;
const int motorFrontRight = 5;

// led outputs
const int ledPin = 13;

void setup() {
  Serial.begin(9600);
  // outputs
  pinMode(motorRearLeft, OUTPUT);
  pinMode(motorRearRight, OUTPUT);
  pinMode(motorFrontLeft, OUTPUT);
  pinMode(motorFrontRight, OUTPUT);
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
  digitalWrite(motorRearLeft, LOW);
  digitalWrite(motorRearRight,HIGH);
  digitalWrite(motorFrontLeft,HIGH);
  digitalWrite(motorFrontRight,LOW);
}

void motorRight(){
  digitalWrite(motorRearLeft, HIGH);
  digitalWrite(motorRearRight,LOW);
  digitalWrite(motorFrontLeft,LOW);
  digitalWrite(motorFrontRight,HIGH);
}

void motorForward(){
  digitalWrite(motorRearLeft, HIGH);
  digitalWrite(motorRearRight,HIGH);
  digitalWrite(motorFrontLeft, LOW);
  digitalWrite(motorFrontRight, LOW);
}

void motorBackward(){
  digitalWrite(motorRearLeft, LOW);
  digitalWrite(motorRearRight,LOW);
  digitalWrite(motorFrontLeft,HIGH);
  digitalWrite(motorFrontRight,HIGH);
}

void motorStop(){
  digitalWrite(motorRearLeft, LOW);
  digitalWrite(motorRearRight,LOW);
  digitalWrite(motorFrontLeft, LOW);
  digitalWrite(motorFrontRight, LOW);
}
