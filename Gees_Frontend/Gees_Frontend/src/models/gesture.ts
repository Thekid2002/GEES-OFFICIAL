export class Gesture {
  id: number | null;
  name: string;
  description: string | null;
  image_url: string | null;

  constructor(id: number | null, name: string, description: string | null, image_url: string | null) {
    this.id = id;
    this.name = name;
    this.description = description;
    this.image_url = image_url;
  }
}
