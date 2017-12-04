void setup() {
  size(400, 400);
  
}

float a = 0;

void draw() {
  a++;
  background(216, 30, 91);
  noStroke();
  drawCircle(width/2, height/2, width/5);
}

void drawCircle(int x, int y, int r) {
  fill(51, 24, 50, 255);
  rotate(radians(a/10000));
  translate(a/10000, a/10000);
  ellipse(x, y, r, r);
  if (r > 1) {
    drawCircle(x/2, y, r/2);
    drawCircle(width/2 + x/2, y, r/2);
    drawCircle(x, height/2 + y/2, r/2);
    drawCircle(x, y/2, r/2);
  }
  
}