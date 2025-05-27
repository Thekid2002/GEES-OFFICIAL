export class GestureActionMapping {
  id: number | null;
  action_id: number;
  gesture_id: number;

  constructor(id: number | null, action_id: number, gesture_id: number) {
    this.id = id;
    this.action_id = action_id;
    this.gesture_id = gesture_id;
  }
}
