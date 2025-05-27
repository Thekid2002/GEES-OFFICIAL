export enum MessageType {
  COUNTDOWN = 'countdown',
  START = 'start',
  STOP = 'stop',
  ERROR = 'error',
}

export class WSMessage {
  msgType: MessageType
  data: any

  constructor(msgType: MessageType, data: any) {
    this.msgType = msgType
    this.data = data
  }
}
