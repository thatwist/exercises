class Ball {
  PVector location;
  PVector velocity = new PVector(3, 3);
  int radius = 10;
  
  
  Ball(PVector location) {
    this.location = location;
  }
  
  void move(PVector boundary) {
    location.add(velocity);
    if ((location.x + radius > boundary.x) || (location.x - radius < 0)) {
      velocity.x = -velocity.x;
    }
    if ((location.y + radius > boundary.y) || (location.y - radius < 0)) {
      velocity.y = -velocity.y;
    }
  }
  
  void draw() {
    background(255);
    fill(50, 100, 100);
    ellipse(location.x, location.y, radius, radius);
  }
}

Ball ball;

void setup() {
  size(300, 200);
  ball = new Ball(new PVector(20, 20));
}

void draw() {
  ball.move(new PVector(width, height));
  ball.draw();
}