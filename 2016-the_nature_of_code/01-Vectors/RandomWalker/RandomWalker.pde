class Walker {
  float x, y;
  Walker() {
    x = 100;
    y = 100;
  }

  void step() {
    float stepX = random(-1, 1);
    float stepY = random(-1, 1);
    x += stepX;
    y += stepY;
  }
  
  void display() {
    point(x, y);
  }
}

Walker walker;

void setup() {
  size(300, 200);
  walker = new Walker();
}

void draw() {
  walker.step();
  walker.display();
}