int cols = 120;
int rows = 50;
float scale = 7f;

float[][] altitudes = new float[rows][cols];
float i = 0;
float k = 0;
int j = 1;

void setup() {
  size(800, 400, P3D);
}

void draw() {
  background(50, 0, 50);
  
  translate(width/2-cols*scale/2, height/2-rows*scale/2 + 150);
  
  noStroke();
  
  i += 0.1;
  if ((k > 100) || k < 0) {
    j = -j;
  }
  k += 0.05 * j;
  
  for (int y = 0; y < rows; y++) {
    for (int x = 0; x < cols; x++) {
      altitudes[y][x] = map(noise(i - y / scale, x / scale ), 0, 1, 0, k);
    }
  }
  
  for (int y = 0; y < rows - 1; y++) {
    beginShape(TRIANGLE_STRIP);
    for (int x = 0; x < cols; x++) {
      fill(20, altitudes[y][x], altitudes[y][x]*2);
      vertex(x*scale, y*scale, altitudes[y][x]);
      vertex(x*scale, (y+1)*scale, altitudes[y+1][x]);
    }
    endShape();
  }
}