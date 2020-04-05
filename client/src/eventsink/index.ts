import Vue from "vue";
import { VueConstructor } from "vue/types/umd";
import { BaseUrl } from "../constants";
const eventUrl = `${BaseUrl}events/dunno`;
type VueSSE = VueConstructor<Vue> & {
  SSE: (
    url: string,
    options: { format: string; withCredentials: boolean }
  ) => Promise<any>;
};
export type eventType<T> = {
  data: T;
  event: string;
};
export class EventSink {
  private server: any;
  public async init(onMessage: Function) {
    const higherOrderVue: VueSSE = Vue as VueSSE;
    this.server = await higherOrderVue.SSE(eventUrl, {
      format: "json",
      withCredentials: true
    });
    this.server.subscribe("", (message: any) => {
      console.log("received ", message);
      onMessage(message);
    });
  }
  public async destory() {
    if (!this.server) {
      return;
    }
    this.server.close();
  }
}
