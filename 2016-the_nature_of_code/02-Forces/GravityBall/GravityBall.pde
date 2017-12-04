class Ball {  
  PVector location;
  PVector velocity;
  PVector acceleration;
  int colorR;
  int colorG;
  int colorB;
  
  int size;

  Ball(PVector location, int size) {
    this.location = location;
    this.size = size;
    this.velocity = new PVector(0, 0);
    this.acceleration = new PVector(0, 0);
    this.colorR = int(random(255));
    this.colorG = 0;
    this.colorB = 204;
  }

  void draw() {
    noStroke();
    fill(this.colorR, this.colorG, this.colorB);
    ellipse(this.location.x, this.location.y, size, size);
  }

  void applyForce(PVector force) {
    PVector f = force.copy();
    this.acceleration.add(f.div(mass()));
  }

  private int mass() {
    return size * 50;
  }

  void update() {
    this.velocity.add(this.acceleration);
    this.location.add(this.velocity);
    if (this.location.y > height) {
      this.location.y = height;
      this.velocity.y = -this.velocity.y;
    }
    if (this.location.x > width) {
      this.location.x = width;
      this.velocity.x = -this.velocity.x;
    }
    if (this.location.x < 0) {
      this.location.x = 0;
      this.velocity.x = -this.velocity.x;
    }
    this.acceleration.mult(0);
  }
}

Ball[] balls;
PVector gravity = new PVector(0, 50);
PVector wind = new PVector(3, 0);

void setup() {
  size(300, 700);
  int numOfBalls = 200;
  balls = new Ball[numOfBalls];
  for (int i = 0; i < balls.length; i++) { 
    balls[i] = new Ball(new PVector(random(width), random(height)), int(random(20) + 10));
  }
}


void draw() {
  background(255, 204, 0);
  for (Ball ball : balls) {
    ball.applyForce(gravity);
    ball.applyForce(wind);
    ball.update();
    ball.draw();
  }
}