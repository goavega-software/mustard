<template>
    <div>
      <div class="buttonClass">
        <button class="button" @click="saveJson">Save widget dimensions</button>
      </div>
      <grid-layout 
              :layout.sync="layout"
              :col-num="6"
              :row-height="185"
              :is-draggable="draggable"
              :is-resizable="resizable"
              :isBounded="bounded"
              :vertical-compact="true"
              :use-css-transforms="true"
      >
        <grid-item 
              v-for="(item, index) in layout"
              v-bind:key=index
              v-bind:class="getClass(index)"
              :static="item.static"
              :x="item.x"
              :y="item.y"
              :w="item.w"
              :h="item.h"
              :i="item.i"
              :minW="minW"
              :minH="minH"
              @resized="resizedEvent"
              @moved="movedEvent"
              @container-resized="containerResizedEvent"
        >
        <div :is="item.component" v-bind="getProps(index)" class="widgetClass"/>
      </grid-item>
      </grid-layout>
    </div>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import Clock from "./Clock.vue";
import TextWidget from "./TextWidget.vue";
import ListWidget from "./ListWidget.vue";
import ComparisonWidget from "./ComparisonWidget.vue";
import WeatherWidget from "./WeatherWidget.vue";
import SlideshowWidget from "./SlideshowWidget.vue";
import { EventSink, eventType } from "../eventsink";
import { BaseUrl } from "../constants";
import { GridLayout, GridItem } from "vue-grid-layout";
import { VNodeChildrenArrayContents } from "vue/types/umd";

type layoutType = {
  x: number;
  y: number;
  w: number;
  h?: number;
  i: string;
  index: number;
  class: string | string[];
  props: Record<string, unknown>;
  state?: Record<string, unknown>;
  component: string;
};

interface Classes {
  column: boolean;
  [key: string]: boolean;
}

@Component({
  components: {
    Clock,
    TextWidget,
    ListWidget,
    ComparisonWidget,
    WeatherWidget,
    SlideshowWidget,
    GridLayout,
    GridItem
  }
})
export default class App extends Vue {
  private eventServer = new EventSink();
  private layout: Array<layoutType> = [];
  private tempDimensionJson: Array<layoutType> = [];
  private dimensionJson: Array<layoutType> = [];

  created() {
    setTimeout(
      async () =>
        await fetch(`${BaseUrl}api/nudge`, {
          method: "POST"
        }),
      1000
    );
  }
  mounted() {
    (async () => {
      this.tempDimensionJson = this.layout;
      this.layout.forEach(layoutItem => {
        if (!layoutItem.state) {
          return;
        }
        const moduleState = {
          state: () => ({ data: {} })
        };
        const stateKey = Object.keys(layoutItem.state)[0];
        this.$store.registerModule(stateKey, moduleState);
      });
    })();
    (async () => {
      await this.eventServer.init((message: eventType<object>) =>
        this.$store.dispatch("change", message)
      );
    })();
  }
  
  data() {
    return {
        layout: [
            {
              "x":0,
              "y":0,
              "w":2,
              "h":2,
              "class": ["x1 y2"],
              "i":"bsGenerator", 
              "props": {
                "eventId": "bsGenerator",
                "background": "#ffdb58"
              },
              "static": false,
              "component": "TextWidget",
              "state": {
                "bsGenerator": {
                  "title": "Y",
                  "subtitle": "X"
                }
              }
            },
            {
              "x":2,
              "y":0,
              "w":2,
              "h":5,
              "class": ["x2 y4"],
              "i":"jotd", 
              "props": {
                "eventId": "jotd"
              },
              "static": false,
              "component": "TextWidget",
              "state": {
                "jotd": {
                  "title": "Y",
                  "subtitle": "X"
                }
              }
            },
            {
              "x":4,
              "y":0,
              "w":2,
              "h":2,
              "class": ["x1 y2"],
              "i":"clock", 
              "props": {
                "eventId": "clockWidget",
                "clockOneTz": "America/Los_Angeles",
                "clockThreeTz": "America/New_York"
              },
              "static": false,
              "clockOneTz": "America/Los_Angeles",
              "clockThreeTz": "America/New_York",
              "component": "Clock",
              "state": {
                "clockWidget": {}
              }
            }
        ],
        draggable: true,
        resizable: true,
        bounded: false,
        index: 0,
        minW: 2,
        minH: 2,
    }
  }

  movedEvent(i:string, newX:number, newY:number) {
    const index = i== 'bsGenerator' ? 0 : (i== 'jotd' ? 1 : 2)
    this.tempDimensionJson[index]['i'] = i;
    this.tempDimensionJson[index]['x'] = newX;
    this.tempDimensionJson[index]['y'] = newY;
  }
  resizedEvent(i:string, newH:number, newW:number, newHPx:number, newWPx:number) {
    const index = i== 'bsGenerator' ? 0 : (i== 'jotd' ? 1 : 2)
    this.tempDimensionJson[index]['i'] = i;
    this.tempDimensionJson[index]['h'] = newH;
    this.tempDimensionJson[index]['w'] = newW;
  }
  containerResizedEvent(i:string, newH:number, newW:number, newHPx:number, newWPx:number) {
    const index = i== 'bsGenerator' ? 0 : (i== 'jotd' ? 1 : 2)
    this.tempDimensionJson[index]['i'] = i;
    this.tempDimensionJson[index]['h'] = newH;
    this.tempDimensionJson[index]['w'] = newW;
  }

  getClass(index: number): { column: boolean } & Record<string, boolean> {
    const classes: Classes = {
      column: false
    };
    const classNames = Array.isArray(this.layout[index].class)
      ? this.layout[index].class
      : [this.layout[index].class];
    (classNames as string[]).forEach((cls: string|number) => {
      classes[cls] = false;
    });
    return classes;
  }

  getProps(index: number): Record<string, unknown> {
    return this.layout[index].props;
  }
  
  saveJson() {
    this.dimensionJson = this.tempDimensionJson;
    console.log("🚀 ~ file: App.vue:252 ~ App ~ saveJson ~ dimensionJson:", this.dimensionJson)
    console.log('button clicked!!!!!!!')
  }
  beforeDestroy() {
    this.eventServer.destory();
  }
}
</script>

<style lang="scss">

.vue-grid-layout {
    background: #333;
    width: 100%;
    height: 100vh !important;
    overflow: hidden;
}
.vue-grid-item:not(.vue-grid-placeholder) {
    background: #ccc;
    border: 0.5px solid black;
}
.vue-grid-item .resizing {
    opacity: 0.9;
}
.vue-grid-item .static {
    background: #cce;
}
.vue-grid-item .no-drag {
    height: 100%;
    width: 100%;
}
.vue-grid-item .minMax {
    font-size: 12px;
}
.vue-grid-item .add {
    cursor: pointer;
}
.vue-draggable-handle {
    position: absolute;
    width: 20px;
    height: 20px;
    top: 0;
    left: 0;
    background: url("data:image/svg+xml;utf8,<svg xmlns='http://www.w3.org/2000/svg' width='10' height='10'><circle cx='5' cy='5' r='5' fill='#999999'/></svg>") no-repeat;
    background-position: bottom right;
    padding: 0 8px 8px 0;
    background-repeat: no-repeat;
    background-origin: content-box;
    box-sizing: border-box;
    cursor: pointer;
}
.buttonClass {
  background: #333;
  display: flex;  
  justify-content: center;  
  align-items: center;
}
.button {
  background-color: #035880;
  color: white;
  width: fit-content;
  max-width: 100%;
  min-width: 100px;
  padding: 10px 24px;
  font-weight: bold;
}
.widgetClass {
  position: relative !important;
}
</style>
