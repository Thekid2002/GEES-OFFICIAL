export class PlotData {
  private data: { x: number[]; y: number[]; z: number[] };
  private maxPoints: number;

  constructor(data: { x: number[]; y: number[]; z: number[] } = { x: [], y: [], z: [] }) {
    this.maxPoints = 8;
    this.data = data;
  }

  // Add a new point to the data
  addPoint(x: number, y: number, z: number) {
    this.data.x.push(x);
    this.data.y.push(y);
    this.data.z.push(z);

    // Limit the number of points
    if (this.data.x.length > this.maxPoints) {
      this.data.x.shift();
      this.data.y.shift();
      this.data.z.shift();
    }
  }

  // Get the current data
  getData() {
    return this.data;
  }
}
