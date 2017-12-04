class Balloon {  
  PVector location;
  PVector velocity;
  PVector acceleration;
  int size;

  Balloon(PVector location, int size) {
    this.location = location;
    this.size = size;
    this.velocity = new PVector(0, 0);
    this.acceleration = new PVector(0, 0);
  }

  void draw() {
    noStroke();
    fill(255, 0, 204);
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
    if (this.location.y < 0) {
      this.location.y = 0;
      this.velocity.y = -this.velocity.y;
    }
  }
}

Balloon balloon;
PVector gravity = new PVector(0, -1);

void setup() {
  size(300, 200);
  int balloonSize = 70; 
  balloon = new Balloon(new PVector(width/2, height), balloonSize);
}


void draw() {
  background(255, 204, 0);
  balloon.applyForce(gravity);
  balloon.update();
  balloon.draw();
}