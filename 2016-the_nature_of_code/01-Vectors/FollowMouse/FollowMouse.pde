class Ball {
  PVector location;
  PVector velocity;
  PVector accelaration;
  
  int maxVelocity = 7;
  
  Ball() {
    location = new PVector(random(width), random(height));
    velocity = new PVector(0, 0);
  }
  
  void update() {
    PVector mouse = new PVector(mouseX, mouseY);
    PVector direction = mouse.sub(location);
    accelaration = direction.normalize().mult(0.1);
    velocity.add(accelaration);
    velocity.limit(maxVelocity);
    location.add(velocity);
  }
  
  void draw() {
    ellipse(location.x, location.y, 40, 40);
  }
}

Ball[] balls;
  
void setup() {
  size(600, 600);
  balls = new Ball[] {
    new Ball(), new Ball(), new Ball(), new Ball(), new Ball(), new Ball()
  };
}

void draw() {
  background(255);
  for (Ball ball : balls) {
    ball.update();
    ball.draw();
  }
}