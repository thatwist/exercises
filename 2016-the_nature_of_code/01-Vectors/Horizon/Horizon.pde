int i = 0;

void setup() {
  size(600, 300);
}

void draw() {
  background(255, 120, 0);
  stroke(0, 0, 0, 150);
  for (int x = 0; x < width; x++) {
    int y =int(noise((x + i)/150.0) * 300);
    line(x, y, x, height);
  }
  i++;
}